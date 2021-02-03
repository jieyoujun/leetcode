package zhandeyarudanchuxulielcof

func ValidateStackSequences(pushed []int, popped []int) bool {
	stack, i := []int{}, 0
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1] // pop
			i++
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
