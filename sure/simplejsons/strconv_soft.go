package simplejsons

import (
	"github.com/bitly/go-simplejson"
	"github.com/yyle88/simplejsonx"
	"github.com/yyle88/sure"
)

func Strconv[T any](simpleJson *simplejson.Json) T {
	res0, err := simplejsonx.Strconv[T](simpleJson)
	sure.Soft(err)
	return res0
}
