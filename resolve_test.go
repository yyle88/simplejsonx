package simplejsonx_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/simplejsonx"
)

func TestExtract_Int(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"age": 18}`))
	require.NoError(t, err)

	res, err := simplejsonx.Extract[int](object, "age")
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18, res)
}

func TestExtract_Mismatch(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"name": "yyle88"}`))
	require.NoError(t, err)

	res, err := simplejsonx.Extract[int](object, "age")
	require.Error(t, err)
	t.Log(err)
	t.Log(res)
	require.Equal(t, 0, res)
}

func TestInspect(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"age": 18}`))
	require.NoError(t, err)

	res, err := simplejsonx.Inspect[int](object, "age")
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18, res)
}

func TestInspect_Mismatch(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"name": "yyle88"}`))
	require.NoError(t, err)

	res, err := simplejsonx.Inspect[int](object, "age")
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 0, res)
}

func TestResolve_Int(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"age": 18}`))
	require.NoError(t, err)

	res, err := simplejsonx.Resolve[int](object.Get("age"))
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18, res)
}

func TestResolve_Mismatch(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": 18}`))
	require.NoError(t, err)
	{
		res, err := simplejsonx.Resolve[int](object.Get("age"))
		require.NoError(t, err)
		require.Equal(t, 18, res)
	}
	{
		res, err := simplejsonx.Resolve[string](object.Get("age"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, "", res) // zero value of string
	}
	{
		res, err := simplejsonx.Resolve[string](object.Get("name"))
		require.NoError(t, err)
		require.Equal(t, "yyle88", res)
	}
	{
		res, err := simplejsonx.Resolve[int](object.Get("name"))
		require.Error(t, err)
		t.Log(err)
		require.Equal(t, 0, res) // zero value
	}
}

func TestResolve_Int64(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"height": 175, "weight": 80}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Resolve[int64](object.Get("height"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(175), res)
	}
	{
		res, err := simplejsonx.Resolve[int64](object.Get("weight"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, int64(80), res)
	}
}

func TestResolve_Float64(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"size": 18.5}`))
	require.NoError(t, err)

	res, err := simplejsonx.Resolve[float64](object.Get("size"))
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, 18.5, res)
}

func TestResolve_String(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"name": "yyle88", "age": 18, "like": "rice"}`))
	require.NoError(t, err)

	res, err := simplejsonx.Resolve[string](object.Get("like"))
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, "rice", res)
}

func TestResolve_Uint64(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"endurance": 30, "persistence": 60}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Resolve[uint64](object.Get("endurance"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(30), res)
	}
	{
		res, err := simplejsonx.Resolve[uint64](object.Get("persistence"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, uint64(60), res)
	}
}

func TestResolve_Bool(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`{"is_tall": true, "is_rich": true, "is_cool": true}`))
	require.NoError(t, err)

	{
		res, err := simplejsonx.Resolve[bool](object.Get("is_tall"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := simplejsonx.Resolve[bool](object.Get("is_rich"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
	{
		res, err := simplejsonx.Resolve[bool](object.Get("is_cool"))
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, true, res)
	}
}

func TestResolve_StringArray(t *testing.T) {
	object, err := simplejsonx.Load([]byte(`["a", "b", "c"]`))
	require.NoError(t, err)

	res, err := simplejsonx.Resolve[[]string](object)
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, []string{"a", "b", "c"}, res)
}

func TestResolve_Array(t *testing.T) {
	{
		object, err := simplejsonx.Load([]byte(`[1, "two", 3.3]`))
		require.NoError(t, err)

		res, err := simplejsonx.Resolve[[]interface{}](object)
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, []interface{}{json.Number("1"), "two", json.Number("3.3")}, res)
	}
	{
		var value any
		decoder := json.NewDecoder(bytes.NewBufferString(`[1, "two", 3.3]`))
		decoder.UseNumber()
		require.NoError(t, decoder.Decode(&value))

		object := simplejsonx.Wrap(value)

		res, err := simplejsonx.Resolve[[]interface{}](object)
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, []interface{}{json.Number("1"), "two", json.Number("3.3")}, res)
	}
	{
		var a any
		require.NoError(t, json.Unmarshal([]byte(`[1, "two", 3.3]`), &a))

		object := simplejsonx.Wrap(a)

		res, err := simplejsonx.Resolve[[]interface{}](object)
		require.NoError(t, err)
		t.Log(res)
		require.Equal(t, []interface{}{float64(1), "two", 3.3}, res)
	}
}

func TestResolve_Map(t *testing.T) {
	object := simplejsonx.Wrap(map[string]interface{}{"key": "value"})

	res, err := simplejsonx.Resolve[map[string]interface{}](object)
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

	object, err := simplejsonx.Load(data)
	require.NoError(t, err)

	res, err := simplejsonx.Extract[[]byte](object, "value")
	require.NoError(t, err)
	t.Log(res)
	require.Equal(t, "abc", string(res))
}

