package array

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		children []int
		cookies  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok-0",
			args: args{
				children: []int{1, 2},
				cookies:  []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "fail-0",
			args: args{
				children: []int{4, 5, 6},
				cookies:  []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "ok-1",
			args: args{
				children: []int{2, 1},
				cookies:  []int{3, 2, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.args.children, tt.args.cookies); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
