package tihuankonggelcof

import (
	"reflect"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	type testCase struct {
		s    string
		want string
	}
	testGroup := map[string]testCase{
		"case1": {"We are happy.", "We%20are%20happy."},
		"case2": {"   We are happy.", "%20%20%20We%20are%20happy."},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := ReplaceSpace(tC.s)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
