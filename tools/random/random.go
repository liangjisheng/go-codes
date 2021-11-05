package random

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
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
