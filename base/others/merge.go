package others

import (
	"sort"
)

type X struct {
	Left int
	Right int

}

type SortX []X

func (s SortX) Len() int {
	return len(s)

}

func (s SortX) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortX) Less(i, j int) bool {
	return s[i].Left < s[j].Left

}

func merge(intervals [][]int) [][]int {
	ret := make([][]int,0)
	helpList := make(SortX,0)
	for _,v:= range intervals{
		t:=X{Left: v[0],Right: v[1]}
		helpList = append(helpList, t)
	}
	sort.Sort(helpList)
	i := 0
	for i< len(helpList){
		j := i+1
		newX := make([]int,2)
		newX[0] = helpList[i].Left
		newX[1] = helpList[i].Right
		for j < len(helpList) && helpList[j].Left <=newX[1]{
			if helpList[j].Right > newX[1]{
				newX[1] = helpList[j].Right
			}
			j++
		}
		ret = append(ret, newX)
		i = j
	}
	return ret
}
