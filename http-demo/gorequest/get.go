package gorequest

import (
	"github.com/parnurzeal/gorequest"
	"log"
)

type Result struct {
	Field string
}

func Get() {
	url := ""
	var result Result
	response, _, errs := gorequest.New().
		Get(url).
		EndStruct(&result)

	if errs != nil {
		log.Print("error: ", errs)
		return
	}

	response.Body.Close()
}
