package zigzagconversion

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	type testCase struct {
		s       string
		numRows int
		want    string
	}
	testGroup := map[string]testCase{
		"case1": {"LEETCODEISHIRING", 3, "LCIRETOESIIGEDHN"},
		"case2": {"LEETCODEISHIRING", 4, "LDREOEIIECIHNTSG"},
		"case3": {"AB", 1, "AB"},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := Convert(tC.s, tC.numRows)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want:%#v, got:%#v\n", tC.want, got)
			}
		})
	}
}
