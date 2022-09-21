package suffixtree

type Edge struct {
	Label []rune
	*Node
}

func newEdge(label []rune, node *Node) *Edge {
	return &Edge{Label: label, Node: node}
}
