package unpackstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenRegularStringUnpackGot(t *testing.T) {
	val, err := DoUnpackString("a4bc2d5e")
	assert.Nil(t, err)
	assert.Equal(t, "aaaabccddddde", val)
}

func TestWhenUnpackableStringGot(t *testing.T) {
	val, err := DoUnpackString("abcd")
	assert.Nil(t, err)
	assert.Equal(t, "abcd", val)
}

func TestWhenWrongStringGot(t *testing.T) {
	_, err := DoUnpackString("45")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "wrong value, that is sucks")
}

func TestWhenEscapedUnpackableStringGot(t *testing.T) {
	v, err := DoUnpackString("qwe\\4\\5")
	assert.Nil(t, err)
	assert.Equal(t, "qwe45", v)
}

func TestWhenEscapedRegularStringGot(t *testing.T) {
	v, err := DoUnpackString("qwe\\45")
	assert.Nil(t, err)
	assert.Equal(t, "qwe44444", v)
}

func TestWhenDoubleEscapedStringGot(t *testing.T) {
	v, err := DoUnpackString("qwe\\\\5")
	assert.NotNil(t, err)
	assert.Equal(t, "qwe\\\\", v)
}

func TestWhenRegularLongStringUnpackGot(t *testing.T) {
	val, err := DoUnpackString("a10bc2d5e")
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