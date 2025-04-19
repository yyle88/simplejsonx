package main

import (
	"fmt"

	"github.com/yyle88/simplejsonx/sure/simplejsonm"
)

func main() {
	// Sample JSON data
	data := []byte(`{"name": "yyle88", "age": 18, "is_rich": true}`)

	// Must Load the JSON data
	simpleJson := simplejsonm.Load(data)

	// Must Extract fields
	name := simplejsonm.Extract[string](simpleJson, "name")
	age := simplejsonm.Extract[int](simpleJson, "age")
	isRich := simplejsonm.Extract[bool](simpleJson, "is_rich")

	// Output the extracted values
	fmt.Println("name:", name, "age:", age, "rich:", isRich)
}
