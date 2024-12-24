package simplejsonx_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/simplejsonx"
)

func TestExtract_Int(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"age": 18}`))
	require.NoError(t, err)

	res, err := simplejsonx.Extract[int](simpleJson, "age")
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18, res)
}

func TestExtract_Mismatch(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"name": "yyle88"}`))
	require.NoError(t, err)

	res, err := simplejsonx.Extract[int](simpleJson, "age")
	require.Error(t, err)
	t.Log(err)
	t.Log(res)
	require.Equal(t, 0, res)
}

func TestInspect(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"age": 18}`))
	require.NoError(t, err)

	res, err := simplejsonx.Inspect[int](simpleJson, "age")
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18, res)
}

func TestInspect_Mismatch(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"name": "yyle88"}`))
	require.NoError(t, err)

	res, err := simplejsonx.Inspect[int](simpleJson, "age")
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 0, res)
}

func TestResolve_Int(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"age": 18}`))
	require.NoError(t, err)

	res, err := simplejsonx.Resolve[int](simpleJson.Get("age"))
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18, res)
}

func TestResolve_Mismatch(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": 18}`))
	require.NoError(t, err)
	{
		res, err := simplejsonx.Resolve[int](simpleJson.Get("age"))
		require.NoError(t, err)
		require.Equal(t, 18, res)
	}
	{
		res, err := simplejsonx.Resolve[string](simpleJson.Get("age"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, "", res) // zero value of string
	}
	{
		res, err := simplejsonx.Resolve[string](simpleJson.Get("name"))
		require.NoError(t, err)
		require.Equal(t, "yyle88", res)
	}
	{
		res, err := simplejsonx.Resolve[int](simpleJson.Get("name"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res) // zero value
	}
}

func TestResolve_Int64(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"height": 175, "weight": 80}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Resolve[int64](simpleJson.Get("height"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(175), res)
	}
	{
		res, err := simplejsonx.Resolve[int64](simpleJson.Get("weight"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(80), res)
	}
}

func TestResolve_Float64(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"size": 18.5}`))
	require.NoError(t, err)

	res, err := simplejsonx.Resolve[float64](simpleJson.Get("size"))
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18.5, res)
}

func TestResolve_String(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": 18, "like": "rice"}`))
	require.NoError(t, err)

	res, err := simplejsonx.Resolve[string](simpleJson.Get("like"))
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, "rice", res)
}

func TestResolve_Uint64(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"endurance": 30, "persistence": 60}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Resolve[uint64](simpleJson.Get("endurance"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(30), res)
	}
	{
		res, err := simplejsonx.Resolve[uint64](simpleJson.Get("persistence"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(60), res)
	}
}

func TestResolve_Bool(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`{"is_tall": true, "is_rich": true, "is_cool": true}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Resolve[bool](simpleJson.Get("is_tall"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := simplejsonx.Resolve[bool](simpleJson.Get("is_rich"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := simplejsonx.Resolve[bool](simpleJson.Get("is_cool"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
}

func TestResolve_StringArray(t *testing.T) {
	simpleJson, err := simplejsonx.Load([]byte(`["a", "b", "c"]`))
	require.NoError(t, err)

	res, err := simplejsonx.Resolve[[]string](simpleJson)
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, []string{"a", "b", "c"}, res)
}

func TestResolve_Array(t *testing.T) {
	{
		simpleJson, err := simplejsonx.Load([]byte(`[1, "two", 3.3]`))
		require.NoError(t, err)

		res, err := simplejsonx.Resolve[[]interface{}](simpleJson)
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, []interface{}{json.Number("1"), "two", json.Number("3.3")}, res)
	}
	{
		var value any
		decoder := json.NewDecoder(bytes.NewBufferString(`[1, "two", 3.3]`))
		decoder.UseNumber()
		require.NoError(t, decoder.Decode(&value))

		simpleJson := simplejsonx.Wrap(value)

		res, err := simplejsonx.Resolve[[]interface{}](simpleJson)
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, []interface{}{json.Number("1"), "two", json.Number("3.3")}, res)
	}
	{
		var a any
		require.NoError(t, json.Unmarshal([]byte(`[1, "two", 3.3]`), &a))

		simpleJson := simplejsonx.Wrap(a)

		res, err := simplejsonx.Resolve[[]interface{}](simpleJson)
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, []interface{}{float64(1), "two", 3.3}, res)
	}
}

func TestResolve_Map(t *testing.T) {
	simpleJson := simplejsonx.Wrap(map[string]interface{}{"key": "value"})

	res, err := simplejsonx.Resolve[map[string]interface{}](simpleJson)
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, map[string]interface{}{"key": "value"}, res)
}

func TestResolve_Bytes(t *testing.T) {
	a := map[string]interface{}{
		"value": "abc",
	}
	data, err := json.Marshal(a)
	require.NoError(t, err)
	t.Log(string(data))

	simpleJson, err := simplejsonx.Load(data)
	require.NoError(t, err)

	res, err := simplejsonx.Extract[[]byte](simpleJson, "value")
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, "abc", string(res))
}
