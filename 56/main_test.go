package mergeintervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// {
		// 	name: "case1",
		// 	args: args{
		// 		intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		// 	},
		// 	want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		// },
		// {
		// 	name: "case2",
		// 	args: args{
		// 		intervals: [][]int{{1, 4}, {4, 5}},
		// 	},
		// 	want: [][]int{{1, 5}},
		// },
		// {
		// 	name: "case3",
		// 	args: args{
		// 		intervals: [][]int{{1, 4}, {5, 6}},
		// 	},
		// 	want: [][]int{{1, 4}, {5, 6}},
		// },
		// {
		// 	name: "case4",
		// 	args: args{
		// 		intervals: [][]int{{1, 4}, {0, 0}},
		// 	},
		// 	want: [][]int{{0, 0}, {1, 4}},
		// },
		// {
		// 	name: "case5",
		// 	args: args{
		// 		intervals: [][]int{{1, 4}, {0, 4}},
		// 	},
		// 	want: [][]int{{0, 4}},
		// },
		// {
		// 	name: "case6",
		// 	args: args{
		// 		intervals: [][]int{{1, 4}, {1, 4}},
		// 	},
		// 	want: [][]int{{1, 4}},
		// },
		{
			name: "case7",
			args: args{
				intervals: [][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}},
			},
			want: [][]int{{2, 4}, {5, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
