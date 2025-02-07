package random

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/duke-git/lancet/v2/random"
)

func RandBytes(n int) []byte {
	return random.RandBytes(n)
}

func RandNumeral(n int) string {
	return random.RandNumeral(6)
}

func RanNum() int64 {
	return rand.Int63()
}

// RanBit 生成指定位数的随机数
func RanBit() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func RanAToB(from, to int64) int64 {
	if from >= to || from == 0 || to == 0 {
		return to
	}
	return rand.Int63n(to-from) + from
}

// 生成随机字符串
var _bytes = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var r *rand.Rand

// GetRandomString 生成随机字符串
func GetRandomString(n int) string {
	result := []byte{}
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	for i := 0; i < n; i++ {
		result = append(result, _bytes[r.Intn(len(_bytes))])
	}
	return string(result)
}

// GetRangeNumString 生成随机数字字符串
func GetRangeNumString(n int) string {
	var _bytes = []byte("0123456789")
	var r *rand.Rand

	result := []byte{}
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	for i := 0; i < n; i++ {
		result = append(result, _bytes[r.Intn(len(_bytes))])
	}
	return string(result)
}

// GetRandInt 生成随机整数 digit：位数
func GetRandInt(min int, max int) int {
	if min > max {
		min = 0
		max = 0
	}
	if max == min {
		return min
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// GetCapitalRandom 生成大写随机bytes
func GetCapitalRandom(len int) []byte {
	s := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, len)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		result[i] = s[r.Intn(26)]
	}

	return result
}

// GetLowerRandom 生成小写随机bytes
func GetLowerRandom(len int) []byte {
	s := []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, len)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		result[i] = s[r.Intn(26)]
	}

	return result
}

func StringSliceItem(array []string) string {
	length := len(array)
	if length == 0 {
		return ""
	}

	index := rand.Intn(length)
	return array[index]
}

const (
	addressAlNum = "0123456789abcdefABCDEF"
)

func Address() string {
	length := len(addressAlNum)
	var res []byte
	for i := 0; i < 40; i++ {
		index := rand.Intn(length)
		res = append(res, addressAlNum[index])
	}
	return "0x" + string(res)
}
