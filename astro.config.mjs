// @ts-check
import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

import tailwindcss from "@tailwindcss/vite";
import starlightSidebarTopics from "starlight-sidebar-topics";

// navs
import { aeExpression } from "./src/nav/ae-expression";
import { aeScripting } from "./src/nav/ae-scripting";
import { javascriptTools } from "./src/nav/javascript-tools";
import { aePlugin } from "./src/nav/ae-plugin";
import { aiScripting } from "./src/nav/ai-scripting";
import { prScripting } from "./src/nav/pr-scripting";
import { prPlugin } from "./src/nav/pr-plugin";
import { houdiniVexNav } from "./src/nav/houdini-vex";

import starlightGiscus from "starlight-giscus";

import sitemap from "@astrojs/sitemap";

import starlightLlmsTxt from "starlight-llms-txt";

// https://astro.build/config
export default defineConfig({
  site: "https://docs.yuelili.com",
  title: "月离文档",
  description:
    "月离文档，专注于Adobe系列软件的脚本开发、插件开发、Houdini Vex等技术的分享与教程。",
  i18n: {
    defaultLocale: "en",
    locales: [
      {
        codes: ["en"],
        path: "en",
      },
      {
        codes: ["zh-CN"],
        path: "zh-cn",
      },
    ],
    routing: {
      prefixDefaultLocale: true,
      redirectToDefaultLocale: false,
      fallbackType: "redirect",
    },
  },
  markdown: {
    shikiConfig: {
      theme: "dracula",
      langs: ["applescript", "vb", "actionscript-3", "javascsript"],
      langAlias: {
        applescript: "applescript",
        actionscript: "actionscript-3",
        vbscript: "vb",
        none: "text",
        vex: "cpp",
      },
      defaultColor: false,
    },
  },

  integrations: [
    starlight({
      title: "月离文档",

      social: [
        {
          icon: "github",
          label: "GitHub",
          href: "https://github.com/yueli-official/docs",
        },
      ],
      components: {
        Head: "./src/components/Head.astro",
        Header: "./src/components/Header.astro",
        TwoColumnContent: "./src/components/TwoColumnContent.astro",
      },
      plugins: [
        starlightSidebarTopics([
          aeExpression,
          aeScripting,
          javascriptTools,
          aePlugin,
          prScripting,
          prPlugin,
          aiScripting,
          houdiniVexNav,
        ]),

        starlightGiscus({
          repo: "yueli-official/docs",
          repoId: "R_kgDOQz-TMw",
          category: "Comments",
          categoryId: "DIC_kwDOQz-TM84C0mHv",
        }),
        starlightLlmsTxt({
          projectName: "月离文档站",

          description: `
这是一个面向中文用户的专业技术文档站点，
内容涵盖 After Effects、Premiere Pro、Houdini、JavaScript、
以及相关脚本、插件与表达式的使用与开发说明。
`,

          details: `
文档主要用于：
- 技术学习与查询
- 表达式 / 脚本 / 插件 API 参考
- 实际项目中的开发与自动化支持

文档语言以中文为主，结构清晰，适合被大型语言模型检索与引用。
`,

          /** 推荐资源（非必须，但很加分） */
          optionalLinks: [
            {
              label: "首页",
              url: "/zh-cn/",
              description: "文档站总入口",
            },
          ],

          /** 按你首页的分类，给 AI 明确分组 */
          customSets: [
            {
              label: "AE 表达式",
              description: "After Effects 表达式中文文档",
              paths: ["zh-cn/ae-expression/**"],
            },
            {
              label: "AE 脚本",
              description: "After Effects 脚本开发文档",
              paths: ["zh-cn/ae-scripting/**"],
            },
            {
              label: "AE 插件",
              description: "After Effects 插件开发文档",
              paths: ["zh-cn/ae-plugin/**"],
            },
            {
              label: "JavaScript 工具",
              description: "AE / Adobe 相关 JavaScript 工具与教程",
              paths: ["zh-cn/javascript-tools/**"],
            },
            {
              label: "AI 脚本",
              description: "AI 相关脚本文档",
              paths: ["zh-cn/ai-scripting/**"],
            },
            {
              label: "PR 脚本",
              description: "Premiere Pro 脚本开发文档",
              paths: ["zh-cn/pr-scripting/**"],
            },
            {
              label: "PR 插件",
              description: "Premiere Pro 插件开发文档",
              paths: ["zh-cn/pr-plugin/**"],
            },
            {
              label: "Houdini VEX",
              description: "Houdini VEX 中文文档",
              paths: ["zh-cn/houdini-vex/**"],
            },
          ],

          /** 这些页面对 AI 最重要，放前面 */
          promote: [
            "zh-cn/index*",
            "**/introduction/**",
            "**/intro/**",
            "**/whats-new*",
            "**/changelog*",
          ],

          /** 不重要 / 过碎的内容可以往后放 */
          demote: ["**/appendix/**", "**/examples/**"],

          /** small 版本里直接排除 */
          exclude: ["**/assets/**", "**/images/**"],

          rawContent: true,
        }),
      ],
    }),
    sitemap(),
  ],
  outDir: "dist/astro",
  vite: {
    plugins: [tailwindcss()],
    // build: {
    //   emptyOutDir: false,
    // },
  },
});
