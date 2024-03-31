package main

type ReqHello1 struct {
	Name string `json:"name"`
}

type ResHello1 struct {
	Name string `json:"name"`
}

type ResponseError struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}
