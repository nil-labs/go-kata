package dfs

type Node struct {
	Value    int
	Children []Node
}

func (n Node) Search(v int) *Node {
	stack := []Node{n}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Value == v {
			return &current
		}
		stack = append(stack, current.Children...)
	}
	return nil
}
