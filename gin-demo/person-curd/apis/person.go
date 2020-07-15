package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go-demos/gin-demo/person-curd/models"

	"github.com/gin-gonic/gin"
)

// Index ...
func Index(c *gin.Context) {
	c.String(http.StatusOK, "It works!")
}

// AddPerson ...
func AddPerson(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	log.Println(firstName, lastName)
	p := models.Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

// UpdatePerson ...
func UpdatePerson(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	p := models.Person{ID: id}

	p.GetPerson()
	if p.FirstName != "" {
		p.FirstName = firstName
		p.LastName = lastName
		ra, err := p.UpdatePerson()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("update successful %d", ra)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

// DelPerson ...
func DelPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	p := models.Person{ID: id}
	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	} else {
		msg := fmt.Sprintf("delete successful %d", ra)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	}
}

// GetPerson ...
func GetPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	p := models.Person{ID: id}
	p.GetPerson()
	if p.FirstName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": p,
			"msg":  "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

// GetPersons ...
func GetPersons(c *gin.Context) {
	var p models.Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": persons,
		"msg":  "success",
	})
}
