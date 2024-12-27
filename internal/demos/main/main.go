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
