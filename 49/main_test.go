package groupanagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type testCase struct {
		strs []string
		want [][]string
	}
	testGroup := map[string]testCase{
		"case1": {
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
	}
	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := GroupAnagrams(tC.strs)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}
}
