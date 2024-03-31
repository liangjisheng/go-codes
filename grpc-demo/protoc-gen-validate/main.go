package main

import (
	"log"
	examplepb "protocvalidate/proto"
)

func main() {
	p := new(examplepb.Person)

	err := p.Validate() // err: Id must be greater than 999
	log.Println(err)
	p.Id = 1000

	err = p.Validate() // err: Email must be a valid email address
	log.Println(err)
	p.Email = "example@web3.com"

	err = p.Validate() // err: Name must match pattern '^[^\d\s]+( [^\d\s]+)*$'
	log.Println(err)
	p.Name = "Protocol Buffer"

	err = p.Validate() // err: Home is required
	log.Println(err)
	p.Home = &examplepb.Person_Location{
		Lat: 37.7,
		Lng: 999,
	}

	err = p.Validate() // err: Home.Lng must be within [-180, 180]
	log.Println(err)
	p.Home.Lng = -122.4

	err = p.Validate() // err: nil
	log.Println(err)
}
