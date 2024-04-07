package 迭代器模式

type Iterator interface {
	hasNext() bool
	getNext()
	CurrentItem() any
}

type Row struct {
	data string
}

type CollectionIterator struct {
	index int
	Rows  []*Row
}

func (u *CollectionIterator) hasNext() bool {
	if u.index < len(u.Rows) {
		return true
	}
	return false

}
func (u *CollectionIterator) getNext() {
	if u.hasNext() {
		u.index++
	}
	return
}

func (u *CollectionIterator) CurrentItem() any {
	return u.Rows[u.index]
}
