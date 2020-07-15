package main

func main() {
	// insert()
	booksDemo()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
