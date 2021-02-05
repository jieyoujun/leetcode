package removeduplicatesfromsortedarrayii

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{

		{
			name: "case2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want:  7,
			want1: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("removeDuplicates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
