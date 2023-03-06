package main

import (
	"context"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp/device"
	"io/ioutil"
	"log"
	"time"

	cdp "github.com/chromedp/chromedp"
)

func main() {
	// 创建新的cdp上下文
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()

	urlstr := `http://wk5.bookan.com.cn/index.html?id=27058&token=&productId=5#/`
	var buf []byte
	if err := cdp.Run(ctx, cdp.Emulate(device.IPhone13ProMax), wkScreenshot(urlstr, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("1.png", buf, 0644); err != nil {
		log.Fatal(err)
	}

	urlstr = `https://zq5.bookan.com.cn/?t=index&id=20954#/`
	if err := cdp.Run(ctx, cdp.Emulate(device.IPhone13ProMax), zqScreenshot(urlstr, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("2.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}

func wkScreenshot(urlstr string, res *[]byte) cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate(urlstr),
		cdp.WaitVisible("#app", cdp.ByID),
		cdp.Sleep(time.Duration(3) * time.Second),
		cdp.Screenshot("#app", res, cdp.NodeVisible, cdp.ByID),
	}
}

func zqScreenshot(urlstr string, quality int, res *[]byte) cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate(urlstr),
		cdp.WaitVisible("#app", cdp.ByID),
		cdp.ActionFunc(scroll("window.scrollTo(0, 1200);")),
		cdp.Sleep(time.Duration(3) * time.Second),
		cdp.ActionFunc(scroll("window.scrollTo(0, 2400);")),
		cdp.Sleep(time.Duration(3) * time.Second),
		cdp.ActionFunc(scroll("window.scrollTo(0, document.body.scrollHeight);")),
		cdp.Sleep(time.Duration(3) * time.Second),
		cdp.FullScreenshot(res, quality),
	}
}

func scroll(script string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		_, exp, err := runtime.Evaluate(script).Do(ctx)
		if err != nil {
			return err
		}
		if exp != nil {
			return exp
		}
		return nil
	}
}
