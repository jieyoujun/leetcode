package palindromenumber

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type testCase struct {
		x  int
		want bool
	}
	testGroup := map[string]testCase{
		"case1":  testCase{1221, true},
		"case2":  testCase{121, true},
		"case3":  testCase{-121, false},
		"case4":  testCase{10, false},
		"case5":  testCase{0, true},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := IsPalindrome(tC.x)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
