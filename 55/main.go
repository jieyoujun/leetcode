package jumpgame

func canJump(nums []int) bool {
	scale := 0
	for i := 0; i < len(nums); i++ {
		if i > scale {
			return false
		}
		if scale < i+nums[i] {
			scale = i + nums[i]
		}
		if scale >= len(nums)-1 {
			return true
		}
	}
	return false
}
