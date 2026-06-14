# Shortlink Go

基于 Go + Gin 的短链接服务。

## 功能

- 生成短链：`POST /shorten` 提交长链接，返回短码
- 访问跳转：`GET /:shortcode` 自动 301 跳转到原链接

## 快速开始

### 运行

`bash
git clone https://github.com/Signal-zxh/shortlink-go.git
cd shortlink-go
go mod tidy
go run main.go
`

### 测试

生成短链：

`bash
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url":"https://go.dev"}'
`

访问短链：浏览器打开 `http://localhost:8080/xxx`（把 xxx 换成返回的短码）

## 技术栈

- Go 1.21+
- Gin
- base58

## License

MIT