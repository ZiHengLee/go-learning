package 代理模式

import "fmt"

type Seller interface {
	sell(name string)
}

type Station struct {
	stock int //库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
	}
}

func (station *Station) remain() bool {
	if station.stock > 0 {
		return true
	}
	return false
}

type StationProxy struct {
	station *Station
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.remain() {
		proxy.station.sell(name)
	} else {
		fmt.Println("票已售空")
	}
}
