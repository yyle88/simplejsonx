package simplejsonx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/simplejsonx"
)

func TestStrconv_Int(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"age": "18"}`))
	require.NoError(t, err)

	res, err := simplejsonx.Strconv[int](object.Get("age"))
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18, res)
}

func TestStrconv_Mismatch(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": "18", "is_student": "true"}`))
	require.NoError(t, err)
	{
		res, err := simplejsonx.Strconv[int](object.Get("age"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 18, res)
	}
	{
		res, err := simplejsonx.Strconv[string](object.Get("age"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "18", res) // Strconv 使用 object.String() 返回字符串
	}
	{
		res, err := simplejsonx.Strconv[string](object.Get("name"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "yyle88", res)
	}
	{
		res, err := simplejsonx.Strconv[int](object.Get("name"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res) // int 的零值
	}
	{
		res, err := simplejsonx.Strconv[bool](object.Get("is_student"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := simplejsonx.Strconv[int](object.Get("is_student"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res) // int 的零值
	}
}

func TestStrconv_Int64(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"height": "175", "weight": "80", "temperature": "-5", "zero": "0"}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Strconv[int64](object.Get("height"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(175), res)
	}
	{
		res, err := simplejsonx.Strconv[int64](object.Get("weight"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(80), res)
	}
	{
		res, err := simplejsonx.Strconv[int64](object.Get("temperature"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(-5), res)
	}
	{
		res, err := simplejsonx.Strconv[int64](object.Get("zero"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(0), res)
	}
}

func TestStrconv_Float64(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"size": "18.5", "pi": "3.14159", "large": "1e6", "small": "1e-6"}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Strconv[float64](object.Get("size"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 18.5, res)
	}
	{
		res, err := simplejsonx.Strconv[float64](object.Get("pi"))
		require.NoError(t, err)
		t.Log(res)
		require.InDelta(t, 3.14159, res, 0.00001) // 浮点数精度验证
	}
	{
		res, err := simplejsonx.Strconv[float64](object.Get("large"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 1e6, res)
	}
	{
		res, err := simplejsonx.Strconv[float64](object.Get("small"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 1e-6, res)
	}
}

func TestStrconv_String(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": "18", "like": "rice", "special": "hello\nworld"}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Strconv[string](object.Get("like"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "rice", res)
	}
	{
		res, err := simplejsonx.Strconv[string](object.Get("special"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "hello\nworld", res)
	}
}

func TestStrconv_Uint64(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"endurance": "30", "persistence": "60", "zero": "0"}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Strconv[uint64](object.Get("endurance"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(30), res)
	}
	{
		res, err := simplejsonx.Strconv[uint64](object.Get("persistence"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(60), res)
	}
	{
		res, err := simplejsonx.Strconv[uint64](object.Get("zero"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(0), res)
	}
}

func TestStrconv_Bool(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"is_tall": "true", "is_rich": "false", "is_cool": "1", "is_smart": "0"}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Strconv[bool](object.Get("is_tall"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := simplejsonx.Strconv[bool](object.Get("is_rich"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, false, res)
	}
	{
		res, err := simplejsonx.Strconv[bool](object.Get("is_cool"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := simplejsonx.Strconv[bool](object.Get("is_smart"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, false, res)
	}
}

func TestStrconv_None(t *testing.T) {
	res, err := simplejsonx.Strconv[int](nil)
	require.Error(t, err)
	t.Log(err)
	require.Equal(t, 0, res) // int 的零值
}

func TestStrconv_InvalidConversion(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"invalid_int": "abc", "invalid_float": "def", "invalid_bool": "yes"}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Strconv[int](object.Get("invalid_int"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res)
	}
	{
		res, err := simplejsonx.Strconv[float64](object.Get("invalid_float"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0.0, res)
	}
	{
		res, err := simplejsonx.Strconv[bool](object.Get("invalid_bool"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, false, res)
	}
}

func TestStrconv_EmptyString(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"empty": ""}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Strconv[string](object.Get("empty"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "", res)
	}
	{
		res, err := simplejsonx.Strconv[int](object.Get("empty"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res)
	}
	{
		res, err := simplejsonx.Strconv[float64](object.Get("empty"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0.0, res)
	}
	{
		res, err := simplejsonx.Strconv[bool](object.Get("empty"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, false, res)
	}
}

func TestStrconv_BoundaryValues(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"max_int64": "9223372036854775807", "min_int64": "-9223372036854775808", "max_uint64": "18446744073709551615", "large_float": "1.7976931348623157e+308", "small_float": "-1.7976931348623157e+308"}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Strconv[int64](object.Get("max_int64"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(9223372036854775807), res)
	}
	{
		res, err := simplejsonx.Strconv[int64](object.Get("min_int64"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(-9223372036854775808), res)
	}
	{
		res, err := simplejsonx.Strconv[uint64](object.Get("max_uint64"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(18446744073709551615), res)
	}
	{
		res, err := simplejsonx.Strconv[float64](object.Get("large_float"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 1.7976931348623157e+308, res)
	}
	{
		res, err := simplejsonx.Strconv[float64](object.Get("small_float"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, -1.7976931348623157e+308, res)
	}
}
