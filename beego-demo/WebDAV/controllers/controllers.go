package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"golang.org/x/net/webdav"
)

var (
	path   = "_testdata"
	prefix = "/"
)

// WebDav WebDav struct
type WebDav struct {
	fs          webdav.FileSystem
	ls          webdav.LockSystem
	HandlerFunc http.HandlerFunc
}

// WebDAVController handles WebDAV requests.
type WebDAVController struct {
	beego.Controller
}

func NewWebDav() *WebDav {
	return &WebDav{}
}

func (wd *WebDav) mount(path string) error {
	if s, err := filepath.Abs(path); err == nil {
		path = s
	}
	wd.fs = webdav.Dir(path)
	wd.ls = webdav.NewMemLS()
	return nil
}

// Main All method handles all requests for WebDAVController.
func (c *WebDAVController) Main() {
	wd := NewWebDav()
	wd.mount(path)

	os.Mkdir(path, os.ModeDir)

	h := &webdav.Handler{
		FileSystem: wd.fs,
		LockSystem: wd.ls,
		Prefix:     prefix,
	}

	h.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request)
}
