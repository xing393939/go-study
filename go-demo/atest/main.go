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

var list = [][3]string{
	{"国家电网有限公司", "微刊", "762"},
	{"国家电网有限公司", "微刊", "34241"},
	{"国家电网有限公司", "微刊", "49154"},
	{"国家电网有限公司", "微刊", "51227"},
	{"中国民航工会", "专区", "22210"},
	{"中国民航工会", "微刊", "21025"},
	{"天津高速公路集团有限公司", "专区", "26293"},
	{"天津高速公路集团有限公司", "微刊", "27476"},
	{"中国煤炭科工集团有限公司", "专区", "26702"},
	{"中国煤炭科工集团有限公司", "微刊", "26124"},
	{"中国工商银行北京市分行", "专区", "41131"},
	{"中国工商银行北京市分行", "微刊", "41133"},
	{"中国铁建股份有限公司", "专区", "20927"},
	{"中国铁建股份有限公司", "微刊", "17543"},
	{"中国铁建股份有限公司", "微刊", "26175"},
	{"中国铁建股份有限公司", "微刊", "26176"},
	{"中国铁建股份有限公司", "微刊", "26177"},
	{"中国铁建股份有限公司", "微刊", "26178"},
	{"中国铁建股份有限公司", "微刊", "26179"},
	{"中国铁建股份有限公司", "微刊", "26180"},
	{"中国铁建股份有限公司", "微刊", "26181"},
	{"中国铁建股份有限公司", "微刊", "26182"},
	{"北京市发展和改革委员会", "专区", "26826"},
	{"北京市发展和改革委员会", "微刊", "26823"},
	{"北京城市副中心投资建设集团有限公司", "专区", "22527"},
	{"北京城市副中心投资建设集团有限公司", "微刊", "22529"},
	{"黑龙江省总工会", "专区", "28525"},
	{"黑龙江省总工会", "微刊", "28526"},
	{"中国北方化学研究院集团有限公司", "专区", "49696"},
	{"中国北方化学研究院集团有限公司", "微刊", "25913"},
	{"天津港集团有限公司", "专区", "21638"},
	{"天津港集团有限公司", "微刊", "21634"},
	{"华润医药控股有限公司", "专区", "30086"},
	{"华润医药控股有限公司", "微刊", "30088"},
	{"华润医药控股有限公司", "微刊", "47865"},
}

func main() {
	// 创建新的cdp上下文
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	for _, row := range list {
		if row[1] == "微刊" {
			getWk(row[2], &buf, ctx)
		} else {
			getZq(row[2], &buf, ctx)
		}
		file := "dist/" + row[0] + "_" + row[1] + "_" + row[2] + "_" + time.Now().Format("2006-01-02")
		if err := ioutil.WriteFile(file+".png", buf, 0644); err != nil {
			log.Fatal(err)
		} else {
			log.Println(file)
		}
	}
}

func getWk(id string, res *[]byte, ctx context.Context) {
	url := `http://wk5.bookan.com.cn/index.html?id=` + id
	if err := cdp.Run(
		ctx,
		cdp.Emulate(device.IPhone13ProMax),
		cdp.Navigate(url),
		cdp.WaitVisible("#app", cdp.ByID),
		cdp.Sleep(time.Duration(3)*time.Second),
		cdp.FullScreenshot(res, 90),
	); err != nil {
		log.Fatal(err)
	}
}

func getZq(id string, res *[]byte, ctx context.Context) {
	url := `https://zq5.bookan.com.cn/?t=index&id=` + id
	if err := cdp.Run(
		ctx,
		cdp.Emulate(device.IPhone13ProMax),
		cdp.Navigate(url),
		cdp.WaitVisible("#app", cdp.ByID),
		cdp.ActionFunc(scroll("window.scrollTo(0, 1200);")),
		cdp.Sleep(time.Duration(3)*time.Second),
		cdp.ActionFunc(scroll("window.scrollTo(0, 2400);")),
		cdp.Sleep(time.Duration(3)*time.Second),
		cdp.ActionFunc(scroll("window.scrollTo(0, document.body.scrollHeight);")),
		cdp.Sleep(time.Duration(3)*time.Second),
		cdp.FullScreenshot(res, 90),
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
