package simplejsonx

import (
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"github.com/yyle88/simplejsonx/internal/utils"
)

// Strconv extracts and converts the value from the provided JSON object into the specified type.
func Strconv[T any](simpleJson *simplejson.Json) (T, error) {
	if simpleJson == nil {
		return utils.Zero[T](), errors.New("parameter simpleJson is missing")
	}
	switch zero := utils.Zero[T](); any(zero).(type) {
	case int:
		stv, err := simpleJson.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.Atoi(stv)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to int")
		}
		return any(res).(T), nil
	case int64:
		stv, err := simpleJson.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.ParseInt(stv, 10, 64)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to int64")
		}
		return any(res).(T), nil
	case float64:
		stv, err := simpleJson.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.ParseFloat(stv, 64)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to float64")
		}
		return any(res).(T), nil
	case string:
		stv, err := simpleJson.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		return any(stv).(T), nil
	case uint64:
		stv, err := simpleJson.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.ParseUint(stv, 10, 64)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to uint64")
		}
		return any(res).(T), nil
	case bool:
		stv, err := simpleJson.String()
		if err != nil {
			return zero, errors.WithMessage(err, "unable to resolve JSON value to string")
		}
		res, err := strconv.ParseBool(stv)
		if err != nil {
			return zero, errors.WithMessage(err, "unable to convert JSON value to bool")
		}
		return any(res).(T), nil
	default:
		return zero, errors.Errorf("unsupported generic type: %T. unable to resolve JSON value.", zero)
	}
}
