package 命令模式

import "testing"

func TestCommand(t *testing.T) {
	tv := TV{}
	openCommand := OpenCommand{tv}
	closeCommand := CloseCommand{tv}
	openCommand.Execute()
	closeCommand.Execute()
}
