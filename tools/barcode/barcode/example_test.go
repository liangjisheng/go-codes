package barcode

import (
	"image/png"
	"os"
	"testing"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func TestQREncode(t *testing.T) {
	// Create the barcode
	qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
