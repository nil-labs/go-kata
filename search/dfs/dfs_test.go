package dfs

import (
	"testing"
)

func TestNode_Search(t *testing.T) {
	tests := []struct {
		name string
		tree Node
		want int
	}{
		{
			name: "Success",
			tree: Node{
				Value: 1,
				Children: []Node{
					{
						Value: 2,
						Children: []Node{
							{
								Value: 6,
							},
						},
					},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.tree.Search(tt.want); got.Value != tt.want {
				t.Errorf("Node.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
