package qrcode

//https://github.com/skip2/go-qrcode/blob/master/example_test.go

import (
	"fmt"
	"image/color"
	"os"
	"testing"

	"github.com/skip2/go-qrcode"
)

func TestExampleEncode(t *testing.T) {
	if png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256); err != nil {
		t.Errorf("Error: %s", err.Error())
	} else {
		fmt.Printf("PNG is %d bytes long", len(png))
	}
}

func TestExampleWriteFile(t *testing.T) {
	filename := "example.png"
	if err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, filename); err != nil {
		if err = os.Remove(filename); err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	}
}

func TestExampleEncodeWithColourAndWithoutBorder(t *testing.T) {
	q, err := qrcode.New("https://example.org", qrcode.Medium)
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	// Optionally, disable the QR Code border.
	q.DisableBorder = true

	// Optionally, set the colours.
	q.ForegroundColor = color.RGBA{R: 0x33, G: 0x33, B: 0x66, A: 0xff}
	q.BackgroundColor = color.RGBA{R: 0xef, G: 0xef, B: 0xef, A: 0xff}

	err = q.WriteFile(256, "example2.png")
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}
}
