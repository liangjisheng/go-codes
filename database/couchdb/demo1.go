package main

// create your own document
type dummyDocument struct {
	// couchdb.Document
	Foo  string `json:"foo"`
	Beep string `json:"beep"`
}

/*func demo1() {
	u, err := url.Parse("http://127.0.0.1:5984/")
	if err != nil {
		panic(err)
	}

	// create a new client
	client, err := couchdb.NewClient(u)
	if err != nil {
		panic(err)
	}

	// get some information about your CouchDB
	info, err := client.Info()
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}*/
