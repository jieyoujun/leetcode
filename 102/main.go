package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result, nextLevel := [][]int{}, []*TreeNode{root}
	for len(nextLevel) > 0 {
		nnxtlvl := len(nextLevel)
		lvlVals := make([]int, nnxtlvl)
		for i := 0; i < nnxtlvl; i++ {
			if nextLevel[i].Left != nil {
				nextLevel = append(nextLevel, nextLevel[i].Left)
			}
			if nextLevel[i].Right != nil {
				nextLevel = append(nextLevel, nextLevel[i].Right)
			}
			lvlVals[i] = nextLevel[i].Val
		}
		result = append(result, lvlVals)
		nextLevel = nextLevel[nnxtlvl:]
	}
	return result
}
