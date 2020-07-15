package v4service

// Add ...
type Add struct {
	A int `json:"a"`
	B int `json:"b"`
}

// AddAck ...
type AddAck struct {
	Res int `json:"res"`
}

// Login ...
type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// LoginAck ...
type LoginAck struct {
	Token string `json:"token"`
}
