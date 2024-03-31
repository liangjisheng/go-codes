package models

import "time"

// Post 主题模型
type Post struct {
	ID        int
	UUID      string
	Body      string
	UserID    int
	ThreadID  int
	CreatedAt time.Time
}

// CreatedAtDate ...
func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// User Get the user who wrote the post
func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", post.UserID).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)
	return
}
