package main

import (
	"fmt"

	"github.com/yyle88/simplejsonx/sure/simplejsonm"
)

func main() {
	// Sample JSON data
	// 示例 JSON 数据
	data := []byte(`{"name": "yyle88", "age": 18, "is_rich": true}`)

	// Must Load the JSON data (panics on error)
	// 必须成功加载 JSON 数据（失败会 panic）
	object := simplejsonm.Load(data)

	// Must Extract fields (panics on error)
	// 必须成功提取字段（失败会 panic）
	name := simplejsonm.Extract[string](object, "name")
	age := simplejsonm.Extract[int](object, "age")
	isRich := simplejsonm.Extract[bool](object, "is_rich")

	// Output the extracted values
	// 输出提取的值
	fmt.Println("name:", name, "age:", age, "rich:", isRich)
}
