package controllers

import (
	"encoding/json"
	"net/http"

	"go-demos/projects/restful/auth"
	"go-demos/projects/restful/helper"
	"go-demos/projects/restful/models"

	"gopkg.in/mgo.v2/bson"
)

const (
	db         = "Movies"
	collection = "UserModel"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.UserName == "" || user.Password == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
		return
	}
	err = models.Insert(db, collection, user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError,
			helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
	}

	helper.ResponseWithJson(w, http.StatusOK,
		helper.Response{Code: http.StatusOK, Msg: "register success."})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
	}

	exist := models.IsExist(db, collection, bson.M{"username": user.UserName})
	if exist {
		token, _ := auth.GenerateToken(&user)
		helper.ResponseWithJson(w, http.StatusOK,
			helper.Response{Code: http.StatusOK, Data: models.JwtToken{Token: token}})
	} else {
		helper.ResponseWithJson(w, http.StatusNotFound,
			helper.Response{Code: http.StatusNotFound, Msg: "the user not exist"})
	}
}
