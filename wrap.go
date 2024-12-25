package simplejsonx

import (
	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
)

// Load creates a simplejson.Json instance from bytes, same as simplejson.NewJson(data).
func Load(data []byte) (simpleJson *simplejson.Json, err error) {
	simpleJson, err = simplejson.NewJson(data)
	if err != nil {
		return simplejson.New(), errors.WithMessage(err, "unable to parse JSON")
	}
	return simpleJson, nil
}

// Wrap creates a simplejson.Json instance with the provided value as the root element.
func Wrap(value interface{}) (simpleJson *simplejson.Json) {
	simpleJson = simplejson.New()
	simpleJson.SetPath([]string{}, value)
	return simpleJson
}

// List converts a slice of items into a slice of simplejson.Json objects.
func List(elements []interface{}) (simpleJsons []*simplejson.Json) {
	simpleJsons = make([]*simplejson.Json, 0, len(elements))
	for _, elem := range elements {
		simpleJsons = append(simpleJsons, Wrap(elem))
	}
	return simpleJsons
}
