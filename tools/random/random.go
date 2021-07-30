package random

import (
	"fmt"
	"math/rand"
	"time"
)

func Random1() int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63()
}

// Random2 生成指定位数的随机数
func Random2() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func Random3(from, to int64) int64 {
	if from >= to || from == 0 || to == 0 {
		return to
	}
	rand.Seed(time.Now().Unix())
	return rand.Int63n(to-from) + from
}
