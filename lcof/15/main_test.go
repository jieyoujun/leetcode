package erjinzhizhong1degeshulcof

import (
	"reflect"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	type testCase struct {
		num  uint32
		want int
	}
	testGroup := map[string]testCase{
		"case1": {11, 3},
		"case2": {128, 1},
		"case3": {4294967293, 31},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := HammingWeight(tC.num)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
