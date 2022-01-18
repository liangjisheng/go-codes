package filedir

import (
	"os"
	"testing"
)

func TestGetFilesAndDirs(t *testing.T) {
	files, _, err := GetFilesAndDirs("/home/ps/image_server/image_data")
	if err != nil {
		t.Error(err)
		return
	}

	//t.Log(files)

	imageFile, err := os.OpenFile("imageFiles.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	defer imageFile.Close()

	for _, file := range files {
		imageFile.WriteString(file + "\n")
	}
}
