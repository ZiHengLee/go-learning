package 适配器模式

import "testing"

func TestAdapter(t *testing.T) {
	var a IAdapter = &ComputerIosAdapter{
		Client: IOS{},
	}
	a.Connect("我的iphone4s")

	var b IAdapter = &ComputerAndroidAdapter{
		Client: Android{},
	}
	b.Connect("我的小米2s")
}
