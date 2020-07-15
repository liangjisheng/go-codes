package v1service

// Add ...
type Add struct {
	A int `json:"a"`
	B int `json:"b"`
}

// AddAck ...
type AddAck struct {
	Res int `json:"res"`
}
