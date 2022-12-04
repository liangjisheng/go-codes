package main

import (
	"fmt"
	"log"
	"mvdan.cc/xurls/v2"
)

func main() {
	rxRelaxed := xurls.Relaxed()
	fmt.Println(rxRelaxed.FindString("Do gophers live in golang.org and foo.com?"))
	fmt.Println(rxRelaxed.FindAllString("Do gophers live in golang.org and foo.com?", -1))
	fmt.Println(rxRelaxed.FindString("This string does not have a URL"))
	//Output:
	//golang.org
	//[golang.org foo.com]
	//

	rxStrict := xurls.Strict()
	fmt.Println(rxStrict.FindAllString("must have scheme: http://foo.com/.", -1)) // []string{"http://foo.com/"}
	fmt.Println(rxStrict.FindAllString("no scheme, no match: foo.com", -1))       // []string{}
	//Output:
	//[http://foo.com/]
	//[]

	input := "The webiste is https://golangbyexample.com:8000/tutorials/intro amd mail to mailto:contactus@golangbyexample.com"
	fmt.Println(rxStrict.FindAllString(input, -1))
	//Output:
	//[https://golangbyexample.com:8000/tutorials/intro mailto:contactus@golangbyexample.com]

	//如果我们想把输出限制在一个特定的方案上，也可以这样做
	xurlsStrict, err := xurls.StrictMatchingScheme("https")
	if err != nil {
		log.Fatalf("Some error occured. Error: %s", err)
	}
	fmt.Println(xurlsStrict.FindAllString(input, -1))
	//Output:
	//[https://golangbyexample.com:8000/tutorials/intro]
}
