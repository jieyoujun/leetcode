package containerwithmostwater

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	type testCase struct {
		height []int
		want   int
	}
	testGroup := map[string]testCase{
		"case1": testCase{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := MaxArea(tC.height)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
