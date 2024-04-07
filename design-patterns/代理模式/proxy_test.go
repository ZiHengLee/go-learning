package 代理模式

import "testing"

func TestProxy(t *testing.T) {

	station := &Station{3}
	proxy := &StationProxy{station}
	station.sell("A")

	proxy.sell("B")
	proxy.sell("C")
	proxy.sell("D")
}
