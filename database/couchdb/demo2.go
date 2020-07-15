package main

import (
	"fmt"
	"time"

	"github.com/joelnb/sofa"
)

func demo2() {
	conn, err := sofa.NewConnection("http://localhost:5984", 10*time.Second, sofa.NullAuthenticator())
	if err != nil {
		panic(err)
	}

	db := conn.Database("new_database")
	doc := &struct {
		sofa.DocumentMetadata
		Name string `json:"name"`
		Type string `json:"type"`
	}{
		DocumentMetadata: sofa.DocumentMetadata{
			ID: "fruit1",
		},
		Name: "apple",
		Type: "fruit",
	}

	rev, err := db.Put(doc)
	if err != nil {
		panic(err)
	}
	fmt.Println(rev)
}
