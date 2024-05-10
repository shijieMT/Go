# go-zero 使用文档
## 安装环境
### 安装goctl
```shell
go install github.com/zeromicro/go-zero/tools/goctl@latest
```
### 安装protoc
```shell
goctl env check --install --verbose --force
go get -u github.com/zeromicro/go-zero@latest
```
### 创建一个api服务
> 微服务多的话，可以建一个apps文件夹，里面放所有服务
> 这里直接在user服务下创建 user/api/*
```shell
goctl api new api
```
### 修改user/api/internal/logic/apilogic.go文件
```go
func (l *ApiLogic) Api(req *types.Request) (resp *types.Response, err error) {
  // todo: add your logic here and delete this line

  return &types.Response{Message: "玛奇朵"}, nil
}
```
### 添加依赖并运行
> go mod tidy  
> Go 会自动处理 go.mod 和 go.sum 文件，确保它们反映了项目的最新依赖状态。
```shell
go mod tidy
```
