[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/simplejsonx/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/simplejsonx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/simplejsonx)](https://pkg.go.dev/github.com/yyle88/simplejsonx)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/simplejsonx/main.svg)](https://coveralls.io/github/yyle88/simplejsonx?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/simplejsonx.svg)](https://github.com/yyle88/simplejsonx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/simplejsonx)](https://goreportcard.com/report/github.com/yyle88/simplejsonx)

# simplejsonx

`simplejsonx` 是个基于泛型的 JSON 解析库，依赖于 `github.com/bitly/go-simplejson`，通过增强类型安全性和灵活性来优化 JSON 处理，同时要求至少使用 Go 1.22 版本（需要支持泛型）。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 安装

```bash
go get github.com/yyle88/simplejsonx
```

## 使用示例

### 1. 基础 JSON 解析与错误处理

此示例展示如何加载 JSON 数据并使用错误处理提取类型化字段。

```go
package main

import (
	"fmt"
	"log"

	"github.com/yyle88/simplejsonx"
)

func main() {
	// 示例 JSON 数据
	data := []byte(`{"name": "yyle88", "age": 18, "is_rich": true}`)

	// 加载 JSON 数据
	object, err := simplejsonx.Load(data)
	if err != nil {
		log.Fatalf("Error loading JSON: %v", err)
	}

	// 提取字段
	name, err := simplejsonx.Extract[string](object, "name")
	if err != nil {
		log.Fatalf("Error extracting 'name': %v", err)
	}

	age, err := simplejsonx.Extract[int](object, "age")
	if err != nil {
		log.Fatalf("Error extracting 'age': %v", err)
	}

	isRich, err := simplejsonx.Extract[bool](object, "is_rich")
	if err != nil {
		log.Fatalf("Error extracting 'is_rich': %v", err)
	}

	// 输出提取的值
	fmt.Println("name:", name, "age:", age, "rich:", isRich)  // 输出: name: yyle88 age: 18 rich: true
}
```

⬆️ **源码：** [源码](internal/demos/demo1x/main.go)

### 2. Must 风格 API（错误时 Panic）

此示例展示 must 风格 API，在遇到错误时会 panic，适用于不期望出错的场景。

```go
package main

import (
	"fmt"

	"github.com/yyle88/simplejsonx/sure/simplejsonm"
)

func main() {
	// 示例 JSON 数据
	data := []byte(`{"name": "yyle88", "age": 18, "is_rich": true}`)

	// 必须成功加载 JSON 数据（失败会 panic）
	object := simplejsonm.Load(data)

	// 必须成功提取字段（失败会 panic）
	name := simplejsonm.Extract[string](object, "name")
	age := simplejsonm.Extract[int](object, "age")
	isRich := simplejsonm.Extract[bool](object, "is_rich")

	// 输出提取的值
	fmt.Println("name:", name, "age:", age, "rich:", isRich)  // 输出: name: yyle88 age: 18 rich: true
}
```

⬆️ **源码：** [源码](internal/demos/demo2x/main.go)

## 功能示例

### 提取 JSON 字段

**基础字段提取：**
```go
res, err := simplejsonx.Extract[int](object, "age")
if err != nil {
    log.Fatalf("Error extracting 'age': %v", err)
}
fmt.Println("Age:", res)  // 输出: 18
```

### 解析 JSON 值

**使用 Resolve 获取类型化值：**
```go
object, err := simplejsonx.Load([]byte(`{"height": 175, "weight": 80}`))
if err != nil {
	log.Fatalf("Error loading JSON: %v", err)
}

height, err := simplejsonx.Resolve[int64](object.Get("height"))
if err != nil {
	log.Fatalf("Error resolving 'height': %v", err)
}
fmt.Println("Height:", height)  // 输出: 175
```

### 使用 Inspect 处理可选字段

**Inspect 在字段缺失时返回零值：**
```go
object, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": 18}`))
if err != nil {
	log.Fatalf("Error loading JSON: %v", err)
}

name, err := simplejsonx.Inspect[string](object, "name")
if err != nil {
    log.Fatalf("Error inspecting 'name': %v", err)
}
fmt.Println("Name:", name)  // 输出: yyle88

address, err := simplejsonx.Inspect[string](object, "address")
if err != nil {
    log.Fatalf("Error inspecting 'address': %v", err)
}
fmt.Println("Address:", address)  // 输出: 空字符串（零值）
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yyle88/simplejsonx.svg?variant=adaptive)](https://starchart.cc/yyle88/simplejsonx)
