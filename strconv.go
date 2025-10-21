package simplejsonx

import (
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"github.com/yyle88/simplejsonx/internal/utils"
)

// Strconv extracts JSON value via string bridge and converts to target type
// Uses two-stage conversion process: JSON → string → target type
// Handles int, int64, float64, string, uint64, boolean using Go's strconv package
//
// Strconv 通过字符串中介提取 JSON 值并转换成目标类型
// 使用两阶段转换过程：JSON → 字符串 → 目标类型
// 使用 Go 的 strconv 包处理 int、int64、float64、string、uint64、bool
func Strconv[T any](object *simplejson.Json) (T, error) {
	if object == nil {
		return utils.Zero[T](), errors.New("parameter object is missing")
	}
	switch zero := utils.Zero[T](); any(zero).(type) {
	case int:
		stringValue, err := object.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.Atoi(stringValue)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to int")
		}
		return any(res).(T), nil
	case int64:
		stringValue, err := object.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.ParseInt(stringValue, 10, 64)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to int64")
		}
		return any(res).(T), nil
	case float64:
		stringValue, err := object.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to float64")
		}
		return any(res).(T), nil
	case string:
		stringValue, err := object.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		return any(stringValue).(T), nil
	case uint64:
		stringValue, err := object.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.ParseUint(stringValue, 10, 64)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to uint64")
		}
		return any(res).(T), nil
	case bool:
		stringValue, err := object.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.ParseBool(stringValue)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to bool")
		}
		return any(res).(T), nil
	default:
		return zero, errors.Errorf("unsupported generic type: %T. unable to resolve JSON value.", zero)
	}
}
