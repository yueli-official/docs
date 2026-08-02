# User 主体合同消费者迁移

## Status

Finished

## Result

- Docs 授权适配器只把 `subject_kind=user` 识别为用户，把 `subject_kind=client` 识别为服务。
- 审计 actor 使用声明的主体类型；本地 bootstrap 管理员改用 Identity Public User Key。
- 测试 principal 由真实 JWT 签发与验签构造。

## Verification

- `go test ./... -count=1 -timeout 240s`
- `go vet ./...`

全局合同和发布顺序记录在 Identity User 合同 Work 中。
