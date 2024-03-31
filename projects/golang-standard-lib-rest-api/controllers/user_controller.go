package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-demos/projects/golang-standard-lib-rest-api/repositories"
	"go-demos/projects/golang-standard-lib-rest-api/requests"
	"go-demos/projects/golang-standard-lib-rest-api/utils/caching"
	"go-demos/projects/golang-standard-lib-rest-api/utils/crypto"
	"log"
	"net/http"
	"strconv"
	"time"
)

// UserController ...
type UserController struct {
	DB    *sql.DB
	Cache caching.Cache
}

// NewUserController 返回带有数据库和缓存对象的结构体
func NewUserController(db *sql.DB, c caching.Cache) *UserController {
	return &UserController{
		DB:    db,
		Cache: c,
	}
}

// Register ...
// curl -X POST http://127.0.0.1:8000/register -d '{ "email":"example@qq.com", "name":"alice", "password":"alice" }'
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var rr requests.RegisterRequest
	err := decoder.Decode(&rr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := repositories.CreateUser(uc.DB, rr.Email, rr.Name, rr.Password)
	if err != nil {
		log.Printf("%+v\n", rr)
		log.Fatalf("Add user to database error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	token, err := crypto.GenerateToken()
	if err != nil {
		log.Fatalf("Generate token Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	oneMonth := time.Duration(60*60*24*30) * time.Second
	err = uc.Cache.Set(fmt.Sprintf("token_%s", token), strconv.Itoa(id), oneMonth)
	if err != nil {
		log.Fatalf("Add token to redis Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	p := map[string]string{
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// Login ...
// curl -X POST http://127.0.0.1:8000/login -d '{ "email":"example@qq.com", "password":"alice" }'
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var lr requests.LoginRequest
	err := decoder.Decode(&lr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := repositories.GetPrivateUserDetailsByEmail(uc.DB, lr.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusBadRequest)
			return
		}
		log.Fatalf("Create User Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	password := crypto.HashPassword(lr.Password, user.Salt)
	if user.Password != password {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	token, err := crypto.GenerateToken()
	if err != nil {
		log.Fatalf("Create User Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	oneMonth := time.Duration(60*60*24*30) * time.Second
	err = uc.Cache.Set(fmt.Sprintf("token_%s", token), strconv.Itoa(user.ID), oneMonth)
	if err != nil {
		log.Fatalf("Create User Error: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	p := map[string]string{
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
