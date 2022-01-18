// Copyright 2017 Emir Ribic. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Golang SwaggerUI example
//
// This documentation describes example APIs found under https://github.com/liangjisheng
//
//     Schemes: http
//     BasePath: /v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: liangjisheng <1294851990@qq.com> https://liangjisheng.github.io
//     Host:
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"log"
	"net/http"
	"swagger/go-swagger/cmd/api"

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"

	_ "swagger/go-swagger/cmd/statik"
	_ "swagger/go-swagger/cmd/swagger"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/swaggerui/", staticServer)
	router.PathPrefix("/swaggerui/").Handler(sh)
	registerV1Routes(router)
	log.Println("server listen on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func registerV1Routes(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()
	api.RegisterRepoRoutes(v1, "/repo")
	api.RegisterUserRoutes(v1, "/user")
}
