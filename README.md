# exam api

## 环境需求
```sh
$ go version
go version go1.12 linux/amd64
$ export GOPROXY=https://goproxy.io
$ go get -u github.com/swaggo/swag/cmd/swag
```

## 开发步骤
```sh
$ cp config/app.toml.example config/app.toml 
$ edit config/app.toml #修改数据库配置
$ swag init
$ go run main.go
```
