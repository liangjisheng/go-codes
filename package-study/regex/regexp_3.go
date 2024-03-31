package main

import (
	"fmt"
	"regexp"
)

// IsTelNumber ...
func IsTelNumber(telNum string) (bool, error) {
	m, err := regexp.MatchString("^[0-9]{11}$", telNum)
	if m {
		return true, err
	}
	return false, err
}

func regexp3() {
	fmt.Println("test go regexp(telNum)")
	retVal, _ := IsTelNumber("15202992212")
	if retVal {
		fmt.Println("this is a telphone address")
	} else {
		fmt.Println("this is noa a telphone address")
	}
}
