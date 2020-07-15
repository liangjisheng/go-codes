package models

import (
	"errors"
	db "go-demos/gin-demo/person-curd/database"
	"log"
)

// Person ...
type Person struct {
	ID        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

// AddPerson ...
func (p *Person) AddPerson() (id int64, err error) {
	if db.SqlDB == nil {
		return -1, errors.New("db is nil")
	}
	// stmt, err := db.SqlDB.Prepare("insert into person(first_name, last_name) values (?, ?)")
	// if err != nil {
	// 	return
	// }
	// res, err := stmt.Exec(p.FirstName, p.LastName)
	// if err != nil {
	// 	return
	// }

	res, err := db.SqlDB.Exec("insert into person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

// UpdatePerson ...
func (p *Person) UpdatePerson() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("UPDATE person SET first_name = ?, last_name = ? WHERE id = ?", p.FirstName, p.LastName, p.ID)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

// DelPerson ...
func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM person WHERE id = ?", p.ID)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

// GetPersons ...
func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("select id, first_name, last_name from person")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

// GetPerson ...
func (p *Person) GetPerson() (err error) {
	log.Println(p.ID)
	db.SqlDB.QueryRow("select id, first_name, last_name from person where id = ?", p.ID).Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
	)
	return
}
