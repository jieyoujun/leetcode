package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	q := []*TreeNode{root}
	v := []string{}
	for len(q) > 0 {
		tnode := q[0]
		q = q[1:]
		if tnode != nil {
			v = append(v, strconv.Itoa(tnode.Val))
			q = append(q, tnode.Left, tnode.Right)
		} else {
			v = append(v, "#")
		}
	}
	return strings.Join(v, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	v := strings.Split(data, ",")
	if v[0] == "#" {
		return nil
	}
	i := 0
	vi, _ := strconv.Atoi(v[i])
	root := &TreeNode{Val: vi}
	q := []*TreeNode{root}
	for len(q) > 0 {
		tnode := q[0]
		q = q[1:]
		i++
		if v[i] != "#" {
			vi, _ = strconv.Atoi(v[i])
			tnode.Left = &TreeNode{Val: vi}
			q = append(q, tnode.Left)
		}
		i++
		if v[i] != "#" {
			vi, _ = strconv.Atoi(v[i])
			tnode.Right = &TreeNode{Val: vi}
			q = append(q, tnode.Right)
		}
	}
	return root
}
