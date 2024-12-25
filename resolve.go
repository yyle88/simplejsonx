package simplejsonx

import (
	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"github.com/yyle88/tern"
)

// Extract retrieves the value associated with the given key in the JSON object.
func Extract[T any](simpleJson *simplejson.Json, key string) (T, error) {
	if simpleJson == nil {
		return tern.Zero[T](), errors.New("parameter simpleJson is missing")
	}
	return Resolve[T](simpleJson.Get(key))
}

// Inspect retrieves the value of the given key if it exists in the JSON object, returning the zero value if the key is missing.
func Inspect[T any](simpleJson *simplejson.Json, key string) (T, error) {
	if simpleJson == nil {
		return tern.Zero[T](), errors.New("parameter simpleJson is missing")
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
	elements, err := simpleJson.Get(key).Array()
	if err != nil {
		return simpleJsons, errors.WithMessage(err, "unable to get list")
	}
	return List(elements), nil
}
