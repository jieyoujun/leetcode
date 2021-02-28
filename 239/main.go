package slidingwindowmaximum

import (
	"container/heap"
	"sort"
)

var a []int

type pq struct{ sort.IntSlice }

func (q pq) Less(i, j int) bool { return a[q.IntSlice[i]] > a[q.IntSlice[j]] }
func (q *pq) Push(x interface{}) {
	q.IntSlice = append(q.IntSlice, x.(int))
}
func (q *pq) Pop() interface{} {
	x := q.IntSlice[len(a)-1]
	q.IntSlice = q.IntSlice[:len(a)-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &pq{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	result := make([]int, 1, n-k+1)
	result[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		result = append(result, nums[q.IntSlice[0]])
	}
	return result

	// O(n*k)超时
	// var result []int
	// for i := 0; i+k <= len(nums); i++ {
	// 	result = append(result, max(nums[i:i+k]))
	// }
	// return result
}

func max(nums []int) int {
	m := nums[0]
	for _, num := range nums[1:] {
		if m < num {
			m = num
		}
	}
	return m
}
