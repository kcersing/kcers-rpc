# kcers
Gym Sysytem

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOPROXY=https://mirrors.aliyun.com/goproxy/


go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
go install github.com/cloudwego/hertz/cmd/hz@latest
go install github.com/cloudwego/cwgo@latest
GO111MODULE=on go install github.com/cloudwego/thriftgo@latest

## 技术栈

| 功能      | 实现              |
|---------|-----------------|
| HTTP 框架 | Hertz           |
| RPC 框架  | Kitex |
| 数据库     | MySQL、Redis     |
| 数据库orm  | ent             |
| 消息队列    | RabbitMQ        |
| 身份鉴权    | JWT             |
| 访问控制    | Casbin          |
| 服务与配置中心 | Consul         |
| 链路追踪    | Jaeger         |
| 限流熔断    | Sentinel       |
| API 文档  | Swagger         |
| 身份鉴权    | JWT 认证          |
| 搜索      | Elasticsearch   |
| 指标监控    | Prometheus      |
| 服务治理    | OpenTelemetry  |
| 对象存储    | MinIO           |
| 图像识别    | 百度 OCR         |
| CI      | GitHub Actions  |
| 代码生成    | hz-thrift       |
