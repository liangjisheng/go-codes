package main

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func main() {
	imgPath := "example.png"
	// 打印图像的元数据
	md := metaData(imgPath)
	fmt.Printf("图像类型：%v\n图像尺寸：%v x %v\n", md.Type, md.Size.Width, md.Size.Height)

	// 调整图像大小
	resize(imgPath, 1200, 800)

	// 旋转图像,传旋转角度
	rotate(imgPath, 180)

	// 自动(随机)图像,根据 EXIF 方向元数据(autoRotateOnly)自动旋转
	autoRotate(imgPath)

	// 格式转换
	convert(imgPath, bimg.JPEG)

	// 添加文字水印
	watermarkText(imgPath)

	// 添加图片水印
	logoPath := "logo.png"
	watermarkImage(imgPath, logoPath)

	// 高斯模糊,模糊程度参数
	gaussianBlur(imgPath, 20, 5)

	// 按长宽式缩略
	smartCrop(imgPath, 400, 300)

	// 缩略图,缩略图像素 200 * 200
	thumbnail(imgPath, 200)

	// 获取图像类型
	imgType := getType(imgPath)
	fmt.Printf("图像类型：%v\n", imgType)
}

func metaData(path string) bimg.ImageMetadata {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	md, err := bimg.Metadata(buffer)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return md
}

func resize(path string, width, height int) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Resize(width, height)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	_, err = bimg.NewImage(newImage).Size()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("001-resize.jpg", newImage)
}

func rotate(path string, angle bimg.Angle) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// 传旋转角度
	newImage, err := bimg.NewImage(buffer).Rotate(angle)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("002-rotate.png", newImage)
}

func autoRotate(path string) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).AutoRotate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("003-auto-rotate.jpg", newImage)
}

func convert(path string, imageType bimg.ImageType) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Convert(imageType)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("004-convert."+bimg.ImageTypes[imageType], newImage)

}

func watermarkText(path string) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// 水印的相关数据
	watermark := bimg.Watermark{
		Text:       "GoCN 社区",
		DPI:        150,
		Margin:     300,
		Opacity:    0.25, // 不透明度
		Width:      500,
		Font:       "sans bold 14",
		Background: bimg.Color{255, 255, 255},
	}

	newImage, err := bimg.NewImage(buffer).Watermark(watermark)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("005-watermark.jpg", newImage)
}

func watermarkImage(path, logo string) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	logoBuffer, err := bimg.Read(logo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// 水印的相关数据
	watermark := bimg.WatermarkImage{
		Left:    100,
		Top:     200,
		Buf:     logoBuffer,
		Opacity: 1, // 不透明度 0~1 之间的浮点数
	}

	newImage, err := bimg.NewImage(buffer).WatermarkImage(watermark)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("006-watermark-image.jpg", newImage)
}

func gaussianBlur(path string, sigma, minAmpl float64) {

	options := bimg.Options{
		GaussianBlur: bimg.GaussianBlur{sigma, minAmpl},
	}

	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("007-gaussianblur.jpg", newImage)
}

func smartCrop(path string, width, height int) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).SmartCrop(width, height)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("008-smartCrop.jpg", newImage)
}

func thumbnail(path string, pixels int) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Thumbnail(pixels)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("009-thumbnail.jpg", newImage)
}

func getType(path string) string {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	typeName := bimg.NewImage(buffer).Type()

	return typeName
}
