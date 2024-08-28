package main

import (
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
	"strings"
)

func main() {
	// 创建一个自定义的HTTP客户端，忽略证书验证
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://192.168.0.123:9200", // 你的Elasticsearch地址
		},
		Username: "elastic",
		Password: "123sc4567",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// 创建Elasticsearch客户端
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 发出一个简单的请求以验证连接
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	// 读取响应主体
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	} else {
		// 打印响应
		fmt.Println(strings.Repeat("=", 37))
		fmt.Println("Elasticsearch Info:")
		fmt.Println(strings.Repeat("=", 37))
		fmt.Println(res)
	}
}
