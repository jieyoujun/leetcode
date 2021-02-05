package yonglianggezhanshixianduilielcof

import "container/list"

type CQueue struct {
	in  *list.List
	out *list.List
}

func Constructor() CQueue {
	return CQueue{
		in:  list.New(),
		out: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.out.Len() == 0 {
		for this.in.Len() > 0 {
			this.out.PushBack(this.in.Remove(this.in.Back()))
		}
	}
	if this.out.Len() != 0 {
		return this.out.Remove(this.out.Back()).(int)
	}
	return -1
}
