#
## jwt结构
> Header（头部）  
Payload（负载）  
Signature（签名）   
Header.Payload.Signature  
解析网站：[JSON WEB TOOL JWT](https://jwt.io/?spm=a2c6h.12873639.article-detail.9.35a94678mCk64c#debugger-io)
### Header（头部）
```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```
### Payload（负载）
> **registered claims(注册声明)**  
> 注册声明是预定义的声明，这些声明的名称和含义在规范中已经定义好。这些声明是可选的，但规范建议在适当的情况下使用。  
> iss: jwt签发者  
> sub: jwt所面向的用户  
> aud: 接收jwt的一方  
> exp: jwt的过期时间，这个过期时间必须要大于签发时间  
> nbf: 定义在什么时间之前，该jwt都是不可用的.  
> iat: jwt的签发时间  
> jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击
>   
> **Public claims(公共的声明)**   
> 公共声明是用户可以自定义的声明，但需要避免与注册声明冲突。这些声明可以用来传输用户定义的、公开的元数据。为了避免冲突，建议使用命名空间格式。
>   
> **Private claims(私人声明)**  
> 私有声明是专门为应用程序和用户之间传递信息而定义的声明。这些声明仅在特定上下文中有意义，通常由合作双方商定其名称和用途。  
> 例如：
> employeeId: 员工编号。
> department: 部门信息。
```json
{
  "iss": "https://your-auth-server.com",// 注册声明
  "sub": "1234567890",
  "aud": "https://your-api.com",
  "exp": 1716239022,
  "iat": 1516239022,
  "com.example.username": "johndoe",// 公共的声明
  "com.example.role": "admin",
  "org.company.department": "engineering",
  "private.employeeId": "E12345",// 私人声明
  "private.accessLevel": "high"
}
```
### Signature（签名） 
服务端通过Header.Payload与秘钥生成Signature，用于检测Header.Payload是否被修改

## go语言jwt代码模版
```go
package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	UserID   uint   `json:"userID"`
	Nickname string `json:"nickname"` // 用户名
	Role     int8   `json:"role"`     // 权限  1 管理员  2 普通用户
}

// CustomClaims 自定义声明
type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

// GenToken 创建 Token
func GenToken(payload JwtPayLoad, accessSecret string, expires int) (string, error) {
	claim := CustomClaims{
		JwtPayLoad: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(accessSecret))
}

// ParseToken 解析 Token
func ParseToken(tokenString string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
```

