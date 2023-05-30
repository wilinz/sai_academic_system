# Login system

## 编译
```shell
 $env:GOOS="linux" ; $env:GOARCH="arm64" ; go build -o loginsystem cmd/main.go
```
```shell
 $env:GOOS="linux" ; $env:GOARCH="amd64" ; go build -o loginsystem  cmd/main.go
```

## 主函数在 cmd/main.go
启动前请修改 db 文件夹里面的数据库配置，如果要发送邮件 请配置 tools/sendemail.go 里面的邮箱信息