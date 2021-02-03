package powxn

import (
	"reflect"
	"testing"
)

func TestMyPow(t *testing.T) {
	type testCase struct {
		x    float64
		n    int
		want float64
	}
	testGroup := map[string]testCase{
		"case1": {
			x:    2.00000,
			n:    10,
			want: 1024.00000,
		},
		"case2": {
			x:    2.10000,
			n:    3,
			want: 9.26100,
		},
		"case3": {
			x:    2.00000,
			n:    -2,
			want: 0.25000,
		},
	}

	for caseName, tC := range testGroup {
		t.Run(caseName, func(t *testing.T) {
			got := MyPow(tC.x, tC.n)
			if !reflect.DeepEqual(tC.want, got) {
				t.Fatalf("want: %#v, got: %#v\n", tC.want, got)
			}
		})
	}

}
