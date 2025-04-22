package simplejsono

import (
	"github.com/bitly/go-simplejson"
	"github.com/yyle88/simplejsonx"
	"github.com/yyle88/sure"
)

func Strconv[T any](simpleJson *simplejson.Json) T {
	res0, err := simplejsonx.Strconv[T](simpleJson)
	sure.Omit(err)
	return res0
}
