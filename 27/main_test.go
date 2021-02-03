package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type testCase struct {
		nums []int
		val  int
		want int
	}
	testGroup := map[string]testCase{
		"case1": testCase{[]int{3, 2, 2, 3}, 3, 2},
		"case3": testCase{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := RemoveElement(tC.nums, tC.val)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
