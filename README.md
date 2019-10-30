# vuecms api

## 环境需求
```sh
$ go version
go version go1.13.1 linux/amd64
$ go env -w GOPROXY=https://goproxy.cn,direct
$ go get -u -insecure github.com/swaggo/swag/cmd/swag
$ go get -u -insecure github.com/jirfag/go-queryset/cmd/goqueryset
```

## 开发步骤s
```sh
$ edit config/app.toml #修改数据库配置
$ swag init
$ go run main.go
```
