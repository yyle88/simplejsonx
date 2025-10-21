package simplejsonx

import (
	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
)

// Load creates simplejson.Json instance from raw JSON bytes
// Parses byte data into structured JSON object representation
// Returns errors with context when JSON parsing fails
//
// Load 从原始 JSON 字节创建 simplejson.Json 实例
// 将字节数据解析成结构化的 JSON 对象表示
// 当 JSON 解析失败时返回带上下文的错误
func Load(data []byte) (object *simplejson.Json, err error) {
	object, err = simplejson.NewJson(data)
	if err != nil {
		return simplejson.New(), errors.WithMessage(err, "unable to parse JSON")
	}
	return object, nil
}

// Wrap creates simplejson.Json instance wrapping given Go values
// Encapsulates Go data structures into JSON object representation
// Enables uniform JSON operations on native Go values
//
// Wrap 创建包装任意 Go 值的 simplejson.Json 实例
// 将任何 Go 数据结构封装成 JSON 对象表示
// 在原生 Go 值上实现统一的 JSON 操作
func Wrap(value interface{}) (object *simplejson.Json) {
	object = simplejson.New()
	object.SetPath([]string{}, value)
	return object
}

// List converts slice of interface{} values into simplejson.Json objects slice
// Wraps each element as distinct JSON object instances
// Provides batch conversion enabling uniform processing of heterogeneous data
//
// List 将 interface{} 值切片转换成 simplejson.Json 对象切片
// 独立包装每个元素，创建独立的 JSON 对象实例
// 提供批量转换，实现对异构数据的统一处理
func List(elements []interface{}) (objects []*simplejson.Json) {
	objects = make([]*simplejson.Json, 0, len(elements))
	for _, elem := range elements {
		objects = append(objects, Wrap(elem))
	}
	return objects
}
