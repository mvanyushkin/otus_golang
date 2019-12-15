package main

func main() {

}

//func TestWhenRegularStringUnpackGot(t *testing.T) {
//	val, err := DoUnpackString("a4bc2d5e")
//	assert.Nil(t, err)
//	assert.Equal(t, "aaaabccddddde", val)
//}
//
//func TestWhenUnpackableStringGot(t *testing.T) {
//	val, err := DoUnpackString("abcd")
//	assert.Nil(t, err)
//	assert.Equal(t, "abcd", val)
//}
//
//func TestWhenWrongStringGot(t *testing.T) {
//	_, err := DoUnpackString("45")
//	assert.NotNil(t, err)
//	assert.Equal(t, err.Error(), "wrong value, that is sucks")
//}
//
//func TestWhenEscapedUnpackableStringGot(t *testing.T) {
//	v, err := DoUnpackString("qwe\\4\\5")
//	assert.Nil(t, err)
//	assert.Equal(t, "qwe45", v)
//}
//
//func TestWhenEscapedRegularStringGot(t *testing.T) {
//	v, err := DoUnpackString("qwe\\45")
//	assert.Nil(t, err)
//	assert.Equal(t, "qwe44444", v)
//}
//
//func TestWhenDoubleEscapedStringGot(t *testing.T) {
//	v, err := DoUnpackString("qwe\\\\5")
//	assert.Nil(t, err)
//	assert.Equal(t, "qwe\\\\\\\\\\", v)
//}
//
//func TestWhenRegularLongStringUnpackGot(t *testing.T) {
//	val, err := DoUnpackString("a10bc2d5e")
//	assert.Nil(t, err)
//	assert.Equal(t, "aaaaaaaaaabccddddde", val)
//}
