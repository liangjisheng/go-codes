package requests

// RegisterRequest ...
type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// LoginRequest ...
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateJobRequest ...
type CreateJobRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdateJobRequest ...
type UpdateJobRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
