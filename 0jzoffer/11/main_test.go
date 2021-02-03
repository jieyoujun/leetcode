package xuanzhuanshuzudezuixiaoshuzilcof

import (
	"reflect"
	"testing"
)

func TestMinArray(t *testing.T) {
	type testCase struct {
		numbers []int
		want    int
	}
	testGroup := map[string]testCase{
		"case1": {[]int{3, 4, 5, 1, 2}, 1},
		"case2": {[]int{2, 2, 2, 0, 1}, 0},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := MinArray(tC.numbers)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
