package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result, nextLevel, left2Right := [][]int{}, []*TreeNode{root}, true
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
			if left2Right {
				lvlVals[i] = nextLevel[i].Val
			} else {
				lvlVals[nnxtlvl-1-i] = nextLevel[i].Val
			}
		}
		result = append(result, lvlVals)
		left2Right = !left2Right
		nextLevel = nextLevel[nnxtlvl:]
	}
	return result
}
