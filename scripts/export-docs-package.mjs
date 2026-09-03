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
  bug: "CAUTION",
  quote: "NOTE",
  注意: "NOTE",
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
      const inlineClose = line.indexOf(":::");
      if (inlineClose >= 0) {
        const content = line.slice(0, inlineClose).trim();
        const suffix = line.slice(inlineClose + 3).trim();
        if (content) output.push(`> ${content}`);
        container = null;
        if (suffix) lines.splice(lineNumber + 1, 0, suffix);
        continue;
      }
      output.push(line ? `> ${line}` : ">");
      continue;
    }

    const start = line.match(/^\s*:::([^\s\[]+)(?:\[([^\]]+)\])?(?:\s+(.*))?$/u);
    if (!start) {
      if (/^\s*:::\s*$/.test(line)) {
        const legacy = lines[lineNumber + 1]?.match(/^\s*([^\s]+)(?:\s+(.*))?$/u);
        const legacyAlert = legacy && directiveTypes[legacy[1].toLowerCase()];
        if (!legacyAlert) throw new Error(`orphan directive close at line ${lineNumber + 1}`);
        output.push(`> [!${legacyAlert}]`);
        if (legacy[2]) output.push(`> ${legacy[2]}`);
        container = { type: legacy[1].toLowerCase(), lineNumber: lineNumber + 1 };
        lineNumber += 1;
        continue;
      }
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
  const sourceByRelative = new Map(files.map((file) => [path.relative(source, file).replaceAll("\\", "/").toLowerCase(), file]));
  const portablePath = (value) => value.toLowerCase().replace(/[()]/g, "");
  const sourceByPortable = new Map();
  for (const [relative, file] of sourceByRelative) {
    const portable = portablePath(relative);
    if (!sourceByPortable.has(portable)) sourceByPortable.set(portable, file);
  }
  const sourceForPage = (page) => sourceByRelative.get(page.toLowerCase())
    || sourceByRelative.get(page.replace(/\.md$/i, "/index.md").toLowerCase())
    || sourceByPortable.get(portablePath(page));
  for (const page of pageSet) if (!sourceForPage(page)) throw new Error(`${source} is missing navigation page ${page}`);
  const outputBySource = new Map();
  for (const page of pageSet) outputBySource.set(sourceForPage(page), page);
  for (const file of files) if (!outputBySource.has(file)) outputBySource.set(file, path.relative(source, file).replaceAll("\\", "/"));
  let directives = 0;
  for (const [sourceFile, relative] of outputBySource) {
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
  return { documents: outputBySource.size, directives, assets };
}

export function exportPackage(key, repoRoot = process.cwd()) {
  const contentRoot = path.join(repoRoot, "src/content/docs");
  const nav = parseNav(path.join(repoRoot, "src/nav", `${key}.ts`));
  const pageOrder = new Map();
  const pageSet = new Set();
  const navigation = buildNavigation(nav.items, key, pageOrder, pageSet);
  const tempRoot = fs.mkdtempSync(path.join(os.tmpdir(), `docs-export-${key}-`));
  try {
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
    return { key, zipPath, reportPath, report };
  } finally {
    fs.rmSync(tempRoot, { recursive: true, force: true });
  }
}

export function discoverDocumentKeys(repoRoot = process.cwd()) {
  const directoryNames = (root) => new Set(
    fs.readdirSync(root, { withFileTypes: true }).filter((entry) => entry.isDirectory()).map((entry) => entry.name),
  );
  const contentRoot = path.join(repoRoot, "src/content/docs");
  const zh = directoryNames(path.join(contentRoot, "zh-cn"));
  const en = directoryNames(path.join(contentRoot, "en"));
  return fs.readdirSync(path.join(repoRoot, "src/nav"), { withFileTypes: true })
    .filter((entry) => entry.isFile() && entry.name.endsWith(".ts"))
    .map((entry) => path.basename(entry.name, ".ts"))
    .filter((key) => zh.has(key) && en.has(key))
    .sort();
}

export function exportAllPackages(repoRoot = process.cwd()) {
  const succeeded = [];
  const failed = [];
  for (const key of discoverDocumentKeys(repoRoot)) {
    try {
      succeeded.push(exportPackage(key, repoRoot));
    } catch (error) {
      failed.push({ key, error: error instanceof Error ? error.message : String(error) });
    }
  }
  const outputDir = path.join(repoRoot, "exports");
  fs.mkdirSync(outputDir, { recursive: true });
  const reportPath = path.join(outputDir, "docs-export-batch.report.json");
  fs.writeFileSync(reportPath, JSON.stringify({ succeeded: succeeded.map(({ key, zipPath, reportPath: itemReportPath }) => ({ key, zipPath, reportPath: itemReportPath })), failed }, null, 2) + "\n");
  return { succeeded, failed, reportPath };
}

if (process.argv[1] && import.meta.url === pathToFileURL(process.argv[1]).href) {
  const key = process.argv[2];
  if (!key) throw new Error("usage: node scripts/export-docs-package.mjs <document-key|--all>");
  if (key === "--all") {
    const result = exportAllPackages();
    process.stdout.write(`${JSON.stringify({ succeeded: result.succeeded.map(({ key: itemKey }) => itemKey), failed: result.failed })}\n${result.reportPath}\n`);
    if (result.failed.length) process.exitCode = 1;
  } else {
    const result = exportPackage(key);
    process.stdout.write(`${JSON.stringify(result.report)}\n${result.zipPath}\n`);
  }
}
