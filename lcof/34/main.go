package erchashuzhongheweimouyizhidelujinglcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	curSum, path, paths := 0, []int{}, [][]int{}
	findPath(root, sum, curSum, &path, &paths)
	return paths
}

func findPath(root *TreeNode, sum, curSum int, path *[]int, paths *[][]int) {
	curSum += root.Val
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil && curSum == sum {
		// tmp := append([]int{}, (*path)...)
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*paths = append(*paths, tmp)
	}
	if root.Left != nil {
		findPath(root.Left, sum, curSum, path, paths)
	}
	if root.Right != nil {
		findPath(root.Right, sum, curSum, path, paths)
	}
	*path = (*path)[:len(*path)-1]
}
