package 责任链模式

import (
	"testing"
)

func TestFilter(t *testing.T) {
	chain := &FilterChain{}
	chain.AddFilter(&AdFilter{})
	chain.AddFilter(&SensitiveFilter{})
	chain.AddFilter(&YellowFilter{})
	chain.Filer("hello world")
}
