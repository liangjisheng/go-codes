package fogleman_gg

import (
	"bytes"
	"github.com/fogleman/gg"
	"testing"
)

func TestTextImage(t *testing.T) {
	const dx = 1024
	const dy = 384
	dc := gg.NewContext(dx, dy)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	// font: /Library/Fonts/Arial Unicode.ttf
	if err := dc.LoadFontFace("ArialUnicode.ttf", 60); err != nil {
		panic(err)
	}

	//dc.DrawStringAnchored("Hello, world, banana.", 0, dc.FontHeight(), 0, 0)
	dc.DrawStringAnchored("4757648372483478349824784", dx/2, dy/2, 0.5, 0.5)
	//dc.DrawStringAnchored("Hello, world, banana.", dx, dy-dc.FontHeight(), 1, 1)

	buf := new(bytes.Buffer)
	dc.EncodePNG(buf)
	t.Log("len", buf.Len())

	dc.SavePNG("out.png")
}