func TestGetList(t *testing.T) {
	jsonData := `{
		"key": [
			{"name": "item1"},
			{"name": "item2"},
			{"name": "item3"}
		]
	}`

	object, err := simplejsonx.Load([]byte(jsonData))
	require.NoError(t, err)

	{
		objects, err := simplejsonx.GetList(object, "key")
		require.NoError(t, err)
		require.Len(t, objects, 3)

		var names = make([]string, 0, len(objects))
		for _, item := range objects {
			name, err := simplejsonx.Extract[string](item, "name")
			require.NoError(t, err)
			names = append(names, name)
		}
		require.Equal(t, []string{"item1", "item2", "item3"}, names)
	}

	{
		objects, err := simplejsonx.GetList(object, "invalidKey")
		require.Error(t, err)
		require.Len(t, objects, 0)
	}
}

func TestInquire(t *testing.T) {
	// Test success scenario
	// 测试成功场景
	object, err := simplejsonx.Load([]byte(`{"age": 18}`))
	require.NoError(t, err)

	{
		res, exists, err := simplejsonx.Inquire[int](object, "age")
		require.NoError(t, err)
		require.True(t, exists, "key should exist")
		t.Log(res)
		require.Equal(t, 18, res)
	}

	// Test missing key scenario
	// 测试键不存在场景
	{
		res, exists, err := simplejsonx.Inquire[int](object, "name")
		require.NoError(t, err)
		require.False(t, exists, "key should not exist")
		t.Log(res)
		require.Equal(t, 0, res) // zero value / 零值
	}

	// Test resolution failure scenario
	// 测试解析失败场景
	{
		res, exists, err := simplejsonx.Inquire[string](object, "age")
		require.Error(t, err, "should fail to resolve int to string")
		require.False(t, exists, "should return false on resolve failure")
		t.Log(res)
		require.Equal(t, "", res) // zero value / 零值
	}
}

func TestAttempt(t *testing.T) {
	// Test success scenario
	// 测试成功场景
	object, err := simplejsonx.Load([]byte(`{"age": 18}`))
	require.NoError(t, err)

	{
		res, ok := simplejsonx.Attempt[int](object, "age")
		require.True(t, ok, "should succeed")
		t.Log(res)
		require.Equal(t, 18, res)
	}

	// Test missing key scenario
	// 测试键不存在场景
	{
		res, ok := simplejsonx.Attempt[int](object, "name")
		require.False(t, ok, "should fail due to missing key")
		t.Log(res)
		require.Equal(t, 0, res) // zero value / 零值
	}

	// Test resolution failure scenario
	// 测试解析失败场景
	{
		res, ok := simplejsonx.Attempt[string](object, "age")
		require.False(t, ok, "should fail to resolve int to string")
		t.Log(res)
		require.Equal(t, "", res) // zero value / 零值
	}
}

func TestExplore(t *testing.T) {
	// Test success scenario with nested path
	// 测试成功场景（嵌套路径）
	object, err := simplejsonx.Load([]byte(`{"user": {"name": "Alice"}}`))
	require.NoError(t, err)

	{
		res, exists, err := simplejsonx.Explore[string](object, "user.name")
		require.NoError(t, err)
		require.True(t, exists, "path should exist")
		t.Log(res)
		require.Equal(t, "Alice", res)
	}

	// Test non-existent path scenario
	// 测试路径不存在场景
	{
		res, exists, err := simplejsonx.Explore[string](object, "user.address")
		require.NoError(t, err)
		require.False(t, exists, "path should not exist")
		t.Log(res)
		require.Equal(t, "", res) // zero value / 零值
	}

	// Test resolution failure scenario
	// 测试解析失败场景
	{
		res, exists, err := simplejsonx.Explore[int](object, "user.name")
		require.Error(t, err, "should fail to resolve string to int")
		require.False(t, exists, "should return false on resolve failure")
		t.Log(res)
		require.Equal(t, 0, res) // zero value / 零值
	}

	// Test empty path scenario
	// 测试空路径场景
	{
		res, exists, err := simplejsonx.Explore[int](object, "")
		require.Error(t, err, "should fail due to empty path")
		require.False(t, exists, "should return false for empty path")
		t.Log(res)
		require.Equal(t, 0, res) // zero value / 零值
	}
}

func TestExploreGetJson(t *testing.T) {
	// Test success scenario with nested path
	// 测试成功场景（嵌套路径）
	object, err := simplejsonx.Load([]byte(`{"user": {"name": "Alice", "events":[{"eventName":"eat"}, {"eventName":"sleep"}]}}`))
	require.NoError(t, err)
	res, ok, err := simplejsonx.Explore[*simplejson.Json](object, "user")
	require.NoError(t, err)
	require.True(t, ok)
	t.Log(res)
	elements, err := simplejsonx.GetList(res, "events")
	require.NoError(t, err)
	t.Log(len(elements))
	require.Len(t, elements, 2)
	for _, item := range elements {
		t.Log(item)
	}
}

func TestExploreGetList(t *testing.T) {
	// Test success scenario with nested path
	// 测试成功场景（嵌套路径）
	object, err := simplejsonx.Load([]byte(`{"user": {"name": "Alice", "events":[{"eventName":"eat"}, {"eventName":"sleep"}]}}`))
	require.NoError(t, err)
	elements, ok, err := simplejsonx.Explore[[]*simplejson.Json](object, "user.events")
	require.NoError(t, err)
	require.True(t, ok)
	t.Log(len(elements))
	require.Len(t, elements, 2)
	for _, item := range elements {
		t.Log(item)
	}
}
