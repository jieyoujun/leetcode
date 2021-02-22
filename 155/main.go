package minstack

type MinStack struct {
	nums []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
	if len(this.mins) == 0 || this.mins[len(this.mins)-1] > x {
		this.mins = append(this.mins, x)
	} else {
		this.mins = append(this.mins, this.mins[len(this.mins)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.nums) > 0 {
		this.nums = this.nums[:len(this.nums)-1]
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.nums) > 0 {
		return this.mins[len(this.mins)-1]
	}
	return -1
}
