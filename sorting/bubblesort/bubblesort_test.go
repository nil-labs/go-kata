package bubblesort

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name     string
		unsorted []int
		sorted   []int
	}{
		{
			name:     "valid",
			unsorted: []int{4, 5, 2, 1, 0},
			sorted:   []int{0, 1, 2, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort(tt.unsorted)
			if !reflect.DeepEqual(tt.unsorted, tt.sorted) {
				t.Errorf("expecting array %v to have order %v", tt.unsorted, tt.sorted)
			}
		})
	}
}
