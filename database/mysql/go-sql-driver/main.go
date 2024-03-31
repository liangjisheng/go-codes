package main

func main() {
	// demo()
	demo1()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
