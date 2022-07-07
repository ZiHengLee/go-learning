package others

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"
)

//nums1 := []int{1,2}
//nums2 := []int{1,2,3,4}
//func TestOthers(t *testing.T) {
//	fmt.Println(isValid("()[]{}"))
//	//fmt.Println(isValid("(()()())())"))
//}
func TestThreeNum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(nums[:6])
	//fmt.Println(threeSum(nums))
}
func TestFindMedianSortedArrays(t *testing.T) {
	tt := time.Now().UTC()
	//fmt.Println(t)
	// 获取当前时间戳
	timestamp := tt.Unix()
	fmt.Println(timestamp)
}

func TestEmoji(t *testing.T) {
	//before,_ := CalculateZeroTimeStamp(0)
	weekBefore, _ := GetFirstDateOfWeek()
	//fmt.Println(before-86400==weekBefore)
	//fmt.Println(GetDateString(weekBefore,before-86400))
	//featuredCount :=1
	//caculateCount:=3
	//featuredRateStr := fmt.Sprintf("%.2f", 100*(float64(featuredCount)/float64(caculateCount)))
	//featuredRate, _ := strconv.ParseFloat(featuredRateStr, 64)
	fmt.Println(weekBefore)
}

func GetDateString(before, after int64) string {
	beforeTime := time.Unix(before, 0)
	_, bm, bd := beforeTime.Date()
	afterTime := time.Unix(after, 0)
	_, am, ad := afterTime.Date()
	strDate := fmt.Sprintf("%d.%d-%d.%d", int(bm), bd, int(am), ad)
	return strDate
}
func TestE(t *testing.T) {
	//timeStr := time.Now().Format("2006-01-02")
	//fmt.Println("timeStr:", timeStr)
	//loc,_ := time.LoadLocation("America/New_York")
	//
	//ss, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	//timeNumber := ss.Unix()
	//fmt.Println(CalculateZeroTimeStamp(1))  //1578844800
	//fmt.Println(CalculateZeroTimeStamp(1))
	fmt.Println(getAwemeId("https://www.tiktok.com/@maggieboodles"))
}

//tiktok资源id计算
func getAwemeId(url string) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	resp, err := client.Do(req)
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return ""
	}
	fmt.Println(resp)
	url = resp.Header.Get("Location")
	print(url)
	if strings.Contains(url, ".html") {
		sIndex := strings.LastIndex(url, "/")
		eIndex := strings.LastIndex(url, ".html")
		if sIndex < eIndex && eIndex < len(url) {
			return url[sIndex+1 : eIndex]
		}
	} else {
		t := strings.Split(url, "/")
		if len(t) > 2 {
			return t[len(t)-2]
		}
	}
	return ""
}

func GetFirstDateOfWeek() (before, after int64) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	loc := time.FixedZone("CST", -4*60*60)
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, offset)
	before = weekStartDate.Unix() - 86400
	return
}
func CalculateZeroTimeStamp(limit int) (start, end int64) {
	timeStr := time.Now().Format("2006-01-02")
	//loc, _ := time.LoadLocation("America/New_York")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	timeNumber := t.Unix()
	return timeNumber, timeNumber + int64(limit*86400)
}

func get_day_continuous(ts int64) (dayContinuous int64, err error) {

	// 时区确认: 东七区
	var cstZone = time.FixedZone("CST", 60*60*7)
	t1 := time.Date(time.Now().In(cstZone).Year(), time.Now().In(cstZone).Month(), time.Now().In(cstZone).Day(), 0, 0, 0, 0, cstZone)
	t2 := time.Date(time.Unix(ts, 0).In(cstZone).Year(), time.Unix(ts, 0).In(cstZone).Month(), time.Unix(ts, 0).In(cstZone).Day(), 0, 0, 0, 0, cstZone)

	dayContinuous = int64(t1.Sub(t2).Hours() / 24)

	return
}
