package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-demos/projects/golang-standard-lib-rest-api/repositories"
	"go-demos/projects/golang-standard-lib-rest-api/requests"
	"go-demos/projects/golang-standard-lib-rest-api/utils/caching"
	"log"
	"net/http"
	"path"
	"strconv"
)

// JobController ...
type JobController struct {
	DB    *sql.DB
	Cache caching.Cache
}

// NewJobController ...
func NewJobController(db *sql.DB, c caching.Cache) *JobController {
	return &JobController{
		DB:    db,
		Cache: c,
	}
}

// Create create job
// curl -X POST http://127.0.0.1:8000/job -d '{ "title":"job1", "description":"job1 description" }' -H "token":"Vv7aImdk0VdqpzYF_I5KvSRJUYukPaR-10-19afm8EjxVN0L2jopznXzWcHck5CT39CtpWkkhb-59gWNZOgbew=="
// curl -X POST http://127.0.0.1:8000/job -d '{ "title":"job2", "description":"job2 description" }' -H "token":"Vv7aImdk0VdqpzYF_I5KvSRJUYukPaR-10-19afm8EjxVN0L2jopznXzWcHck5CT39CtpWkkhb-59gWNZOgbew=="
func (jc *JobController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	token := r.Header.Get("token")
	userIDStr, err := jc.Cache.Get(fmt.Sprintf("token_%s", token))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatalf("Convert user id to int: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var cjr requests.CreateJobRequest
	err = decoder.Decode(&cjr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	_, err = repositories.CreateJob(jc.DB, cjr.Title, cjr.Description, userID)
	if err != nil {
		log.Fatalf("Creating a job: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Job ...
// curl -X GET http://127.0.0.1:8000/job/1
// curl -X PUT http://127.0.0.1:8000/job/1 -d '{ "title":"job1-update", "description":"description update" }' -H "token":"Vv7aImdk0VdqpzYF_I5KvSRJUYukPaR-10-19afm8EjxVN0L2jopznXzWcHck5CT39CtpWkkhb-59gWNZOgbew=="
// curl -X DELETE http://127.0.0.1:8000/job/2 -d '{ "title":"job1-update", "description":"description update" }' -H "token":"Vv7aImdk0VdqpzYF_I5KvSRJUYukPaR-10-19afm8EjxVN0L2jopznXzWcHck5CT39CtpWkkhb-59gWNZOgbew=="
func (jc *JobController) Job(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	JobID, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	job, err := repositories.GetJobByID(jc.DB, JobID)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(job)
		return
	}

	token := r.Header.Get("token")
	userIDStr, err := jc.Cache.Get(fmt.Sprintf("token_%s", token))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatalf("Convert user id to int: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if userID != job.UserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
		var ujr requests.UpdateJobRequest
		err = decoder.Decode(&ujr)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = repositories.UpdateJob(jc.DB, job.ID, ujr.Title, ujr.Description)
		if err != nil {
			log.Fatalf("Updating a job: %s", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "DELETE" {
		err = repositories.DeleteJob(jc.DB, job.ID)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
}

// Feed ...
// curl -s -X GET 'http://127.0.0.1:8000/job/feed?page=1&results_per_page=5'
// curl -X GET http://127.0.0.1:8000/job/feed
func (jc *JobController) Feed(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	page := 1
	var err error
	pageStr, ok := r.URL.Query()["page"]
	if ok {
		page, err = strconv.Atoi(pageStr[0])
		if err != nil {
			page = 1
		}
	}

	resultsPerPage := 10
	resultsPerPageStr, ok := r.URL.Query()["results_per_page"]
	if ok {
		resultsPerPage, err = strconv.Atoi(resultsPerPageStr[0])
		if err != nil {
			resultsPerPage = 1
		}
	}

	jobs, err := repositories.GetJobs(jc.DB, page, resultsPerPage)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(jobs)
}
