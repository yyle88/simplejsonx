package simplejson_must

import (
	"github.com/bitly/go-simplejson"
	"github.com/yyle88/simplejsonx"
	"github.com/yyle88/sure"
)

func Extract[T any](simpleJson *simplejson.Json, key string) T {
	res0, err := simplejsonx.Extract[T](simpleJson, key)
	sure.Must(err)
	return res0
}

func Resolve[T any](simpleJson *simplejson.Json) T {
	res0, err := simplejsonx.Resolve[T](simpleJson)
	sure.Must(err)
	return res0
}
