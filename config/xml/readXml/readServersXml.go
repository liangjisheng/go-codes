package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type RecurlyServers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName     xml.Name `xml:"server"`
	ServerName  string   `xml:"serverName"`
	ServerIP    string   `xml:"serverIP"`
	Description string   `xml:",innerxml"`
}

func readServersXml() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	v := RecurlyServers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Println(v.XMLName)
	fmt.Println(v.Version)
	fmt.Println(v.Svs[0].Description)
}
