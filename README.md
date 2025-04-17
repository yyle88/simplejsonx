[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/simplejsonx/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/simplejsonx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/simplejsonx)](https://pkg.go.dev/github.com/yyle88/simplejsonx)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/simplejsonx/master.svg)](https://coveralls.io/github/yyle88/simplejsonx?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/simplejsonx.svg)](https://github.com/yyle88/simplejsonx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/simplejsonx)](https://goreportcard.com/report/github.com/yyle88/simplejsonx)

# simplejsonx

`simplejsonx` is a generic-based JSON parsing package that depends on `github.com/bitly/go-simplejson`, enhancing type safety and flexibility in JSON processing. It requires at least Go 1.22 (which supports generics).

## CHINESE README

[中文说明](README.zh.md)

## Installation

```bash
go get github.com/yyle88/simplejsonx
```

## Usage Example

### 1. Reading and Loading JSON Data

First, use `simplejsonx.Load` to read and load JSON data.

```go
package main

import (
	"fmt"
	"github.com/yyle88/simplejsonx"
	"log"
)

func main() {
	// Sample JSON data
	data := []byte(`{"name": "yyle88", "age": 18, "is_rich": true}`)

	// Load the JSON data
	simpleJson, err := simplejsonx.Load(data)
	if err != nil {
		log.Fatalf("Error loading JSON: %v", err)
	}

	// Extract fields
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

	// Output the extracted values
	fmt.Println("name:", name, "age:", age, "rich:", isRich)
}
```

### 2. Extracting JSON Fields

Use `Extract` to extract JSON fields and ensure the correct type is returned.

```go
res, err := simplejsonx.Extract[int](simpleJson, "age")
if err != nil {
    log.Fatalf("Error extracting 'age': %v", err)
}
fmt.Println("Age:", res)  // Output: 18
```

### 3. Resolving JSON Values

Use the `Resolve` method to obtain the JSON field data in the specified type.

```go
simpleJson, err := simplejsonx.Load([]byte(`{"height": 175, "weight": 80}`))
if err != nil {
	log.Fatalf("Error loading JSON: %v", err)
}

height, err := simplejsonx.Resolve[int64](simpleJson.Get("height"))
if err != nil {
	log.Fatalf("Error resolving 'height': %v", err)
}

fmt.Println("Height:", height)  // Output: 175
```

### 4. Using `Inspect` to Extract Fields

The `Inspect` method is similar to `Extract`, but if the key does not exist, it returns the zero value without causing an error. It only returns an error in the case of invalid arguments or other issues. When using this method, you need to handle the returned zero value and errors correctly.

```go
simpleJson, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": 18}`))
if err != nil {
	log.Fatalf("Error loading JSON: %v", err)
}

// Use Inspect to extract fields. If the key does not exist, it returns the zero value (e.g., empty string, zero integer).
name, err := simplejsonx.Inspect[string](simpleJson, "name")
if err != nil {
    log.Fatalf("Error inspecting 'name': %v", err)
}
fmt.Println("Name:", name)  // Output: yyle88

// Extract a non-existent field, which returns the zero value (e.g., empty string).
address, err := simplejsonx.Inspect[string](simpleJson, "address")
if err != nil {
    log.Fatalf("Error inspecting 'address': %v", err)
}
fmt.Println("Address:", address)  // Output: Empty string (zero value)
```

---

## License

`simplejsonx` is open-source and released under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `simplejsonx`!** 🎉

Give me stars. Thank you!!!

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/simplejsonx.svg?variant=adaptive)](https://starchart.cc/yyle88/simplejsonx)
