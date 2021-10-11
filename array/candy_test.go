package array

import (
	"reflect"
	"testing"
)

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ok-0",
			args: args{
				ratings: []int{1, 0, 2},
			},
			want: []int{2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
