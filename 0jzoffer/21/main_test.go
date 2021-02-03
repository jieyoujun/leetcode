package diaozhengshuzushunxushiqishuweiyuoushuqianmianlcof

import (
	"reflect"
	"testing"
)

func TestExchange(t *testing.T) {
	type testCase struct {
		nums []int
		want []int
	}
	testGroup := map[string]testCase{
		"case1": {[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
		"case2": {[]int{1, 3, 5}, []int{1, 3, 5}},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := Exchange(tC.nums)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
