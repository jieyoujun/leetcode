package stringtointegeratoi

import (
	"reflect"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	type testCase struct {
		str  string
		want int
	}
	testGroup := map[string]testCase{
		"case1":  testCase{"42", 42},
		"case2":  testCase{"    -42    ", -42},
		"case3":  testCase{"4193 with words", 4193},
		"case4":  testCase{"words and 987", 0},
		"case5":  testCase{"    0000000000000   ", 0},
		"case6":  testCase{"0000000000000000123000", 123000},
		"case7":  testCase{"9223372036854775808", 2147483647},
		"case8":  testCase{"2147483646", 2147483646},
		"case9":  testCase{"2147483648", 2147483647},
		"case10": testCase{"-91283472332", -2147483648},
		"case11": testCase{"-2147483647", -2147483647},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := MyAtoi(tC.str)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
