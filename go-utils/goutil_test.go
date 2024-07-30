package utils

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func TestGoUtil(t *testing.T) {
	assert.Equal(t, true, goutil.IsEmpty(nil))
	assert.Equal(t, false, goutil.IsEmpty("abc"))

	assert.Equal(t, true, goutil.IsEqual("a", "a"))
	assert.Equal(t, true, goutil.IsEqual([]string{"a"}, []string{"a"}))
	assert.Equal(t, true, goutil.IsEqual(23, 23))

	assert.Equal(t, true, goutil.Contains("abc", "a"))
	assert.Equal(t, true, goutil.Contains([]string{"abc", "def"}, "abc"))
	assert.Equal(t, true, goutil.Contains(map[int]string{2: "abc", 4: "def"}, 4))

	// convert type
	str := goutil.String(23)     // "23"
	iVal := goutil.Int("-2")     // 2
	i64Val := goutil.Int64("-2") // -2
	u64Val := goutil.Uint("2")   // 2
	t.Log(str, iVal, i64Val, u64Val)

	arrutil.IntsHas([]int{2, 4, 5}, 2)          // True
	arrutil.Int64sHas([]int64{2, 4, 5}, 2)      // True
	arrutil.StringsHas([]string{"a", "b"}, "a") // True

	// list and val interface{}
	arrutil.Contains([]uint32{9, 2, 3}, 9) // True
}

func TestCliUtil(t *testing.T) {
	args := cliutil.ParseLine(`./app top sub --msg "has multi words"`)
	dump.P(args)

	s := cliutil.BuildLine("./myapp", []string{
		"-a", "val0",
		"-m", "this is message",
		"arg0",
	})
	fmt.Println("Build line:", s)
}
