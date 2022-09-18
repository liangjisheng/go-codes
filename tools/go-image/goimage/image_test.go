package goimage

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"testing"
)

const (
	dx = 500
	dy = 200
)

func TestImage(t *testing.T) {
	file, err := os.Create("test.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//调用NewAlpha函数，实现image接口
	alpha := image.NewAlpha(image.Rect(0, 0, dx, dy))
	//遍历每一个像素点，设置图片，使用Alpha对图片进行了透明度渐变设定
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			alpha.Set(x, y, color.Alpha{A: uint8(x % 256)}) //设定alpha图片的透明度
		}
	}

	//在这一行的后边，尝试调用上边提到的各种方法
	fmt.Println(alpha.At(400, 100))    //查看在指定位置的像素
	fmt.Println(alpha.Bounds())        //查看图片边界
	fmt.Println(alpha.Opaque())        //查看是否图片完全透明
	fmt.Println(alpha.PixOffset(1, 1)) //查看指定点相对于第一个点的距离
	fmt.Println(alpha.Stride)          //查看两个垂直像素之间的距离
	jpeg.Encode(file, alpha, nil)      //将alpha中的信息写入图片文件中
}

func TestImage1(t *testing.T) {
	file, err := os.Create("test1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//新建图片接口起始点是0,0，定义矩形边界
	rgba := image.NewRGBA(image.Rect(0, 0, dx, dy))
	//为图片添加颜色
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			//set函数设定(x,y)的color
			//NRGBA，保存像素信息
			rgba.Set(x, y, color.NRGBA{uint8(x % 256), uint8(y % 256), 0, 255})
		}
	}

	fmt.Println(rgba.At(400, 100))    //{144 100 0 255}
	fmt.Println(rgba.Bounds())        //(0,0)-(500,200)
	fmt.Println(rgba.Opaque())        //true，其完全透明
	fmt.Println(rgba.PixOffset(1, 1)) //2004
	fmt.Println(rgba.Stride)          //2000
	jpeg.Encode(file, rgba, nil)      //将image信息存入文件中
}

func TestImage2(t *testing.T) {
	file, err := os.Create("test2.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//新建图片接口起始点是0,0，定义矩形边界
	rgba := image.NewRGBA(image.Rect(0, 0, dx, dy))
	//为图片添加颜色
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			//set函数设定(x,y)的color
			//NRGBA，保存像素信息
			rgba.Set(x, y, color.NRGBA{240, 255, 255, 255})
		}
	}
	jpeg.Encode(file, rgba, nil) //将image信息存入文件中
}
