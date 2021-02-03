package longestpalindrome

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type testCase struct {
		s    string
		want string
	}
	testGroup := map[string]testCase{
		"case1": testCase{"babad", "aba"},
		"case2": testCase{"cbbd", "bb"},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := LongestPalindrome(tC.s)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
