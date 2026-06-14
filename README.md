# Shortlink Go

基于 Go + Gin + MySQL 的短链接服务。

## 功能

- 生成短链：`POST /shorten` 提交长链接，返回短码
- 访问跳转：`GET /:shortcode` 自动 301 跳转到原链接
- 访问统计：`GET /stats/:shortcode` 查看短链被访问次数
- 数据持久化：短码映射关系存储在 MySQL 中，服务重启不丢失

## 快速开始

### 前置条件

- Go 1.21+
- Docker（用于运行 MySQL）

### 启动 MySQL

`bash
docker run -d \
  --name shortlink-mysql \
  -e MYSQL_ROOT_PASSWORD=123456 \
  -e MYSQL_DATABASE=shortlink \
  -p 3306:3306 \
  -v mysql-data:/var/lib/mysql \
  mysql:8.0
`

### 运行服务

`bash
git clone https://github.com/Signal-zxh/shortlink-go.git
cd shortlink-go
go mod tidy
go run main.go
`

### 测试

`bash
# 生成短链
curl -X POST http://localhost:8081/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://go.dev"}'

# 访问短链（用返回的 short_code）
# 浏览器打开 http://localhost:8081/r
`
# 查看统计
curl http://localhost:8081/stats/r


## 技术栈

- Go 1.21+
- Gin
- MySQL 8.0
- base58

## License

MIT