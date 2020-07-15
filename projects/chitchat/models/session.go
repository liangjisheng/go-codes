package models

import "time"

// Session ...
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?", session.UUID).
		Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.ID != 0 {
		valid = true
	}
	return
}

// DeleteByUUID Delete session from database
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.UUID)
	return
}

// User Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", session.UserID).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// SessionDeleteAll Delete all sessions from database
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}
