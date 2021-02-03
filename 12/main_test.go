package integertoroman

import (
	"reflect"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	type testCase struct {
		num  int
		want string
	}
	testGroup := map[string]testCase{
		"case1": testCase{3, "III"},
		"case2": testCase{4, "IV"},
		"case3": testCase{9, "IX"},
		"case4": testCase{58, "LVIII"},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := IntToRoman(tC.num)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
