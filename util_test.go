package wercker

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStruct2MapReturnsFalse(t *testing.T) {
	result, ok := struct2map("keepitgreen")

	require.False(t, ok, "")
	assert.Nil(t, result, "")
}

type unannotatedStruct struct {
	Key1 string
	Key2 int
}

func TestStruct2MapUnannotatedStruct(t *testing.T) {
	input := unannotatedStruct{Key1: "Value1", Key2: 2}

	result, ok := struct2map(input)

	require.True(t, ok, "")
	assert.Equal(t, "Value1", result["Key1"], "")
	assert.Equal(t, 2, result["Key2"], "")
}

type annotatedStruct struct {
	Key1 string `map:"key1"`
	Key2 int    `map:"key2"`
}

func TestStruct2MapAnnotatedStruct(t *testing.T) {
	input := annotatedStruct{Key1: "Value1", Key2: 2}

	result, ok := struct2map(input)

	require.True(t, ok, "")
	assert.Equal(t, "Value1", result["key1"], "")
	assert.Equal(t, 2, result["key2"], "")
}

func TestStruct2MapPointerStruct(t *testing.T) {
	input := annotatedStruct{Key1: "Value1", Key2: 2}

	result, ok := struct2map(&input)

	require.True(t, ok, "")
	assert.Equal(t, "Value1", result["key1"], "")
	assert.Equal(t, 2, result["key2"], "")
}

type fancyAnnotatedStruct struct {
	// a dash should be ignored
	Key1 string `map:"-"`

	// These should contain a empty string
	Key2 string `map:"key2"`
	Key3 string `map:"key3,omitempty"`

	// These should contain a value
	Key4 string `map:"key4"`
	Key5 string `map:"key5,omitempty"`
}

func TestStruct2MapFancyAnnotatedStruct(t *testing.T) {
	input := fancyAnnotatedStruct{
		Key1: "value omitted",
		Key2: "",
		Key3: "",
		Key4: "value4",
		Key5: "value5",
	}

	result, ok := struct2map(input)
	require.True(t, ok, "")

	assert.Equal(t, 3, len(result), "incorrect number of pairs")

	// Key1 should not be present at all
	_, key1Present := result["Key1"]
	assert.False(t, key1Present, "Key1 should not be present")

	// key2 should be present with an empty string
	key2, key2Present := result["key2"]
	if assert.True(t, key2Present, "key2 should be present") {
		assert.Equal(t, "", key2, "key2 does not have the correct value")
	}

	// key3 should not be present
	_, key3Present := result["key3"]
	assert.False(t, key3Present, "key3 should not be present")

	// key4 should be present containing value4
	key4, key4Present := result["key4"]
	if assert.True(t, key4Present, "key4 should be present") {
		assert.Equal(t, "value4", key4, "key4 does not have the correct value")
	}

	// key5 should be present containing value4
	key5, key5Present := result["key5"]
	if assert.True(t, key5Present, "key5 should be present") {
		assert.Equal(t, "value5", key5, "key5 does not have the correct value")
	}

	// More keys that should not be present:

	// - should not be present
	_, dashPresent := result["-"]
	assert.False(t, dashPresent, "- should not be present")

	// key3,omitempty should not be present
	_, key3OmitEmptyPresent := result["key3,omitempty"]
	assert.False(t, key3OmitEmptyPresent, "key3,omitempty should not be present")

	// key5,omitempty should not be present
	_, key5OmitEmptyPresent := result["key5,omitempty"]
	assert.False(t, key5OmitEmptyPresent, "key5,omitempty should not be present")
}

func TestIsEmptyValue(t *testing.T) {
	tests := []interface{}{
		"",
		[]string{},
		map[string]string{},
		int(0),
		uint(0),
		int8(0),
		uint8(0),
		int16(0),
		uint16(0),
		int32(0),
		uint32(0),
		int64(0),
		uint64(0),
		uintptr(0),
		float32(0),
		float64(0),
		false,
	}

	for _, test := range tests {
		v := reflect.ValueOf(test)
		result := isEmptyValue(v)
		assert.True(t, result, fmt.Sprintf("value: %#v", v))
	}
}
