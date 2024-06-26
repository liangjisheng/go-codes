package v1transport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"v1/v1endpoint"
	"v1/v1service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// NewHTTPHandler ...
func NewHTTPHandler(endpoint v1endpoint.EndPointServer) http.Handler {
	options := []httptransport.ServerOption{
		// 程序中的全部报错都会走这里面
		httptransport.ServerErrorEncoder(errorEncode),
	}

	m := http.NewServeMux()
	m.Handle("/sum", httptransport.NewServer(
		endpoint.AddEndPoint,
		decodeHTTPADDRequest,      // 解析请求值
		encodeHTTPGenericResponse, // 返回值
		options...,
	))
	return m
}

func decodeHTTPADDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		in  v1service.Add
		err error
	)
	in.A, err = strconv.Atoi(r.FormValue("a"))
	in.B, err = strconv.Atoi(r.FormValue("b"))
	if err != nil {
		return in, err
	}
	return in, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("encodeHTTPGenericResponse", response)
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncode(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncode(_ context.Context, err error, w http.ResponseWriter) {
	fmt.Println("errorEncoder", err.Error())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

type errorWrapper struct {
	Error string `json:"errors"`
}
