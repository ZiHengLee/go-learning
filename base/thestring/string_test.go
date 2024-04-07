package thestring

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestFindLongestSubstringLength(t *testing.T) {
	res := longestPalindrome("abba")
	print(res)
}
func Hex2Dec(val string) int {
	bs, _ := hex.DecodeString(val)
	num := binary.BigEndian.Uint16(bs[:2])
	fmt.Println(num)
	return 1
}


func TestOther(t *testing.T) {
	//var a map[int64]int64
	//a= map[int64]int64{}
	//a[1] =1
	//print(a)
	//s := "Hello World!"
	//b := []byte(s)
	//
	//sEnc := base64.StdEncoding.EncodeToString(b)
	//fmt.Printf("enc=[%s]\n", sEnc)
	sEnc:= "cj0xJi1waz02MjM4Njg5MDA2"
	sDec, err := base64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		fmt.Printf("base64 decode failure, error=[%v]\n", err)
	} else {
		fmt.Printf("dec=[%s]\n", sDec)
	}

	a := "r=1&-pk=6238688822"
	y := []byte(a)
	kk := base64.StdEncoding.EncodeToString(y)
	fmt.Printf("dec=[%s]\n", kk)
}

type RequestParam struct {
	Owner    string `json:"owner"`
	Contract string `json:"contract"`
	Default  bool   `json:"default"`
}

// HTTP返回Body
type HTTPRspBody struct {
	Code int          `json:"code"`
	Data *ApproveData `json:"data"`
}
type ApproveData struct {
	Approved bool `json:"approved"`
}

func ApprovePostHttp(req *RequestParam) (result *HTTPRspBody, err error) {
	url := "http://10.13.153.100:6001/opensea/is_approved_all"
	reqParam, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 准备: HTTP请求
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	// Read: HTTP结果
	rspBody, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(rspBody, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func TestApproveHttp(t *testing.T) {
	req := &RequestParam{
		Owner:    "0x641695f6c63d1aa6217e168d0a7657b0f5ffe760",
		Contract: "0x8a198b42ad0966703b927f2d7e396371198479d4",
	}
	resp, _ := ApprovePostHttp(req)
	print(resp.Data.Approved)
}

func TestProxy(t *testing.T){
	uri, err := url.Parse("http://maga2021-zone-custom:maga2021@proxy.ipidea.io:2333")

	if err != nil{
		log.Fatal("parse url error: ", err)
	}
	log.Println(uri.User)

	client := http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(uri),
		},
	}

	resp, err := client.Get("https://ikzttp.mypinata.cloud/ipfs/QmQFkLSQysj94s5GvTHPyzTxrawwtjgiiYS2TBLgrvw8CW/1")
	if err != nil{
		print(err)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(data))
}

func TestFindKthLargest(t *testing.T) {
	a := []int{3,2,1,5,6,4}
	print(heapSort(a,2))
}

func TestPermutation(t *testing.T) {
	a := []int{1,2,3,3}
	retList := make([][]int,0)
	PermutationInt(a,&retList,0,4)
	fmt.Println(retList)
}