package shuzuzhongdenixuduilcof

import (
	"reflect"
	"testing"
)

func TestReversePairs(t *testing.T) {
	type testCase struct {
		nums []int
		want int
	}
	testGroup := map[string]testCase{
		"case1": {[]int{7, 5, 6, 4}, 5},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := ReversePairs(tC.nums)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
