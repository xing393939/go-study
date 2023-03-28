package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	op = "https://opbackstage.cn/opbackstage"
	rk = "/answerfun/rankPersonList?order_by=passed_levels%20desc,used_time_sum%20asc,update_time%20asc"
	rp = "https://bizcache7n.cn/url/html/yyzcharts/index.html?actuniqid="
)

func main() {
	// 创建新的cdp上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	str, _ := os.Getwd()

	org10Map := split(str + "/Sheet1.csv")
	orgAllMap := split(str + "/Sheet2.csv")

	//rank(org10Map, orgAllMap, ctx)
	report(org10Map, orgAllMap, ctx)
}

func rank(org10Map, orgAllMap map[string][3]string, ctx context.Context) {
	done := make(chan string)
	chromedp.ListenTarget(ctx, func(v interface{}) {
		if ev, ok := v.(*browser.EventDownloadProgress); ok {
			completed := "(unknown)"
			if ev.TotalBytes != 0 {
				completed = fmt.Sprintf("%0.2f%%", ev.ReceivedBytes/ev.TotalBytes*100.0)
			}
			log.Printf("state: %s, completed: %s\n", ev.State.String(), completed)
			if ev.State == browser.DownloadProgressStateCompleted {
				done <- ev.GUID
			}
		}
	})

	str, _ := os.Getwd()
	for _, row := range org10Map {
		dir := str + "\\" + row[2]
		_ = os.MkdirAll(dir, 0777)
		name := dir + "/" + row[1] + ".xlsx"
		if PathExists(name) {
			log.Println(name + " exists")
			continue
		}
		actOrgId := orgAllMap[row[0]][2]
		if err := chromedp.Run(ctx,
			browser.SetDownloadBehavior(browser.SetDownloadBehaviorBehaviorAllowAndName).
				WithDownloadPath(dir).
				WithEventsEnabled(true),
			chromedp.Navigate(op+rk+"&offset=0&limit=0&action_flag=export&actorgid="+actOrgId),
		); err != nil {
			log.Println(err)
		}
		guid := <-done
		log.Println(guid)
		fs, _ := ioutil.ReadDir(dir)
		for _, file := range fs {
			if !file.IsDir() && !strings.Contains(file.Name(), ".xlsx") {
				_ = os.Rename(dir+"/"+file.Name(), name)
			}
		}
	}
}

func report(org10Map, orgAllMap map[string][3]string, ctx context.Context) {
	var res []byte
	str, _ := os.Getwd()
	for _, row := range org10Map {
		dir := str + "\\" + row[2]
		_ = os.MkdirAll(dir, 0777)
		name := dir + "/" + row[1] + ".png"
		if PathExists(name) {
			log.Println(name + " exists")
			continue
		}
		actOrgId := orgAllMap[row[0]][1]
		if err := chromedp.Run(ctx,
			chromedp.Emulate(device.IPhone13ProMax),
			chromedp.Navigate(rp+actOrgId),
			chromedp.WaitVisible("img.down", chromedp.ByQuery),
			chromedp.ActionFunc(scroll("document.querySelector('img.down').style.display='none';")),
			chromedp.Sleep(time.Duration(1)*time.Second),
			chromedp.FullScreenshot(&res, 90),
		); err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(name, res, 0644); err != nil {
			log.Fatal(err)
		} else {
			log.Println(name)
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

func split(filename string) map[string][3]string {
	org10str, _ := ioutil.ReadFile(filename)
	org10arr := strings.Split(string(org10str), "\n")
	rs := make(map[string][3]string)
	for _, str := range org10arr {
		row := strings.Split(str, ",")
		if len(row) == 3 {
			rs[row[0]] = [3]string{
				row[0],
				strings.TrimSpace(row[1]),
				strings.TrimSpace(row[2]),
			}
		}
	}
	return rs
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
