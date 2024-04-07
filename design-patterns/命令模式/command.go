package 命令模式

import "fmt"

// ICommand 命令
type ICommand interface {
	Execute() error
}

type TV struct{}

func (tv TV) Open() {
	fmt.Println("开机")
}

func (tv TV) Close() {
	fmt.Println("关机")
}

type OpenCommand struct {
	receiver TV
}

func (oc OpenCommand) Execute() {
	oc.receiver.Open()
}

type CloseCommand struct {
	receiver TV
}

func (cc CloseCommand) Execute() {
	cc.receiver.Close()
}
