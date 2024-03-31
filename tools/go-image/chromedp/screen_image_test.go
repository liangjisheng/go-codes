package chromedp__test

import (
	"context"
	"io/ioutil"
	"log"
	"testing"

	"github.com/chromedp/chromedp"
)

func TestImage(t *testing.T) {
	// 生成 chrome 实例
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// 设置截图某一个元素
	var buf []byte
	if err := chromedp.Run(ctx, elementScreenshot(`https://pkg.go.dev/`, `img.Homepage-logo`, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	// 全部截取，设置压缩比为90%
	if err := chromedp.Run(ctx, fullScreenshot(`https://brank.as/`, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
}

// elementScreenshot 截图页面某一个元素
func elementScreenshot(url, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// fullScreenshot 截图整个屏幕
func fullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.FullScreenshot(res, quality),
	}
}
