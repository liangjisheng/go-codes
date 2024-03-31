package stringimage

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/golang/freetype"
)

func TestImage(t *testing.T) {
	createImage("alice", "jpg")
	//drawAscii()
}

const (
	dx = 200
	dy = 200
)

func createImage(textName string, fileType string) {
	imgFile, _ := os.Create(textName + "." + fileType)
	defer imgFile.Close()
	//创建位图,坐标x,y,长宽x,y
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	//读字体数据
	//cp /System/Library/Fonts/*.ttf .
	fontBytes, err := ioutil.ReadFile("SFNS.ttf")
	if err != nil {
		log.Println(err)
		return
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(40)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	//c.SetSrc(image.White)
	c.SetSrc(image.Black)
	//设置字体显示位置
	//pt := freetype.Pt(5, 20+int(c.PointToFixed(40)>>8))
	//pt := freetype.Pt(dx/2, dy/2)
	//_, err = c.DrawString(textName, pt)
	if err != nil {
		log.Println(err)
		return
	}

	switch fileType {
	case "png":
		err = png.Encode(imgFile, img)
	case "jpg", "jpeg":
		err = jpeg.Encode(imgFile, img, nil)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func drawAscii() {
	//灰度替换字符
	base := "@#&$%*o!;."
	file1, _ := os.Open("alice.png") //图像名称
	image1, _ := png.Decode(file1)
	bounds := image1.Bounds() //获取图像的边界信息
	logo := ""                //存储最终的字符画string
	for y := 0; y < bounds.Dy(); y += 2 {
		for x := 0; x < bounds.Dx(); x++ {
			pixel := image1.At(x, y)   //获取像素点
			r, g, b, _ := pixel.RGBA() //获取像素点的rgb
			r = r & 0xFF
			g = g & 0xFF
			b = b & 0xFF
			//灰度计算
			gray := 0.299*float64(r) + 0.578*float64(g) + 0.114*float64(b)
			temp := fmt.Sprintf("%.0f", gray*float64(len(base)+1)/255)
			index, _ := strconv.Atoi(temp)
			//根据灰度索引字符并保存
			if index >= len(base) {
				logo += " "
			} else {
				logo += string(base[index])
			}
		}
		logo += "\r\n"
	}
	file1.Close()
	//输出字符画
	log.Printf("\033[31;1m%s , %d", logo, len(logo))
}
