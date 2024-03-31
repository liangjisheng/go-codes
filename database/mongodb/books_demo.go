package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	db         = "test"
	collection = "books"
)

func errorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message:%q}", message)
}

func responseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	w.WriteHeader(code)
	w.Write(json)
}

type book struct {
	ISBN    string   `json:"isbn"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Price   string   `json:"price"`
}

func booksDemo() {
	session, err := mgo.Dial("localhost:27017")
	checkErr(err)
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/books"), allBooks(session))
	mux.HandleFunc(pat.Post("/books"), addBook(session))
	mux.HandleFunc(pat.Get("/books/:isbn"), bookByISBN(session))
	mux.HandleFunc(pat.Put("/books/:isbn"), updateBook(session))
	mux.HandleFunc(pat.Delete("/books/:isbn"), deleteBook(session))

	http.ListenAndServe(":8080", mux)
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(db).C(collection)

	index := mgo.Index{
		Key:        []string{"isbn"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	checkErr(err)
}

func allBooks(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()
		c := session.DB(db).C(collection)

		var books []book
		err := c.Find(bson.M{}).All(&books)
		if err != nil {
			errorWithJSON(w, "database error", http.StatusInternalServerError)
			log.Println("failed get all books:", err)
			return
		}

		respBody, err := json.MarshalIndent(books, "", "")
		checkErr(err)

		responseWithJSON(w, respBody, http.StatusOK)
	}
}

func addBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		var b book
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&b)
		if err != nil {
			errorWithJSON(w, "incorrect body", http.StatusBadRequest)
			return
		}

		c := session.DB(db).C(collection)

		err = c.Insert(b)
		if err != nil {
			if mgo.IsDup(err) {
				errorWithJSON(w, "book with this ISBN already exists", http.StatusBadRequest)
				return
			}

			errorWithJSON(w, "database error", http.StatusInternalServerError)
			log.Println("failed insert book:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", r.URL.Path+"/"+b.ISBN)
		w.WriteHeader(http.StatusCreated)
	}
}

func bookByISBN(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		isbn := pat.Param(r, "isbn")
		c := session.DB(db).C(collection)

		var b book
		err := c.Find(bson.M{"isbn": isbn}).One(&b)
		if err != nil {
			errorWithJSON(w, "database error", http.StatusInternalServerError)
			log.Println("failed find book:", err)
			return
		}

		if b.ISBN == "" {
			errorWithJSON(w, "book not found", http.StatusNotFound)
			return
		}

		respBody, err := json.MarshalIndent(b, "", "")
		if err != nil {
			log.Fatal(err)
		}

		responseWithJSON(w, respBody, http.StatusOK)
	}
}

func updateBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		var b book
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&b)
		if err != nil {
			errorWithJSON(w, "incorrect body", http.StatusBadRequest)
			return
		}

		isbn := pat.Param(r, "isbn")
		c := session.DB(db).C(collection)

		err = c.Update(bson.M{"isbn": isbn}, &b)
		if err != nil {
			switch err {
			default:
				errorWithJSON(w, "database error", http.StatusInternalServerError)
				log.Println("failed update book:", err)
				return
			case mgo.ErrNotFound:
				errorWithJSON(w, "book not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func deleteBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		isbn := pat.Param(r, "isbn")
		c := session.DB(db).C(collection)

		err := c.Remove(bson.M{"isbn": isbn})
		if err != nil {
			switch err {
			default:
				errorWithJSON(w, "database error", http.StatusInternalServerError)
				log.Println("failed update book:", err)
				return
			case mgo.ErrNotFound:
				errorWithJSON(w, "book not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
