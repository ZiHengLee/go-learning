package 享元模式

var units = map[int]*WordUnit{
	1: {
		Id:   1,
		Data: "我",
		Type: "中国字",
	},
	2: {
		Id:   2,
		Data: "爱",
		Type: "中国字",
	},
	3: {
		Id:   3,
		Data: "中",
		Type: "中国字",
	},
	4: {
		Id:   4,
		Data: "国",
		Type: "中国字",
	},
}

// WordUnit 文字享元
type WordUnit struct {
	Id   uint
	Data string
	Type string
}

func Sentence(ids []int) string {
	result := ""
	for _, id := range ids {
		result += units[id].Data
	}
	return result
}
