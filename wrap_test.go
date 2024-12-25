package simplejsonx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/simplejsonx"
)

func TestLoad_ValidJSON(t *testing.T) {
	data := []byte(`{"key": "value"}`)

	simpleJson, err := simplejsonx.Load(data)
	require.NoError(t, err)

	res, err := simpleJson.Get("key").String()
	require.NoError(t, err)
	require.Equal(t, "value", res)
}

func TestLoad_InvalidJSON(t *testing.T) {
	data := []byte(`invalid message`)

	simpleJson, err := simplejsonx.Load(data)
	require.Error(t, err)
	require.NotNil(t, simpleJson)
}

func TestLoad_EmptyInput(t *testing.T) {
	data := []byte(``)

	simpleJson, err := simplejsonx.Load(data)
	require.Error(t, err)
	require.NotNil(t, simpleJson)
}

func TestWrap(t *testing.T) {
	value := map[string]interface{}{
		"key": "abc",
	}

	simpleJson := simplejsonx.Wrap(value)
	require.NotNil(t, simpleJson)

	res, err := simpleJson.Get("key").String()
	require.NoError(t, err)
	require.Equal(t, "abc", res)
}

func TestWrap_WithPrimitiveValue(t *testing.T) {
	value := 88

	simpleJson := simplejsonx.Wrap(value)
	require.NotNil(t, simpleJson)

	res, err := simpleJson.Int()
	require.NoError(t, err)
	require.Equal(t, 88, res)
}

func TestWrap_InvalidNone(t *testing.T) {
	simpleJson := simplejsonx.Wrap(nil)
	require.NotNil(t, simpleJson)

	{
		res, err := simpleJson.Int()
		require.Error(t, err)
		require.Equal(t, 0, res)
	}

	{
		res, err := simpleJson.Get("abc").String()
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
	simpleJson, err := simplejsonx.Load(data)
	require.NoError(t, err)
	infos, err := simpleJson.Get("infos").Array()
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
