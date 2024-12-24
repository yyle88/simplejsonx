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
