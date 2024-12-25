package simplejson_soft

import (
	"github.com/bitly/go-simplejson"
	"github.com/yyle88/simplejsonx"
	"github.com/yyle88/sure"
)

func Load(data []byte) (simpleJson *simplejson.Json) {
	simpleJson, err := simplejsonx.Load(data)
	sure.Soft(err)
	return simpleJson
}

func Wrap(value interface{}) (simpleJson *simplejson.Json) {
	simpleJson = simplejsonx.Wrap(value)
	return simpleJson
}

func List(items []interface{}) (elements []*simplejson.Json) {
	elements = simplejsonx.List(items)
	return elements
}
