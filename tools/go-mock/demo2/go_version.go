package demo2

import "demo2/spider"

// GetGoVersion ...
func GetGoVersion(s spider.Spider) string {
	return s.GetBody()
}
