package quicksort

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		arr    []int
		low    int
		high   int
		expect []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Valid",
			args: args{
				arr:    []int{5, 4, 3, 2, 1},
				expect: []int{1, 2, 3, 4, 5},
				low:    0,
				high:   4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort(tt.args.arr, tt.args.low, tt.args.high)
			if !reflect.DeepEqual(tt.args.arr, tt.args.expect) {
				t.Error(fmt.Sprintf("expecting array %v to have order %v", tt.args.arr, tt.args.expect))
			}
		})
	}
}
