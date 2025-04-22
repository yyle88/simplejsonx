package simplejsonx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrconv_Int(t *testing.T) {
	simpleJson, err := Load([]byte(`{"age": "18"}`))
	require.NoError(t, err)

	res, err := Strconv[int](simpleJson.Get("age"))
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18, res)
}

func TestStrconv_Mismatch(t *testing.T) {
	simpleJson, err := Load([]byte(`{"name": "yyle88", "age": "18", "is_student": "true"}`))
	require.NoError(t, err)
	{
		res, err := Strconv[int](simpleJson.Get("age"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 18, res)
	}
	{
		res, err := Strconv[string](simpleJson.Get("age"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "18", res) // Strconv 使用 simpleJson.String() 返回字符串
	}
	{
		res, err := Strconv[string](simpleJson.Get("name"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "yyle88", res)
	}
	{
		res, err := Strconv[int](simpleJson.Get("name"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res) // int 的零值
	}
	{
		res, err := Strconv[bool](simpleJson.Get("is_student"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := Strconv[int](simpleJson.Get("is_student"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res) // int 的零值
	}
}

func TestStrconv_Int64(t *testing.T) {
	simpleJson, err := Load([]byte(`{"height": "175", "weight": "80", "temperature": "-5", "zero": "0"}`))
	require.NoError(t, err)

	{
		res, err := Strconv[int64](simpleJson.Get("height"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(175), res)
	}
	{
		res, err := Strconv[int64](simpleJson.Get("weight"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(80), res)
	}
	{
		res, err := Strconv[int64](simpleJson.Get("temperature"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(-5), res)
	}
	{
		res, err := Strconv[int64](simpleJson.Get("zero"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(0), res)
	}
}

func TestStrconv_Float64(t *testing.T) {
	simpleJson, err := Load([]byte(`{"size": "18.5", "pi": "3.14159", "large": "1e6", "small": "1e-6"}`))
	require.NoError(t, err)

	{
		res, err := Strconv[float64](simpleJson.Get("size"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 18.5, res)
	}
	{
		res, err := Strconv[float64](simpleJson.Get("pi"))
		require.NoError(t, err)
		t.Log(res)
		require.InDelta(t, 3.14159, res, 0.00001) // 浮点数精度验证
	}
	{
		res, err := Strconv[float64](simpleJson.Get("large"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 1e6, res)
	}
	{
		res, err := Strconv[float64](simpleJson.Get("small"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 1e-6, res)
	}
}

func TestStrconv_String(t *testing.T) {
	simpleJson, err := Load([]byte(`{"name": "yyle88", "age": "18", "like": "rice", "special": "hello\nworld"}`))
	require.NoError(t, err)

	{
		res, err := Strconv[string](simpleJson.Get("like"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "rice", res)
	}
	{
		res, err := Strconv[string](simpleJson.Get("special"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "hello\nworld", res)
	}
}

func TestStrconv_Uint64(t *testing.T) {
	simpleJson, err := Load([]byte(`{"endurance": "30", "persistence": "60", "zero": "0"}`))
	require.NoError(t, err)

	{
		res, err := Strconv[uint64](simpleJson.Get("endurance"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(30), res)
	}
	{
		res, err := Strconv[uint64](simpleJson.Get("persistence"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(60), res)
	}
	{
		res, err := Strconv[uint64](simpleJson.Get("zero"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(0), res)
	}
}

func TestStrconv_Bool(t *testing.T) {
	simpleJson, err := Load([]byte(`{"is_tall": "true", "is_rich": "false", "is_cool": "1", "is_smart": "0"}`))
	require.NoError(t, err)

	{
		res, err := Strconv[bool](simpleJson.Get("is_tall"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := Strconv[bool](simpleJson.Get("is_rich"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, false, res)
	}
	{
		res, err := Strconv[bool](simpleJson.Get("is_cool"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := Strconv[bool](simpleJson.Get("is_smart"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, false, res)
	}
}

func TestStrconv_None(t *testing.T) {
	res, err := Strconv[int](nil)
	require.Error(t, err)
	t.Log(err)
	require.Equal(t, 0, res) // int 的零值
}

func TestStrconv_InvalidConversion(t *testing.T) {
	simpleJson, err := Load([]byte(`{"invalid_int": "abc", "invalid_float": "def", "invalid_bool": "yes"}`))
	require.NoError(t, err)

	{
		res, err := Strconv[int](simpleJson.Get("invalid_int"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res)
	}
	{
		res, err := Strconv[float64](simpleJson.Get("invalid_float"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0.0, res)
	}
	{
		res, err := Strconv[bool](simpleJson.Get("invalid_bool"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, false, res)
	}
}

func TestStrconv_EmptyString(t *testing.T) {
	simpleJson, err := Load([]byte(`{"empty": ""}`))
	require.NoError(t, err)

	{
		res, err := Strconv[string](simpleJson.Get("empty"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, "", res)
	}
	{
		res, err := Strconv[int](simpleJson.Get("empty"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res)
	}
	{
		res, err := Strconv[float64](simpleJson.Get("empty"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0.0, res)
	}
	{
		res, err := Strconv[bool](simpleJson.Get("empty"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, false, res)
	}
}

func TestStrconv_BoundaryValues(t *testing.T) {
	simpleJson, err := Load([]byte(`{"max_int64": "9223372036854775807", "min_int64": "-9223372036854775808", "max_uint64": "18446744073709551615", "large_float": "1.7976931348623157e+308", "small_float": "-1.7976931348623157e+308"}`))
	require.NoError(t, err)

	{
		res, err := Strconv[int64](simpleJson.Get("max_int64"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(9223372036854775807), res)
	}
	{
		res, err := Strconv[int64](simpleJson.Get("min_int64"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(-9223372036854775808), res)
	}
	{
		res, err := Strconv[uint64](simpleJson.Get("max_uint64"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(18446744073709551615), res)
	}
	{
		res, err := Strconv[float64](simpleJson.Get("large_float"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, 1.7976931348623157e+308, res)
	}
	{
		res, err := Strconv[float64](simpleJson.Get("small_float"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, -1.7976931348623157e+308, res)
	}
}
