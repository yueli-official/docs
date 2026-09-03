import fs from "node:fs";
import os from "node:os";
import path from "node:path";
import process from "node:process";
import vm from "node:vm";
import { spawnSync } from "node:child_process";
import { pathToFileURL } from "node:url";

const directiveTypes = {
  note: "NOTE",
  info: "NOTE",
  tip: "TIP",
  important: "IMPORTANT",
  warning: "WARNING",
  caution: "CAUTION",
  danger: "CAUTION",
};

export function convertDirectiveContainers(markdown) {
  const lines = markdown.replace(/\r\n/g, "\n").split("\n");
  const output = [];
  let fence = "";
  let container = null;

  for (let lineNumber = 0; lineNumber < lines.length; lineNumber += 1) {
    const line = lines[lineNumber];
    const fenceMatch = line.match(/^\s*(```+|~~~+)/);
    if (fenceMatch && !container) {
      const marker = fenceMatch[1][0];
      if (!fence) fence = marker;
      else if (fence === marker) fence = "";
      output.push(line);
      continue;
    }
    if (fence) {
      output.push(line);
      continue;
    }

    if (container) {
      if (/^\s*:::\s*$/.test(line)) {
        container = null;
        continue;
      }
      output.push(line ? `> ${line}` : ">");
      continue;
    }

    const start = line.match(/^\s*:::([A-Za-z][\w-]*)(?:\[([^\]]+)\])?(?:\s+(.*))?$/);
    if (!start) {
      if (/^\s*:::\s*$/.test(line)) throw new Error(`orphan directive close at line ${lineNumber + 1}`);
      output.push(line);
      continue;
    }
    const [, rawType, bracketTitle, rawRest = ""] = start;
    const type = rawType.toLowerCase();
    const alert = directiveTypes[type];
    if (!alert) throw new Error(`unsupported directive ${type} at line ${lineNumber + 1}`);
    const inlineClose = rawRest.indexOf(":::");
    const content = inlineClose >= 0 ? rawRest.slice(0, inlineClose).trim() : rawRest.trim();
    const suffix = inlineClose >= 0 ? rawRest.slice(inlineClose + 3).trim() : "";
    output.push(`> [!${alert}]${bracketTitle ? ` ${bracketTitle.trim()}` : ""}`);
    if (content) output.push(`> ${content}`);
    if (inlineClose >= 0) {
      if (suffix) output.push("", suffix);
    } else {
      container = { type, lineNumber: lineNumber + 1 };
    }
  }
  if (container) throw new Error(`unclosed directive ${container.type} at line ${container.lineNumber}`);
  return output.join("\n");
}

function parseNav(file) {
  const source = fs.readFileSync(file, "utf8");
  const objectSource = source.slice(source.indexOf("{"), source.lastIndexOf("};") + 1);
  return vm.runInNewContext(`(${objectSource})`, Object.create(null), { timeout: 1000 });
}

function markdownFiles(root) {
  const result = [];
  for (const entry of fs.readdirSync(root, { withFileTypes: true })) {
    const full = path.join(root, entry.name);
    if (entry.isDirectory()) result.push(...markdownFiles(full));
    else if (entry.isFile() && entry.name.toLowerCase().endsWith(".md")) result.push(full);
  }
  return result.sort();
}

function addFrontMatter(raw, relative, order) {
  const converted = convertDirectiveContainers(raw);
  const id = relative.replace(/\.md$/i, "").replaceAll("\\", "/");
  const additions = [`id: ${id}`, `order: ${order}`];
  if (converted.startsWith("---\n")) {
    const end = converted.indexOf("\n---", 4);
    if (end >= 0) {
      const frontMatter = converted.slice(4, end).split("\n").filter((line) => !/^(id|order):\s*/.test(line));
      return `---\n${frontMatter.join("\n")}\n${additions.join("\n")}\n---${converted.slice(end + 4)}`;
    }
  }
  return `---\n${additions.join("\n")}\n---\n\n${converted}`;
}

function buildNavigation(items, key, pageOrder, pageSet) {
  return items.map((item, index) => {
    if (item.link) {
      const relative = item.link.replace(new RegExp(`^/?${key}/`), "") + ".md";
      pageOrder.set(relative, index + 1);
      pageSet.add(relative);
      return { page: relative };
    }
    return {
      group: item.translations?.["zh-CN"] || item.label,
      translations: { "en-US": item.label },
      children: buildNavigation(item.items || [], key, pageOrder, pageSet),
    };
  });
}

function copyLocale(source, target, pageOrder, pageSet) {
  const files = markdownFiles(source);
  const relativeSet = new Set(files.map((file) => path.relative(source, file).replaceAll("\\", "/")));
  for (const page of pageSet) if (!relativeSet.has(page)) throw new Error(`${source} is missing navigation page ${page}`);
  let directives = 0;
  for (const relative of relativeSet) {
    const sourceFile = path.join(source, relative);
    const raw = fs.readFileSync(sourceFile, "utf8");
    directives += (raw.match(/^:::[A-Za-z][\w-]*/gm) || []).length;
    const targetFile = path.join(target, relative);
    fs.mkdirSync(path.dirname(targetFile), { recursive: true });
    fs.writeFileSync(targetFile, addFrontMatter(raw, relative, pageOrder.get(relative) ?? 999), "utf8");
  }
  let assets = 0;
  const copyAssets = (sourceDir, targetDir) => {
    for (const entry of fs.readdirSync(sourceDir, { withFileTypes: true })) {
      const sourceFile = path.join(sourceDir, entry.name);
      const targetFile = path.join(targetDir, entry.name);
      if (entry.isDirectory()) copyAssets(sourceFile, targetFile);
      else if (/\.(png|jpe?g|webp|gif)$/i.test(entry.name)) {
        fs.mkdirSync(targetDir, { recursive: true });
        fs.copyFileSync(sourceFile, targetFile);
        assets += 1;
      }
    }
  };
  copyAssets(source, target);
  return { documents: relativeSet.size, directives, assets };
}

export function exportPackage(key, repoRoot = process.cwd()) {
  const contentRoot = path.join(repoRoot, "src/content/docs");
  const nav = parseNav(path.join(repoRoot, "src/nav", `${key}.ts`));
  const pageOrder = new Map();
  const pageSet = new Set();
  const navigation = buildNavigation(nav.items, key, pageOrder, pageSet);
  const tempRoot = fs.mkdtempSync(path.join(os.tmpdir(), `docs-export-${key}-`));
  const packageRoot = path.join(tempRoot, key);
  fs.mkdirSync(packageRoot, { recursive: true });
  const zh = copyLocale(path.join(contentRoot, "zh-cn", key), packageRoot, pageOrder, pageSet);
  const en = copyLocale(path.join(contentRoot, "en", key), path.join(packageRoot, "locales/en-US"), pageOrder, pageSet);
  if (zh.documents !== en.documents) throw new Error(`locale document count mismatch: zh-CN=${zh.documents}, en-US=${en.documents}`);
  const manifest = {
    $schema: "https://docs.yueli.dev/schemas/docs-import-v1.json",
    schemaVersion: 1,
    defaultLocale: "zh-CN",
    locales: { "zh-CN": ".", "en-US": "locales/en-US" },
    navigation,
  };
  fs.writeFileSync(path.join(packageRoot, "docs.json"), JSON.stringify(manifest, null, 2) + "\n");
  const outputDir = path.join(repoRoot, "exports");
  fs.mkdirSync(outputDir, { recursive: true });
  const zipPath = path.join(outputDir, `${key}-docs-v1.zip`);
  const reportPath = path.join(outputDir, `${key}-docs-v1.report.json`);
  const command = `Compress-Archive -Path '${packageRoot.replaceAll("'", "''")}\\*' -DestinationPath '${zipPath.replaceAll("'", "''")}' -CompressionLevel Optimal -Force`;
  const zipped = spawnSync("pwsh", ["-NoProfile", "-Command", command], { encoding: "utf8" });
  if (zipped.status !== 0) throw new Error(zipped.stderr || zipped.stdout || "failed to create ZIP");
  const report = { package: "Docs Import Package v1", key, navigationPages: pageSet.size, locales: { "zh-CN": zh, "en-US": en } };
  fs.writeFileSync(reportPath, JSON.stringify(report, null, 2) + "\n");
  fs.rmSync(tempRoot, { recursive: true, force: true });
  return { zipPath, reportPath, report };
}

if (process.argv[1] && import.meta.url === pathToFileURL(process.argv[1]).href) {
  const key = process.argv[2];
  if (!key) throw new Error("usage: node scripts/export-docs-package.mjs <document-key>");
  const result = exportPackage(key);
  process.stdout.write(`${JSON.stringify(result.report)}\n${result.zipPath}\n`);
}
