package utils

import "fmt"

// KeyOfSet KeyOfSet
func KeyOfSet(uid int32) string {
	return fmt.Sprintf("t:uid:set:%d", uid)
}

// KeyOfSession KeyOfSession
func KeyOfSession(uid int32) string {
	return fmt.Sprintf("t:%d", uid)
}

// KeyOfToken KeyOfToken
func KeyOfToken(token string) string {
	return fmt.Sprintf("t:%s", token)
}
