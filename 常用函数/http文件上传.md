
# net/http 文件上传
## 文件上传
```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // 解析表单数据
	
    err := r.ParseMultipartForm(10 << 20) // 这行代码解析表单数据，允许最大 10MB 的文件上传。
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    // 获取上传的文件
    file, fileHeader, err := r.FormFile("file") // 这行代码从表单中读取文件字段，返回文件接口、文件头信息和错误。
    if err != nil {
        http.Error(w, "Unable to get the file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // 打印文件信息
    fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
    fmt.Printf("File Size: %+v\n", fileHeader.Size)
    fmt.Printf("MIME Header: %+v\n", fileHeader.Header)

    // 创建目标文件
    dst, err := os.Create("/path/to/save/" + fileHeader.Filename)
    if err != nil {
        http.Error(w, "Unable to create the file", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    // 将上传的文件内容复制到目标文件
    _, err = io.Copy(dst, file)
    if err != nil {
        http.Error(w, "Unable to save the file", http.StatusInternalServerError)
        return
    }

    // 返回成功响应
    fmt.Fprintf(w, "File uploaded successfully: %s\n", fileHeader.Filename)
}

func main() {
    http.HandleFunc("/upload", uploadHandler)
    http.ListenAndServe(":8080", nil)
}
```
## go-zero中文件上传（一）
- 项目结构：
```shell
your_project_name/
├── internal/
│   ├── handler/
│   │   └── upload_handler.go
│   ├── logic/
│   │   └── upload_logic.go
│   ├── svc/
│   │   └── service_context.go
│   └── types/
│       └── types.go
└── main.go
```
- 请求和响应类型
```go
package types

type UploadFileResponse struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}
```
- 业务逻辑层
```go
func (l *UploadLogic) UploadFile(file multipart.File, header *multipart.FileHeader) (*types.UploadFileResponse, error) {
    defer file.Close()

    // 创建目标文件
    dstPath := filepath.Join("/path/to/save", header.Filename)
    dst, err := os.Create(dstPath)
    if err != nil {
        return nil, err
    }
    defer dst.Close()

    // 将上传的文件内容复制到目标文件
    size, err := io.Copy(dst, file)
    if err != nil {
        return nil, err
    }

    // 返回响应
    return &types.UploadFileResponse{
        Filename: header.Filename,
        Size:     size,
    }, nil
}
```
- 处理器
```go
func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 解析上传的文件
        file, header, err := r.FormFile("file")
        if err != nil {
            httpx.Error(w, err)
            return
        }

        // 调用业务逻辑处理文件上传
        l := logic.NewUploadLogic(r.Context(), svcCtx)
        resp, err := l.UploadFile(file, header)
        if err != nil {
            httpx.Error(w, err)
        } else {
            httpx.OkJson(w, resp)
        }
    }
}
```
- 配置文件
```shell
Name: your_project_name
Host: 0.0.0.0
Port: 8080

RestConf:
  Host: 0.0.0.0
  Port: 8080
```
## go-zero中文件上传（二）在Handler中写业务逻辑
```go
func FileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		file, fileHead, err := r.FormFile("file")
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		// 读取文件数据
		fileData, _ := io.ReadAll(file)

		l := logic.NewFileLogic(r.Context(), svcCtx)
		resp, err := l.File(&req)
        // 添加各种需要回传的信息
		resp.Src = newFileModel.WebPath()
		
		response.Response(r, w, resp, err)
	}
}

```
