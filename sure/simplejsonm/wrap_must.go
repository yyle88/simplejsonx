package simplejsonm

import (
	"github.com/bitly/go-simplejson"
	"github.com/yyle88/simplejsonx"
	"github.com/yyle88/sure"
)

func Load(data []byte) (object *simplejson.Json) {
	object, err := simplejsonx.Load(data)
	sure.Must(err)
	return object
}

func Wrap(value interface{}) (object *simplejson.Json) {
	object = simplejsonx.Wrap(value)
	return object
}

func List(elements []interface{}) (objects []*simplejson.Json) {
	objects = simplejsonx.List(elements)
	return objects
}
