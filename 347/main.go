package topkfrequentelements

import "container/heap"

type pq [][2]int // num, num_count

func (q pq) Len() int           { return len(q) }
func (q pq) Less(i, j int) bool { return q[i][1] > q[j][1] }
func (q pq) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *pq) Push(x interface{}) {
	*q = append(*q, x.([2]int))
}
func (q *pq) Pop() interface{} {
	x := (*q)[len(*q)-1]
	(*q) = (*q)[:len(*q)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	q := &pq{}
	heap.Init(q)
	for num, count := range m {
		heap.Push(q, [2]int{num, count})
	}
	result := make([]int, 0, k)
	for k > 0 {
		result = append(result, heap.Pop(q).([2]int)[0])
		k--
	}
	return result
}
