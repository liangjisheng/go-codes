package demo1

// StringSliceEqual ...
func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// []string{}和[]string(nil) 这时两个字符串切片的长度都是0 但肯定不相等
	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// ShouldSummerBeComming 实现一个简单的assertion函数
func ShouldSummerBeComming(actual interface{}, expected ...interface{}) string {
	if actual == "summer" && expected[0] == "comming" {
		return ""
	}
	return "summer is not comming!"
}
