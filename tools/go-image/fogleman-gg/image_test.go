package fogleman_gg

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"testing"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//https://blog.csdn.net/baidu_32452525/article/details/119194879

func TestImage(t *testing.T) {
	waterImage, err := gg.LoadPNG("out.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 通过图片实例化gg
	dc := gg.NewContextForImage(waterImage)
	dc.SetRGBA(1, 1, 1, 0)

	dc.SetRGB(0, 0, 0)
	// 加载字体，设置大小
	if err := dc.LoadFontFace("Uchen-Regular.ttf", 96); err != nil {
		fmt.Println(err.Error())
		return
	}
	// 给图片添加文字，位置在（x, y） 处
	s := "alice"
	//dc.DrawStringWrapped(s, 10, 25, 0, 0, float64(dc.Width())*0.9, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(s, 10, 25, 0, 0, float64(dc.Width())*0.9, 1.5, gg.AlignCenter)

	newFile, err := os.Create("alice.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer newFile.Close()

	// 将文件保存输出，并设置压缩比
	err = jpeg.Encode(newFile, dc.Image(), &jpeg.Options{Quality: 60})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

//otf 字体文件加载
//LoadFontFace() 这个方法只能加载 ttf 字体文件，也就是 true type font，无法加载 otf 字体文件，也就是 open type font
//所以如果需要加载 otf 字体文件，则需要换一个姿势

func getOpenTypeFontFace(fontFilePath string, fontSize, dpi float64) (*font.Face, error) {
	fontData, fontFileReadErr := ioutil.ReadFile(fontFilePath)
	if fontFileReadErr != nil {
		return nil, fontFileReadErr
	}
	otfFont, parseErr := opentype.Parse(fontData)
	if parseErr != nil {
		return nil, parseErr
	}
	otfFace, newFaceErr := opentype.NewFace(otfFont, &opentype.FaceOptions{
		Size: fontSize,
		DPI:  dpi,
	})
	if newFaceErr != nil {
		return nil, newFaceErr
	}
	return &otfFace, nil
}
