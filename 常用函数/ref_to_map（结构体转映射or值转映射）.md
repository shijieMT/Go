# ref_to_map
## todo 验证传递指针是否可以成功转换
## ref_to_map.go
```go
package test
import "reflect"

func RefToMap(data any, tag string) map[string]any {
	maps := map[string]any{}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data) 
	// 循环处理所有数据
	for i := 0; i < t.NumField(); i++ {
		// 拿到第i个数据的类型
		field := t.Field(i)
		getTag, ok := field.Tag.Lookup(tag)
		// 非tag类型数据，continue
		if !ok { 
			continue
		}
		// 拿到第i个数据的值
		val := v.Field(i)
		// 如果是0值，continue
		if val.IsZero() { 
			continue
		}
		// 1. 数据的类型为 struct，递归处理
		if field.Type.Kind() == reflect.Struct { 
			newMaps := RefToMap(val.Interface(), tag)
			maps[getTag] = newMaps
			continue
		}
		// 2. 数据的类型为 Ptr，递归处理Ptr指向的内容
		if field.Type.Kind() == reflect.Ptr { 
			if field.Type.Elem().Kind() == reflect.Struct {
				newMaps := RefToMap(val.Elem().Interface(), tag)
				maps[getTag] = newMaps
				continue
			}
			maps[getTag] = val.Elem().Interface()
			continue
		}
		// 3. 普通类型数据
		maps[getTag] = val.Interface()

	}
	return maps
}
```
## my_test.go
```go
package test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
	Phone   *string `json:"phone"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func main() {
	phone := "123456789"
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
		Phone: &phone,
	}

	result := RefToMap(person, "json")
	fmt.Printf("%+v\n", result)
}

func TestName(t *testing.T) {
	main()
}
```