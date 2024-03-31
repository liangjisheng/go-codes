package main

import (
	"encoding/xml"
	"fmt"
)

type Action struct {
	XMLName     string `xml:"action"`
	Application string `xml:"application,attr"`
	Data        string `xml:",chardata"`
}

type Condition struct {
	XMLName     string `xml:"condition"`
	Field       string `xml:"field,attr"`
	Expresstion string `xml:"expresstion,attr"`
	Actions     []Action
}

type Extension struct {
	XMLName string `xml:"extension"`
	Name    string `xml:"name,attr"`
	Cond    Condition
}

func marshal() {
	var actions []Action
	actions = append(actions, Action{"", "answer", "raw text"})
	actions = append(actions, Action{"", "esf_page_group", ""})
	condition := Condition{"", "destination_number", "^pagegroups$|^7243$", actions}
	extension := Extension{"", "rtp_multicast_page", condition}

	data, _ := xml.MarshalIndent(extension, "", "    ")
	fmt.Printf("%s\n", data)
}
