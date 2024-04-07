package 责任链模式

import (
	"fmt"
)

type Filter interface {
	Filter(content string)
}

type FilterChain struct {
	filters []Filter
	index   int
}

func (c *FilterChain) AddFilter(filter Filter) {
	c.filters = append(c.filters, filter)
}

func (c *FilterChain) Filer(content string) {
	for c.Next(content) {
	}
}

func (c *FilterChain) Next(content string) bool {
	if c.index == len(c.filters) {
		return false
	}
	c.filters[c.index].Filter(content)
	c.index++
	return true
}

type AdFilter struct {
}

func (ad *AdFilter) Filter(content string) {
	fmt.Println("广告过滤")
}

type YellowFilter struct {
}

func (yellow *YellowFilter) Filter(content string) {
	fmt.Println("色情过滤")
}

// SensitiveFilter 敏感词过滤
type SensitiveFilter struct {
}

func (sensitive *SensitiveFilter) Filter(content string) {
	fmt.Println("敏感词过滤")
}
