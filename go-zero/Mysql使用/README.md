
# go-zero & Mysql
## 原生操作
todo
### 
## 使用 goctl 生成 model 代码
[官方文档](https://go-zero.dev/docs/tasks/cli/mysql)
### model/user.sql
```sql
CREATE TABLE user
(
    id        bigint AUTO_INCREMENT,
    username  varchar(36) NOT NULL,
    password  varchar(64) default '',
    UNIQUE name_index (username),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci;
```
### MySQL导入 sql文件
```shell
create database go_zero;
use go_zero;
source user.sql
```
goctl model mysql ddl --src user.sql --dir .
### 创建api文件
```api
type LoginRequest {
	UserName string `json:"username"`
	Password string `json:"password"`
}

@server (
	prefix: /api/users
)
service users {
	@handler login
	post /login (LoginRequest) returns (string)
}
```
### 生成代码
```shell
goctl api go -api .\user.api -dir .
```
## 生成api代码
/api/user.api
```api
type LoginRequest {
UserName string `json:"username"`
Password string `json:"password"`
}

@server (
prefix: /api/users
)
service users {
@handler login
post /login (LoginRequest) returns (string)
}
```
生成代码
```shell
goctl api go -api user.api -dir .
```
users.yaml 配置MySQL
```yaml
Name: users
Host: 0.0.0.0
Port: 8888
Mysql:
  DataSource: root:135789@tcp(127.0.0.1:3306)/go_zero?charset=utf8mb4&parseTime=True&loc=Local
```
internal/config/config.go
```go
type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
}
```
internal/svc/servicecontext.go
```go
type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(mysqlConn),
	}
}
```
internal/logic/loginlogic.go
```go
func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
// todo: add your logic here and delete this line
rep, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
Username: req.UserName,
Password: req.Password,
})
if err != nil {
return "", err
}
fmt.Println("返回值", rep)
return "插入成功", nil
}
```
