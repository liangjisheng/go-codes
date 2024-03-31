package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"testing"

	"v4/utils"
)

// Person ...
type Person struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

var wg sync.WaitGroup

func loginRequest() {
	defer wg.Done()
	person := Person{
		Account:  "alice",
		Password: "123456",
	}
	personByte, err := json.Marshal(person)
	if err != nil {
		utils.GetLogger().Error("json.Marshal(person) err:")
		return
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/login", bytes.NewBuffer(personByte))
	if err != nil {
		utils.GetLogger().Error("http.NewRequest err:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.GetLogger().Error("client.Do(req) err")
		return
	}
	defer resp.Body.Close()
}

func TestLoginRequest(t *testing.T) {
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go loginRequest()
	}
	wg.Wait()
}
