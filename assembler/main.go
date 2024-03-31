package main

func add(a, b int) int {
	sum := 0 // 不设置该局部变量sum，add栈空间大小会是0
	sum = a + b
	return sum

}

// go tool compile -N -l -S main.go

func main() {
	println(add(1, 2))
}
