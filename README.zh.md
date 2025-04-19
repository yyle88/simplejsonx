# simplejsonx

`simplejsonx` 是个基于泛型的 JSON 解析库，依赖于 `github.com/bitly/go-simplejson`，通过增强类型安全性和灵活性来优化 JSON 处理，同时要求至少使用 Go 1.22 版本（需要支持泛型）。

## 英文文档

[ENGLISH README](README.md)

## 安装

```bash
go get github.com/yyle88/simplejsonx
```

## 使用示例

### 1. 读取和加载 JSON 数据

首先，通过 `simplejsonx.Load` 读取和加载 JSON 数据。

```go
package main

import (
	"fmt"
	"github.com/yyle88/simplejsonx"
	"log"
)

func main() {
	// 示例 JSON 数据
	data := []byte(`{"name": "yyle88", "age": 18, "is_rich": true}`)

	// 加载 JSON 数据
	simpleJson, err := simplejsonx.Load(data)
	if err != nil {
		log.Fatalf("Error loading JSON: %v", err)
	}

	// 提取字段
	name, err := simplejsonx.Extract[string](simpleJson, "name")
	if err != nil {
		log.Fatalf("Error extracting 'name': %v", err)
	}

	age, err := simplejsonx.Extract[int](simpleJson, "age")
	if err != nil {
		log.Fatalf("Error extracting 'age': %v", err)
	}

	isRich, err := simplejsonx.Extract[bool](simpleJson, "is_rich")
	if err != nil {
		log.Fatalf("Error extracting 'is_rich': %v", err)
	}

	// 输出提取的值
	fmt.Println("name:", name, "age:", age, "rich:", isRich)
}
```

### 2. 提取 JSON 字段

使用 `Extract` 提取 JSON 字段，并确保返回正确的类型。

```go
res, err := simplejsonx.Extract[int](simpleJson, "age")
if err != nil {
    log.Fatalf("Error extracting 'age': %v", err)
}
fmt.Println("Age:", res)  // 输出: 18
```

### 3. 提取 JSON 的值

通过 `Resolve` 方法，获取指定类型的 JSON 字段数据。

```go
simpleJson, err := simplejsonx.Load([]byte(`{"height": 175, "weight": 80}`))
if err != nil {
	log.Fatalf("Error loading JSON: %v", err)
}

height, err := simplejsonx.Resolve[int64](simpleJson.Get("height"))
if err != nil {
	log.Fatalf("Error resolving 'height': %v", err)
}

fmt.Println("Height:", height)  // 输出: 175
```

### 4. 使用 `Inspect` 提取字段

`Inspect` 方法与 `Extract` 类似，但如果键不存在，则返回零值，并且不会产生错误。它返回的错误仅在参数或其他问题时发生。使用时，您需要正确处理返回的零值和错误。

```go
simpleJson, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": 18}`))
if err != nil {
	log.Fatalf("Error loading JSON: %v", err)
}

// 使用 Inspect 提取字段，如果键不存在，则返回零值（例如，空字符串、零整数等）
name, err := simplejsonx.Inspect[string](simpleJson, "name")
if err != nil {
    log.Fatalf("Error inspecting 'name': %v", err)
}
fmt.Println("Name:", name)  // 输出: yyle88

// 提取一个不存在的字段，返回零值（例如，空字符串）
address, err := simplejsonx.Inspect[string](simpleJson, "address")
if err != nil {
    log.Fatalf("Error inspecting 'address': %v", err)
}
fmt.Println("Address:", address)  // 输出: 空字符串（零值）
```

---

## 许可

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
