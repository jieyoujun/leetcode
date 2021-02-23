package zhandeyarudanchuxulielcof

import (
	"reflect"
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	type testCase struct {
		pushed []int
		popped []int
		want   bool
	}
	testGroup := map[string]testCase{
		"case1": {
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 3, 2, 1},
			true,
		},
		"case2": {
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 5, 1, 2},
			false,
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := ValidateStackSequences(tC.pushed, tC.popped)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
