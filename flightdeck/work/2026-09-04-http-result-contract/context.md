# Context

- Docs 拥有文档集、导入、评论与授权等业务 DTO 和错误语义；Foundation 只拥有 catalog/operation schema、生成器和 Runtime。
- 当前 7 个 operation 只是已声明基线，必须以 canonical OpenAPI 为准验证是否覆盖全部 API。
- 本地验证通过 Workspace CLI 显式选择 Docs/Foundation/Identity/Asset checkout；本地数据库均为可丢弃测试数据。
- 本阶段不创建 tag、Release 或镜像。

