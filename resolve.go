// Package simplejsonx: Type-safe JSON parsing engine with Go generics support
// Built on github.com/bitly/go-simplejson with enhanced compile-time type-safe checking
// Provides Extract, Inspect, Resolve functions enabling safe conversion to 13+ types
// Supports primitives, lists, maps, and nested JSON objects with zero-cost type checking
// Features code-generated exception handling variants in must/omit/soft modes
//
// simplejsonx: 类型安全的 JSON 解析引擎，支持 Go 泛型
// 基于 github.com/bitly/go-simplejson 构建，具有增强的编译时类型安全性
// 提供 Extract、Inspect、Resolve 函数，实现安全转换到 13+ 种类型
// 支持基础类型、数组、映射和嵌套 JSON 对象，具有零成本类型检查
// 具有代码生成的错误处理变体，包含 must/omit/soft 模式
package simplejsonx

import (
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"github.com/yyle88/simplejsonx/internal/utils"
)

// Extract retrieves and parses the value at the specified key into the target type
// Returns errors when the key is missing, when type conversion fails
// Supports automatic type conversion based on generic type parameters
//
// Extract 检索指定键的值并解析成目标类型
// 当键缺失或类型转换失败时返回错误
// 支持基于泛型类型参数的自动类型转换
func Extract[T any](object *simplejson.Json, key string) (T, error) {
	if object == nil {
		return utils.Zero[T](), errors.New("parameter object is missing")
	}
	if key == "" {
		return utils.Zero[T](), errors.New("parameter key is missing")
	}
	return Resolve[T](object.Get(key))
}

// Inspect retrieves and parses the value at the specified key when present
// Returns zero value without errors when key is absent, errors on conversion failures
// Provides lenient access patterns that accept fields in JSON data
//
// Inspect 检索并解析指定键的值（当键存在时）
// 当键不存在时返回零值且不报错，仅在转换失败时返回错误
// 提供宽松的访问模式，允许 JSON 结构中存在可选字段
func Inspect[T any](object *simplejson.Json, key string) (T, error) {
	if object == nil {
		return utils.Zero[T](), errors.New("parameter object is missing")
	}
	if key == "" {
		return utils.Zero[T](), errors.New("parameter key is missing")
	}
	value, exist := object.CheckGet(key)
	if !exist {
		return utils.Zero[T](), nil
	}
	return Resolve[T](value)
}

// Resolve extracts and converts JSON value into the target type
// Supports comprehensive type conversion via simplejson.Json methods
// Handles primitives (int, int64, float64, string, uint64, bool)
// Handles arrays ([]string, []interface{}, []*simplejson.Json)
// Handles complex types (map[string]interface{}, []byte, *simplejson.Json)
//
// Resolve 提取 JSON 值并转换成目标类型
// 支持使用 simplejson.Json 方法进行全面的类型转换
// 处理基础类型（int、int64、float64、string、uint64、bool）
// 处理数组类型（[]string、[]interface{}、[]*simplejson.Json）
// 处理复杂类型（map[string]interface{}、[]byte、*simplejson.Json）
func Resolve[T any](object *simplejson.Json) (T, error) {
	if object == nil {
		return utils.Zero[T](), errors.New("parameter object is missing")
	}
	switch zero := utils.Zero[T](); any(zero).(type) {
	case int:
		res, err := object.Int()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to int")
		}
		return any(res).(T), nil
	case int64:
		res, err := object.Int64()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to int64")
		}
		return any(res).(T), nil
	case float64:
		res, err := object.Float64()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to float64")
		}
		return any(res).(T), nil
	case string:
		res, err := object.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		return any(res).(T), nil
	case uint64:
		res, err := object.Uint64()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to uint64")
		}
		return any(res).(T), nil
	case bool:
		res, err := object.Bool()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to bool")
		}
		return any(res).(T), nil
	case []string:
		res, err := object.StringArray()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to []string")
		}
		return any(res).(T), nil
	case []interface{}:
		res, err := object.Array()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to []interface{}")
		}
		return any(res).(T), nil
	case map[string]interface{}:
		res, err := object.Map()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to map[string]interface{}")
		}
		return any(res).(T), nil
	case []byte:
		res, err := object.Bytes()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to []byte")
		}
		return any(res).(T), nil
	case *simplejson.Json:
		return any(object).(T), nil
	case []*simplejson.Json:
		elements, err := object.Array()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to []interface{}")
		}
		return any(List(elements)).(T), nil
	default:
		return zero, errors.Errorf("unsupported generic type: %T. unable to resolve JSON value.", zero)
	}
}

