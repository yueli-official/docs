# Docs

Docs 是面向公开阅读的文档发布与导航上下文。它用 Document Collection 组织相关文档，用 Documentation Locale 表达语言，
用 Version Family 关联独立版本化文档集，并以 Logical Document Key 关联跨语言、跨版本的同一页面。

## Language

**Document Collection（文档集）**:
组织同一产品或主题文档、语言和版本的公开阅读单元。
_Avoid_: 独立语言站点、目录文件夹

**Documentation Locale（文档语言）**:
一个 Document Collection 明确启用的稳定 BCP 47 语言标识；它拥有面向读者的名称、方向、启用状态、顺序和默认语义。
_Avoid_: 作者自由输入字符串、独立文档集、界面语言

**Default Documentation Locale（默认文档语言）**:
一个 Document Collection 中恰好一个已启用的默认 Documentation Locale；新文档和缺失译文回退以它为准。
_Avoid_: 浏览器自动语言、全站固定语言

**Locale Derivation（语言派生）**:
在同一 Document Collection 内，以一个已有 Documentation Locale 的完整文档树建立另一语言的翻译草稿；派生文档使用新
Document ID、保留 Logical Document Key 与层级，之后由目标语言独立编辑。
_Avoid_: 自动翻译、覆盖已有译文、跨文档集复制、共享正文

**Version Family（版本族）**:
把同一产品或主题的多个 Versioned Document Collection 关联起来的稳定分组；它只表达版本关系，不拥有文档内容。
_Avoid_: 在一个文档集内部维护版本矩阵、Blog Series、目录文件夹

**Versioned Document Collection（版本化文档集）**:
Version Family 中只承载一个语义版本内容的 Document Collection；语言、权限、发布状态和文档树均由该文档集独立维护。
_Avoid_: Document Version 行、共享文档树、Git branch

**Semantic Documentation Version（文档语义版本）**:
Versioned Document Collection 在 Version Family 内唯一且不可变的 `major.minor.patch` 标识，例如 `2.1.0`；排序遵循
SemVer 数值语义，不使用字符串排序，也不保存 `v` 前缀。
_Avoid_: 展示标题、任意版本字符串、可修改 Label

**Collection Derivation（文档集派生）**:
从一个 Document Collection 克隆语言注册、文档树和展示元数据以创建新集合的事实；派生后两个集合独立演进。
_Avoid_: 实时同步、继承、不可变快照

**Logical Document Key（逻辑文档键）**:
同一逻辑页面跨语言、跨关联文档集共享的稳定身份，当前由 `translationKey` 承载；切换语言或版本族成员时通过它解析目标页面。
_Avoid_: Slug、Title、数据库 Document ID

**Document Badge（文档标记）**:
贴近文档标题展示的短文字或 Tabler 图标，用于表达 `v2`、`Beta`、`已废弃` 等轻量提示；它不改变发布、权限或路由语义。
_Avoid_: Document Status、Semantic Documentation Version、分类标签
