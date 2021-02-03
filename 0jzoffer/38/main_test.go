package zifuchuandepailielcof

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	type testCase struct {
		s    string
		want []string
	}
	testGroup := map[string]testCase{
		"case1": {"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := Permutation(tC.s)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
