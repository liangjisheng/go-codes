package main

import (
	"fmt"
	"regexp"
)

func main() {
	// re1()
	// re2()
	// re3()

	// regexp1()
	// regexp2()
	// regexp3()
	// regexp4()
	// regexp5()
	// regexp6()
	// regexp7()
	// regexp8()

	regtest()
}

func regtest() {
	regex := "^[A-Z]{3}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{5}$"
	target := "APL-7K2R-EFGG-FKAV-E3CMH"
	match, _ := regexp.MatchString(regex, target)
	fmt.Println(match)
}
