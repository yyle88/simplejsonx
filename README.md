[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/simplejsonx/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/simplejsonx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/simplejsonx)](https://pkg.go.dev/github.com/yyle88/simplejsonx)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/simplejsonx/main.svg)](https://coveralls.io/github/yyle88/simplejsonx?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/simplejsonx.svg)](https://github.com/yyle88/simplejsonx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/simplejsonx)](https://goreportcard.com/report/github.com/yyle88/simplejsonx)

# simplejsonx

`simplejsonx` is a generic-based JSON parsing package that depends on `github.com/bitly/go-simplejson`, enabling type-safe and flexible JSON processing. It requires at least Go 1.22 (which supports generics).

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Installation

```bash
go get github.com/yyle88/simplejsonx
```

## Usage Example

### 1. Basic JSON Parsing with Error Handling

This example demonstrates how to load JSON data and extract typed fields with error handling.

```go
package main

import (
	"fmt"
	"log"

	"github.com/yyle88/simplejsonx"
)

func main() {
	// Sample JSON data
	data := []byte(`{"name": "yyle88", "age": 18, "is_rich": true}`)

	// Load the JSON data
	object, err := simplejsonx.Load(data)
	if err != nil {
		log.Fatalf("Error loading JSON: %v", err)
	}

	// Extract fields
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

	// Output the extracted values
	fmt.Println("name:", name, "age:", age, "rich:", isRich)  // Output: name: yyle88 age: 18 rich: true
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### 2. Must-Style API (Panic on Error)

This example shows the must-style API that panics on errors, useful when errors are unexpected.

```go
package main

import (
	"fmt"

	"github.com/yyle88/simplejsonx/sure/simplejsonm"
)

func main() {
	// Sample JSON data
	data := []byte(`{"name": "yyle88", "age": 18, "is_rich": true}`)

	// Must Load the JSON data (panics on error)
	object := simplejsonm.Load(data)

	// Must Extract fields (panics on error)
	name := simplejsonm.Extract[string](object, "name")
	age := simplejsonm.Extract[int](object, "age")
	isRich := simplejsonm.Extract[bool](object, "is_rich")

	// Output the extracted values
	fmt.Println("name:", name, "age:", age, "rich:", isRich)  // Output: name: yyle88 age: 18 rich: true
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

## Examples

### Extracting JSON Fields

**Basic field extraction:**
```go
res, err := simplejsonx.Extract[int](object, "age")
if err != nil {
    log.Fatalf("Error extracting 'age': %v", err)
}
fmt.Println("Age:", res)  // Output: 18
```

### Resolving JSON Values

**Using Resolve to get typed values:**
```go
object, err := simplejsonx.Load([]byte(`{"height": 175, "weight": 80}`))
if err != nil {
	log.Fatalf("Error loading JSON: %v", err)
}

height, err := simplejsonx.Resolve[int64](object.Get("height"))
if err != nil {
	log.Fatalf("Error resolving 'height': %v", err)
}
fmt.Println("Height:", height)  // Output: 175
```

### Using Inspect for Optional Fields

**Inspect returns zero value when field is missing:**
```go
object, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": 18}`))
if err != nil {
	log.Fatalf("Error loading JSON: %v", err)
}

name, err := simplejsonx.Inspect[string](object, "name")
if err != nil {
    log.Fatalf("Error inspecting 'name': %v", err)
}
fmt.Println("Name:", name)  // Output: yyle88

address, err := simplejsonx.Inspect[string](object, "address")
if err != nil {
    log.Fatalf("Error inspecting 'address': %v", err)
}
fmt.Println("Address:", address)  // Output: Blank string (zero value)
```

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/simplejsonx.svg?variant=adaptive)](https://starchart.cc/yyle88/simplejsonx)
