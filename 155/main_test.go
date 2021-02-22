package minstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinStack(t *testing.T) {
	// 输入：
	// ["MinStack","push","push","push","getMin","pop","top","getMin"]
	// [[],[-2],[0],[-3],[],[],[],[]]

	// 输出：
	// [null,null,null,null,-3,null,0,-2]
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	assert.Equal(t, -3, stack.GetMin())
	stack.Pop()
	stack.Pop()
	assert.Equal(t, -2, stack.GetMin())
}
