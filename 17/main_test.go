package lettercombinationsofaphonenumber

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type testCase struct {
		digits string
		want   []string
	}
	testGroup := map[string]testCase{
		"case1": {"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := LetterCombinations(tC.digits)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
