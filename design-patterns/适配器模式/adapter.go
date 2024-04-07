package 适配器模式

import "fmt"

type IAdapter interface {
	Connect(name string) error
}

type IOS struct{}

func (c *IOS) LightningConnect(name string) error {
	fmt.Printf("%s connected", name)
	return nil
}

type Android struct{}

func (c *Android) TypeCConnect(name string) error {
	fmt.Printf("%s connected", name)
	return nil
}

type ComputerIosAdapter struct {
	Client IOS
}

func (c *ComputerIosAdapter) Connect(name string) error {
	c.Client.LightningConnect(name)
	return nil
}

type ComputerAndroidAdapter struct {
	Client Android
}

func (c *ComputerAndroidAdapter) Connect(name string) error {
	c.Client.TypeCConnect(name)
	return nil
}
