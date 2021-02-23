package bazifuchuanzhuanhuanchengzhengshulcof

import "testing"

func Test_strToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{str: "20000000000000000000"},
			want: 2147483647,
		},
		{
			name: "case1",
			args: args{str: "42"},
			want: 42,
		},
		{
			name: "case2",
			args: args{str: "  -42"},
			want: -42,
		},
		{
			name: "case3",
			args: args{str: "4193 with words"},
			want: 4193,
		},
		{
			name: "case4",
			args: args{str: "words and 987"},
			want: 0,
		},
		{
			name: "case5",
			args: args{str: "-91283472332"},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToInt(tt.args.str); got != tt.want {
				t.Errorf("strToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
