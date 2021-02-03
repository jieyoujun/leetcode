package multiplystrings

import (
	"reflect"
	"testing"
)

func TestMultiply(t *testing.T) {
	type testCase struct {
		num1 string
		num2 string
		want string
	}
	testGroup := map[string]testCase{
		"case1": {"2", "3", "6"},
		"case2": {"123", "456", "56088"},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := Multiply(tC.num1, tC.num2)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
