package spf13_afero__test

import (
	"github.com/spf13/afero"
	"log"
	"os"
	"testing"
)

func TestDemo(t *testing.T) {
	var appFs = afero.NewOsFs()
	log.Println(appFs.Name())
	aFile, _ := appFs.Create("a.txt")
	defer aFile.Close()
	_, _ = aFile.Write([]byte("test"))

	aFileInfo, _ := appFs.Stat("a.txt")
	log.Println(aFileInfo.Name())
	log.Println(aFileInfo.Mode())
	log.Println(aFileInfo.IsDir())

	exist, _ := afero.DirExists(appFs, "/root")
	log.Println(exist)
}

func TestExist(t *testing.T) {
	appFS := afero.NewMemMapFs()
	// create test files and directories
	appFS.MkdirAll("src/a", 0755)
	afero.WriteFile(appFS, "src/a/b", []byte("file b"), 0644)
	afero.WriteFile(appFS, "src/c", []byte("file c"), 0644)
	name := "src/c"
	_, err := appFS.Stat(name)
	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", name)
	}
}
