package main

// DB ...
type DB interface {
	Get(key string) (int, error)
}

// GetFromDB ...
func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err != nil {
		return value
	}
	return -1
}

func main() {

}
