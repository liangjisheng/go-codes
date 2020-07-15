package controllers

import (
	"encoding/json"
	"go-demos/microservice/dao"
	"go-demos/microservice/helper"
	"go-demos/microservice/models"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

const (
	db         = "ShowTimes"
	collection = "ShowModel"
)

func CreateTimes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var times models.ShowTimes
	if err := json.NewDecoder(r.Body).Decode(&times); err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest, "invalid request")
		return
	}
	times.Id = bson.NewObjectId().Hex()
	if err := dao.Insert(db, collection, times); err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.ResponseWithJson(w, http.StatusOK, times)
}

func ShowTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	datestr := params["date"]
	var times models.ShowTimes
	if err := dao.FindOne(db, collection, bson.M{"date": datestr}, nil, &times); err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.ResponseWithJson(w, http.StatusOK, times)
}

func ShowAll(w http.ResponseWriter, r *http.Request) {
	var times []models.ShowTimes
	if err := dao.FindAll(db, collection, nil, nil, &times); err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.ResponseWithJson(w, http.StatusOK, times)
}
