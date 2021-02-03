package erchasousuoshudehouxubianlixulielcof

func VerifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	i := 0
	for ; i < len(postorder)-1; i++ {
		if postorder[i] > postorder[len(postorder)-1] {
			break
		}
	}
	for j := i; j < len(postorder)-1; j++ {
		if postorder[j] < postorder[len(postorder)-1] {
			return false
		}
	}
	left, right := true, true
	if i > 0 {
		left = VerifyPostorder(postorder[:i])
	}
	if i < len(postorder)-1 {
		right = VerifyPostorder(postorder[i : len(postorder)-1])
	}
	return left && right
}
