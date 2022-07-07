package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// 生成一个数据流
func asStream(done <-chan struct{}) <-chan interface{} {
	s := make(chan interface{})
	values := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(s)
		for _, v := range values { // 从数组生成
			select {
			case <-done:
				return
			case s <- v:
				fmt.Println("in", v)
			}
		}
	}()
	return s
}
func reduce(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil { // 异常检查
		return nil
	}

	out := <-in         // 先读取第一个元素
	for v := range in { // 实现reduce的主要逻辑
		fmt.Println("out", v)
		out = fn(out, v)
	}

	return out
}

func mapChan(in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{}) //创建一个输出chan
	if in == nil {                // 异常检查
		close(out)
		return out
	}

	go func() { // 启动一个goroutine,实现map的主要逻辑
		defer close(out)
		for v := range in { // 从输入chan读取数据，执行业务操作，也就是map操作
			out <- fn(v)
		}
	}()

	return out
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

var client *http.Client

func init() {
	tr := &http.Transport{
		MaxIdleConns:    100,
		IdleConnTimeout: 1 * time.Second,
	}
	client = &http.Client{Transport: tr}
}

type info struct {
	Data interface{} `json:"data"`
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 11; i++ {
		wg.Add(1)
		go func(int2 int) {
			defer wg.Done()
			req, err := http.NewRequest("GET", "http://localhost:8090", nil)
			if err != nil {
				fmt.Printf("初始化http客户端处错误:%v", err)
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("初始化http客户端处错误:%v", err)
				return
			}
			defer resp.Body.Close()
			nByte, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("读取http数据失败:%v", err)
				return
			}
			fmt.Printf("接收到到值:%v\n", string(nByte))
		}(i)
	}
	wg.Wait()

	fmt.Printf("请求完毕\n")
}
