package assetdemo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// assert 断言不通过的时候, 会继续往下执行
// require 断言不通过的时候, 会报错, 不会接着往下执行

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	var a *int
	assert.Nil(t, a)

	// assert for not nil (good when you expect something)
	a = new(int)
	assert.NotNil(t, a)
	//if assert.NotNil(t, a) {
	//	// now we know that object isn't nil, we are safe to make
	//	// further assertions without causing any errors
	//	assert.Equal(t, "Something", object.Value)
	//}

}

func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}
