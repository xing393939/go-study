package main

import (
	"context"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp/device"
	"io/ioutil"
	"log"
	"os"
	"time"

	cdp "github.com/chromedp/chromedp"
)

var list = [][3]string{
	//{"华润医药控股有限公司", "微刊", "47865"},
}

func main() {
	// 创建新的cdp上下文
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()

	day := time.Now().Format("2006-01-02")
	_ = os.MkdirAll(day, 0711)
	log.Println("start...")

	var buf []byte
	for _, row := range list {
		file := day + "/" + row[0] + "_" + row[1] + "_" + row[2] + "_" + day + ".png"
		if PathExists(file) {
			log.Println(file + " exists")
			continue
		}
		if row[1] == "微刊" {
			getWk(row[2], &buf, ctx)
		} else {
			getZq(row[2], &buf, ctx)
		}
		if err := ioutil.WriteFile(file, buf, 0644); err != nil {
			log.Fatal(err)
		} else {
			log.Println(file)
		}
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks cdp.Tasks) cdp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}

func getWk(id string, res *[]byte, ctx context.Context) {
	url := `http://wk5.bookan.com.cn/index.html?id=` + id
	if err := cdp.Run(
		ctx,
		cdp.Emulate(device.IPhone13ProMax),
		RunWithTimeOut(&ctx, 5, cdp.Tasks{
			cdp.Navigate(url),
			cdp.WaitVisible("#app", cdp.ByID),
			cdp.Sleep(time.Duration(3) * time.Second),
			cdp.FullScreenshot(res, 90),
		}),
	); err != nil {
		log.Fatal(err)
	}
}

func getZq(id string, res *[]byte, ctx context.Context) {
	url := `https://zq5.bookan.com.cn/?t=index&id=` + id
	if err := cdp.Run(
		ctx,
		cdp.Emulate(device.IPhone13ProMax),
		RunWithTimeOut(&ctx, 12, cdp.Tasks{
			cdp.Navigate(url),
			cdp.WaitVisible("#app", cdp.ByID),
			cdp.ActionFunc(scroll("window.scrollTo(0, 1200);")),
			cdp.Sleep(time.Duration(3) * time.Second),
			cdp.ActionFunc(scroll("window.scrollTo(0, 2400);")),
			cdp.Sleep(time.Duration(3) * time.Second),
			cdp.ActionFunc(scroll("window.scrollTo(0, document.body.scrollHeight);")),
			cdp.Sleep(time.Duration(3) * time.Second),
			cdp.FullScreenshot(res, 90),
		}),
	); err != nil {
		log.Fatal(err)
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
