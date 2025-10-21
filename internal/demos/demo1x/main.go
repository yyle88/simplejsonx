package main

import (
	"fmt"
	"log"

	"github.com/yyle88/simplejsonx"
)

func main() {
	// Sample JSON data
	// 示例 JSON 数据
	data := []byte(`{"name": "yyle88", "age": 18, "is_rich": true}`)

	// Load the JSON data
	// 加载 JSON 数据
	object, err := simplejsonx.Load(data)
	if err != nil {
		log.Fatalf("Error loading JSON: %v", err)
	}

	// Extract fields
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

	// Output the extracted values
	// 输出提取的值
	fmt.Println("name:", name, "age:", age, "rich:", isRich)
}
