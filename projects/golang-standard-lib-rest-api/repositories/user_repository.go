package repositories

import (
	"database/sql"
	"go-demos/projects/golang-standard-lib-rest-api/models"
	"go-demos/projects/golang-standard-lib-rest-api/utils/crypto"
)

// GetUserByID ...
func GetUserByID(db *sql.DB, id int) (*models.User, error) {
	const query = `
		select
			id,
			email,
			name
		from
			users
		where
			id = $1
	`
	var user models.User
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name)
	return &user, err
}

// GetUserByEmail ...
func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	const query = `
		select
			id,
			email,
			name
		from
			users
		where
			email = $1
	`
	var user models.User
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name)
	return &user, err
}

// GetPrivateUserDetailsByEmail ...
func GetPrivateUserDetailsByEmail(db *sql.DB, email string) (*models.PrivateUserDetails, error) {
	const query = `
		select
			id,
			password,
			salt
		from
			users
		where
			email = $1
	`
	var u models.PrivateUserDetails
	err := db.QueryRow(query, email).Scan(&u.ID, &u.Password, &u.Salt)
	return &u, err
}

// CreateUser create user
func CreateUser(db *sql.DB, email, name, password string) (int, error) {
	const query = `
		insert into users (
			email,
			name,
			password,
			salt
		) values (
			$1,
			$2,
			$3,
			$4
		) returning id
	`

	salt := crypto.GenerateSalt()
	hashedPassword := crypto.HashPassword(password, salt)
	var id int
	err := db.QueryRow(query, email, name, hashedPassword, salt).Scan(&id)
	return id, err
}
