package simplifypath

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "case2",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "case3",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
		{
			name: "case4",
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
		{
			name: "case5",
			args: args{
				path: "/a/../../b/../c//.//",
			},
			want: "/c",
		},
		{
			name: "case6",
			args: args{
				path: "/a//b////c/d//././/..",
			},
			want: "/a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
