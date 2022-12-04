package captcha

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestImgText(t *testing.T) {
	img := ImgText(800, 200, "hello")

	file, _ := os.Create("img.png")
	defer file.Close()

	_, _ = io.Copy(file, bytes.NewBuffer(img))
}
