package serializeanddeserializebinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SerializeAndDeserializeBinaryTree(t *testing.T) {
	codec := Constructor()
	var tree *TreeNode
	tree = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	assert.Equal(t, "1,2,3,#,#,4,5,#,#,#,#", codec.serialize(tree))
	assert.Equal(t, tree, codec.deserialize("1,2,3,#,#,4,5,#,#,#,#"))
}
