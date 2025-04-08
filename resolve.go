package simplejsonx

import (
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"github.com/yyle88/tern"
)

// Extract retrieves the value associated with the given key in the JSON object and parses it into the specified type.
func Extract[T any](simpleJson *simplejson.Json, key string) (T, error) {
	if simpleJson == nil {
		return tern.Zero[T](), errors.New("parameter simpleJson is missing")
	}
	if key == "" {
		return tern.Zero[T](), errors.New("parameter key is missing")
	}
	return Resolve[T](simpleJson.Get(key))
}

// Inspect retrieves the value of the given key from the JSON object if it exists, parsing it into the specified type, or returns the zero value if the key is missing.
func Inspect[T any](simpleJson *simplejson.Json, key string) (T, error) {
	if simpleJson == nil {
		return tern.Zero[T](), errors.New("parameter simpleJson is missing")
	}
	if key == "" {
		return tern.Zero[T](), errors.New("parameter key is missing")
	}
	value, exist := simpleJson.CheckGet(key)
	if !exist {
		return tern.Zero[T](), nil
	}
	return Resolve[T](value)
}

/*
Resolve extracts the value from the provided JSON and convert it to typed value.
Supports the following functions:
func (j *Json) Int() (int, error)
func (j *Json) Int64() (int64, error)
func (j *Json) Float64() (float64, error)
func (j *Json) String() (string, error)
func (j *Json) Uint64() (uint64, error)
func (j *Json) Bool() (bool, error)
func (j *Json) StringArray() ([]string, error)
func (j *Json) Array() ([]interface{}, error)
func (j *Json) Map() (map[string]interface{}, error)
func (j *Json) Bytes() ([]byte, error)
*/

// Resolve extracts and converts the value from the provided JSON object into the specified type.
func Resolve[T any](simpleJson *simplejson.Json) (T, error) {
	if simpleJson == nil {
		return tern.Zero[T](), errors.New("parameter simpleJson is missing")
	}
	switch zero := tern.Zero[T](); any(zero).(type) {
	case int:
		res, err := simpleJson.Int()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to int")
		}
		return any(res).(T), nil
	case int64:
		res, err := simpleJson.Int64()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to int64")
		}
		return any(res).(T), nil
	case float64:
		res, err := simpleJson.Float64()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to float64")
		}
		return any(res).(T), nil
	case string:
		res, err := simpleJson.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		return any(res).(T), nil
	case uint64:
		res, err := simpleJson.Uint64()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to uint64")
		}
		return any(res).(T), nil
	case bool:
		res, err := simpleJson.Bool()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to bool")
		}
		return any(res).(T), nil
	case []string:
		res, err := simpleJson.StringArray()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to []string")
		}
		return any(res).(T), nil
	case []interface{}:
		res, err := simpleJson.Array()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to []interface{}")
		}
		return any(res).(T), nil
	case map[string]interface{}:
		res, err := simpleJson.Map()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to map[string]interface{}")
		}
		return any(res).(T), nil
	case []byte:
		res, err := simpleJson.Bytes()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to []byte")
		}
		return any(res).(T), nil
	default:
		return zero, errors.Errorf("unsupported generic type: %T. unable to resolve JSON value.", zero)
	}
}

// GetList retrieves a slice of simplejson.Json objects from the key's list (a[key]).
func GetList(simpleJson *simplejson.Json, key string) (simpleJsons []*simplejson.Json, err error) {
	if simpleJson == nil {
		return nil, errors.New("parameter simpleJson is missing")
	}
	if key == "" {
		return nil, errors.New("parameter key is missing")
	}
	elements, err := simpleJson.Get(key).Array()
	if err != nil {
		return simpleJsons, errors.WithMessage(err, "unable to get list")
	}
	return List(elements), nil
}

// Inquire queries the JSON object for the specified key, returning the parsed value, a boolean indicating success, and an error if resolution fails.
func Inquire[T any](simpleJson *simplejson.Json, key string) (T, bool, error) {
	if simpleJson == nil {
		return tern.Zero[T](), false, errors.New("parameter simpleJson is missing")
	}
	if key == "" {
		return tern.Zero[T](), false, errors.New("parameter key is missing")
	}
	value, exist := simpleJson.CheckGet(key)
	if !exist {
		return tern.Zero[T](), false, nil
	}
	res, err := Resolve[T](value)
	if err != nil {
		return tern.Zero[T](), false, errors.WithMessage(err, "unable to resolve JSON value")
	}
	return res, true, nil
}

// Attempt tries to retrieve and parse the value associated with the given key from the JSON object, returning the parsed value and a boolean indicating success.
func Attempt[T any](simpleJson *simplejson.Json, key string) (T, bool) {
	if simpleJson == nil {
		return tern.Zero[T](), false
	}
	if key == "" {
		return tern.Zero[T](), false
	}
	value, exist := simpleJson.CheckGet(key)
	if !exist {
		return tern.Zero[T](), false
	}
	res, err := Resolve[T](value)
	if err != nil {
		return tern.Zero[T](), false
	}
	return res, true
}

// Explore navigates the JSON object using a dot-separated path (e.g., "user.name"), returning the parsed value, a boolean indicating success, and an error if resolution fails.
func Explore[T any](simpleJson *simplejson.Json, path string) (T, bool, error) {
	if simpleJson == nil {
		return tern.Zero[T](), false, errors.New("parameter simpleJson is missing")
	}
	if path == "" {
		return tern.Zero[T](), false, errors.New("parameter path is missing")
	}
	var value = simpleJson
	var exist bool
	for _, key := range strings.Split(path, ".") {
		value, exist = value.CheckGet(key)
		if !exist {
			return tern.Zero[T](), false, nil
		}
	}
	res, err := Resolve[T](value)
	if err != nil {
		return tern.Zero[T](), false, errors.WithMessage(err, "unable to resolve JSON value")
	}
	return res, true, nil
}
