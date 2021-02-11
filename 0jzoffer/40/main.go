package zuixiaodekgeshulcof

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]

	// 最小堆
	// h := &MinIntHeap{}
	// *h = append(*h, arr...)
	// heap.Init(h)
	// res := []int{}
	// for ; k > 0; k-- {
	// 	res = append(res, heap.Pop(h).(int))
	// }
	// return res
}

type MinIntHeap []int

func (mih MinIntHeap) Len() int           { return len(mih) }
func (mih MinIntHeap) Less(i, j int) bool { return mih[i] < mih[j] }
func (mih MinIntHeap) Swap(i, j int)      { mih[i], mih[j] = mih[j], mih[i] }

func (mih *MinIntHeap) Push(x interface{}) {
	*mih = append(*mih, x.(int))
}
func (mih *MinIntHeap) Pop() interface{} {
	x := (*mih)[len(*mih)-1]
	*mih = (*mih)[:len(*mih)-1]
	return x
}
