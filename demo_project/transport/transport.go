package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

func GetClient() (httpCli *http.Client) {
	client := new(http.Client)
	//uri, _ := url.Parse("http://maga2021-zone-custom:maga2021@proxy.ipidea.io:2333")
	var transport http.RoundTripper = &http.Transport{
		//Proxy:              http.ProxyURL(uri),
		DisableKeepAlives:  false,
		DisableCompression: false,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 10 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          1,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   30 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
	}
	client.Transport = transport
	return client
}

func main() {
	client := GetClient()
	wg := sync.WaitGroup{}
	//fmt.Println(client.Transport)
	s := 0
	for {
		wg.Add(10)
		s += 1
		for i := 0; i < 10; i++ {
			go func(j int) {
				_, err := client.Get("https://www.baidu.com")
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(j)
				time.Sleep(time.Second * 5)
				wg.Done()
			}(i)
		}
		fmt.Println()
		wg.Wait()
		if s == 1000 {
			break
		}
	}
	time.Sleep(time.Second * 5)
	res, err := client.Get("https://www.baidu.com")
	fmt.Println(res, err)
	time.Sleep(time.Second * 60)
}
