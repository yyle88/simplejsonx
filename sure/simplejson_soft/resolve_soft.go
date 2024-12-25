package simplejson_soft

import (
	"github.com/bitly/go-simplejson"
	"github.com/yyle88/simplejsonx"
	"github.com/yyle88/sure"
)

func Extract[T any](simpleJson *simplejson.Json, key string) T {
	res0, err := simplejsonx.Extract[T](simpleJson, key)
	sure.Soft(err)
	return res0
}

func Inspect[T any](simpleJson *simplejson.Json, key string) T {
	res0, err := simplejsonx.Inspect[T](simpleJson, key)
	sure.Soft(err)
	return res0
}

func Resolve[T any](simpleJson *simplejson.Json) T {
	res0, err := simplejsonx.Resolve[T](simpleJson)
	sure.Soft(err)
	return res0
}

func GetList(simpleJson *simplejson.Json, key string) (simpleJsons []*simplejson.Json) {
	simpleJsons, err := simplejsonx.GetList(simpleJson, key)
	sure.Soft(err)
	return simpleJsons
}
