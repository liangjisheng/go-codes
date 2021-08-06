package gorequest

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type ReqStruct struct {
	Name string
}

type ResStruct struct {
	Res string
}

func Post() ([]byte, error) {
	url := ""
	req := ReqStruct{
		Name: "name",
	}

	var resp ResStruct
	response, res, errs := gorequest.New().
		Post(url).SendStruct(req).
		EndStruct(&resp)
		//End()

	if errs != nil {
		return nil, fmt.Errorf("%#v", errs)
	}

	response.Body.Close()

	return []byte(res), nil
}

type ReqJsonRpc struct {
	Id      int32         `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type JsonRpcRes struct {
	Result interface{}  `json:"result"`
	Error  JSONRPCError `json:"error"`
}

type JSONRPCError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func JsonRpcPost(params []interface{}) (*JsonRpcRes, error) {
	var url string
	var js JsonRpcRes

	response, _, errs := gorequest.New().
		Post(url).
		SendStruct(ReqJsonRpc{
			Id:      1,
			JsonRpc: "2.0",
			Method:  "tx_submit",
			Params:  params,
		}).
		EndStruct(&js)

	if errs != nil {
		return nil, fmt.Errorf("%#v", errs)
	}

	response.Body.Close()

	return &js, nil
}
