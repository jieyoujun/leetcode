package countandsay

import (
	"reflect"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	type testCase struct {
		n    int
		want string
	}
	testGroup := map[string]testCase{
		"case1": {1, "1"},
		"case2": {4, "1211"},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := CountAndSay(tC.n)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
