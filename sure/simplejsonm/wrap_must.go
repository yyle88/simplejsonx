package simplejsonm

import (
	"github.com/bitly/go-simplejson"
	"github.com/yyle88/simplejsonx"
	"github.com/yyle88/sure"
)

func Load(data []byte) (simpleJson *simplejson.Json) {
	simpleJson, err := simplejsonx.Load(data)
	sure.Must(err)
	return simpleJson
}

func Wrap(value interface{}) (simpleJson *simplejson.Json) {
	simpleJson = simplejsonx.Wrap(value)
	return simpleJson
}

func List(elements []interface{}) (simpleJsons []*simplejson.Json) {
	simpleJsons = simplejsonx.List(elements)
	return simpleJsons
}
