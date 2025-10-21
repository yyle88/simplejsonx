package simplejsonx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/simplejsonx"
)

func TestLoad_ValidJSON(t *testing.T) {
	data := []byte(`{"key": "value"}`)

	object, err := simplejsonx.Load(data)
	require.NoError(t, err)

	res, err := object.Get("key").String()
	require.NoError(t, err)
	require.Equal(t, "value", res)
}

func TestLoad_InvalidJSON(t *testing.T) {
	data := []byte(`invalid message`)

	object, err := simplejsonx.Load(data)
	require.Error(t, err)
	require.NotNil(t, object)
}

func TestLoad_EmptyInput(t *testing.T) {
	data := []byte(``)

	object, err := simplejsonx.Load(data)
	require.Error(t, err)
	require.NotNil(t, object)
}

func TestWrap(t *testing.T) {
	value := map[string]interface{}{
		"key": "abc",
	}

	object := simplejsonx.Wrap(value)
	require.NotNil(t, object)

	res, err := object.Get("key").String()
	require.NoError(t, err)
	require.Equal(t, "abc", res)
}

func TestWrap_WithPrimitiveValue(t *testing.T) {
	value := 88

	object := simplejsonx.Wrap(value)
	require.NotNil(t, object)

	res, err := object.Int()
	require.NoError(t, err)
	require.Equal(t, 88, res)
}

func TestWrap_InvalidNone(t *testing.T) {
	object := simplejsonx.Wrap(nil)
	require.NotNil(t, object)

	{
		res, err := object.Int()
		require.Error(t, err)
		require.Equal(t, 0, res)
	}

	{
		res, err := object.Get("abc").String()
		require.Error(t, err)
		require.Equal(t, "", res)
	}
}

func TestList(t *testing.T) {
	data := []byte(`{
	"infos": [
		{"code":1, "name":"a"},
		{"code":2, "name":"b"},
		{"code":3, "name":"c"}
	],
	"ranks": ["x", "y", "z"]
}`)
	object, err := simplejsonx.Load(data)
	require.NoError(t, err)
	infos, err := object.Get("infos").Array()
	require.NoError(t, err)
	elements := simplejsonx.List(infos)

	var resMap = map[string]int{}
	for _, elem := range elements {
		code, err := elem.Get("code").Int()
		require.NoError(t, err)
		name, err := elem.Get("name").String()
		require.NoError(t, err)
		t.Log(code, name)
		resMap[name] = code
	}
	require.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, resMap)
}
