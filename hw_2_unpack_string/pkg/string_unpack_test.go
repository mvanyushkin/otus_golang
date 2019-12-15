package unpackstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenRegularStringUnpackGot(t *testing.T) {
	unpacker := New()
	val, err := unpacker.Do("a4bc2d5e")
	assert.Nil(t, err)
	assert.Equal(t, "aaaabccddddde", val)
}

func TestWhenUnpackableStringGot(t *testing.T) {
	unpacker := New()
	val, err := unpacker.Do("abcd")
	assert.Nil(t, err)
	assert.Equal(t, "abcd", val)
}

func TestWhenWrongStringGot(t *testing.T) {
	unpacker := New()
	_, err := unpacker.Do("45")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "wrong value, that is sucks")
}

func TestWhenEscapedUnpackableStringGot(t *testing.T) {
	unpacker := New()
	v, err := unpacker.Do("qwe\\4\\5")
	assert.Nil(t, err)
	assert.Equal(t, "qwe45", v)
}

func TestWhenEscapedRegularStringGot(t *testing.T) {
	unpacker := New()
	v, err := unpacker.Do("qwe\\45")
	assert.Nil(t, err)
	assert.Equal(t, "qwe44444", v)
}

func TestWhenDoubleEscapedStringGot(t *testing.T) {
	unpacker := New()
	v, err := unpacker.Do("qwe\\\\5")
	assert.Nil(t, err)
	assert.Equal(t, "qwe\\\\\\\\\\", v)
}

func TestWhenRegularLongStringUnpackGot(t *testing.T) {
	unpacker := New()
	val, err := unpacker.Do("a10bc2d5e")
	assert.Nil(t, err)
	assert.Equal(t, "aaaaaaaaaabccddddde", val)
}

// a4bc2d5e

//"a4bc2d5e" => "aaaabccddddde"
//"abcd" => "abcd"
//"45" => "" (некорректная строка)
//"qwe\4\5" => "qwe45" (*)
//"qwe\45" => "qwe44444" (*)
//"qwe\\5" => "qwe\\\\\" (*)
