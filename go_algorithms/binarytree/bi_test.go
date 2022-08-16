package binarytree

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestBisearch(t *testing.T) {
	type PushReviewMsgParam struct {
		Apps []string `json:"apps,omitempty"`
		Rid  int64    `json:"rid"`
		Pid  int64    `json:"pid"`
		Mid  int64    `json:"mid"`
		Dmid int64    `json:"dmid"`
		Type string   `json:"type"`
	}
	type NewPushParam struct {
		Apps    []string    `json:"apps"`
		ToMid   int64       `json:"to_mid"`
		Type    int         `json:"type"`
		Ct      int64       `json:"ct"`
		Subtype int         `json:"subtype"`
		FromMid int64       `json:"from_mid,omitempty"`
		Src     interface{} `json:"src,omitempty"`
		New     interface{} `json:"new,omitempty"`
		//Counters []NewPushCounterSt `json:"counters,omitempty"`
		Aps interface{} `json:"aps,omitempty"`
	}
	aa := NewPushParam{
		Type:    123,
		Subtype: 5,
		Apps:    []string{"1", "2", "3"},
	}
	bb, _ := json.Marshal(&aa)
	//print(bb)
	cc := &PushReviewMsgParam{}
	err := json.Unmarshal(bb, cc)
	print(err.Error())
	print(cc.Apps[0])
}

func TestSimple(t *testing.T) {
	type TData struct {
		List    []int `json:"list,omitempty"`
		Version int   `json:"version"`
	}
	var newData TData

	result, _ := json.Marshal(newData)
	fmt.Printf("%+v", result)
}
func DecayDays(ct int64) (decayRate float64) {
	//e^(-0.2*limit)
	limit := (time.Now().Unix() - ct) / 86400
	if limit < 0 {
		limit = 0
	} else if limit > 20 {
		limit = 20
	}
	print(limit)
	decayRate = math.Pow(math.E, -float64(limit)*0.2)
	return
}

func TestReBuildTree(t *testing.T) {
	pres,mids:=[]int{1,2,3},[]int{2,1,3}
	head := buildTree(pres,mids)
	//fmt.Printf("%+v", head.Right.Data)
	//PreOrder(head)
	ret := zigzagLevelOrder(head)
	fmt.Println(ret)
}

func TestMaxPathSum(t *testing.T) {
	a := TreeNode{Val: -2}
	b := TreeNode{Val: 1}
	a.Left = &b
	fmt.Println(maxPathSum(&a))
}

func TestIsCompleteTree(t *testing.T) {
	a := TreeNode{Val: -2}
	b := TreeNode{Val: 1}
	a.Left = &b
	fmt.Println(kthLargest(&a,1))
}