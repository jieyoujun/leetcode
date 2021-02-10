package pathsumii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	dfs(root, targetSum, nil, &result)
	return result
}

func dfs(root *TreeNode, targetsum int, path []int, result *[][]int) {
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val == targetsum {
			*result = append(*result, append([]int{}, path...))
			return
		}
		return
	}
	if root.Left != nil {
		dfs(root.Left, targetsum-root.Val, path, result)
	}
	if root.Right != nil {
		dfs(root.Right, targetsum-root.Val, path, result)
	}
}