// GetList retrieves JSON list at the specified key as simplejson.Json objects slice
// Converts each list element into distinct simplejson.Json objects
// Returns errors when key is missing, when value is not a list
//
// GetList 检索指定键的 JSON 数组并转换成 simplejson.Json 对象切片
// 将每个数组元素转换成独立的 simplejson.Json 包装对象
// 当键缺失或值不是数组时返回错误
func GetList(object *simplejson.Json, key string) (objects []*simplejson.Json, err error) {
	if object == nil {
		return nil, errors.New("parameter object is missing")
	}
	if key == "" {
		return nil, errors.New("parameter key is missing")
	}
	elements, err := object.Get(key).Array()
	if err != nil {
		return objects, errors.WithMessage(err, "unable to get list")
	}
	return List(elements), nil
}

// Inquire queries JSON object at the specified key with tri-state result pattern
// Returns parsed value, existence boolean, and possible conversion errors
// Distinguishes between missing keys and type conversion failures
//
// Inquire 查询 JSON 对象中指定键的值，使用三态结果模式
// 返回解析后的值、存在性布尔值和可能的转换错误
// 区分缺失的键和类型转换失败的情况
func Inquire[T any](object *simplejson.Json, key string) (T, bool, error) {
	if object == nil {
		return utils.Zero[T](), false, errors.New("parameter object is missing")
	}
	if key == "" {
		return utils.Zero[T](), false, errors.New("parameter key is missing")
	}
	value, exist := object.CheckGet(key)
	if !exist {
		return utils.Zero[T](), false, nil
	}
	res, err := Resolve[T](value)
	if err != nil {
		return utils.Zero[T](), false, errors.WithMessage(err, "unable to resolve JSON value")
	}
	return res, true, nil
}

// Attempt retrieves and parses value at the specified key with success indication
// Returns parsed value and boolean indicating complete operation success
// Without notification returns zero value and false on failures without exposing errors
//
// Attempt 检索并解析指定键的值，附带成功指示器
// 返回解析后的值和表示整体操作成功的布尔值
// 在任何失败情况下静默返回零值和 false，不暴露错误
func Attempt[T any](object *simplejson.Json, key string) (T, bool) {
	if object == nil {
		return utils.Zero[T](), false
	}
	if key == "" {
		return utils.Zero[T](), false
	}
	value, exist := object.CheckGet(key)
	if !exist {
		return utils.Zero[T](), false
	}
	res, err := Resolve[T](value)
	if err != nil {
		return utils.Zero[T](), false
	}
	return res, true
}

// Explore navigates nested JSON structure via dot-separated path notation
// Traverses multiple levels using path like "user.profile.name"
// Returns parsed value, existence boolean, and possible conversion errors
//
// Explore 通过点分隔路径表示法导航嵌套 JSON 结构
// 使用类似 "user.profile.name" 的路径遍历多层级结构
// 返回解析后的值、存在性布尔值和可能的转换错误
func Explore[T any](object *simplejson.Json, path string) (T, bool, error) {
	if object == nil {
		return utils.Zero[T](), false, errors.New("parameter object is missing")
	}
	if path == "" {
		return utils.Zero[T](), false, errors.New("parameter path is missing")
	}
	value := object
	var exist bool
	for _, key := range strings.Split(path, ".") {
		value, exist = value.CheckGet(key)
		if !exist {
			return utils.Zero[T](), false, nil
		}
	}
	res, err := Resolve[T](value)
	if err != nil {
		return utils.Zero[T](), false, errors.WithMessage(err, "unable to resolve JSON value")
	}
	return res, true, nil
}
