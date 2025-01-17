package suffixtree

import (
	"sort"
)

type Node struct {
	/*
	 * The payload array used to store the data (indexes) associated with this node.
	 * In this case, it is used to store all property indexes.
	 */
	Data []int
	/**
	 * The set of edges starting from this node
	 */
	Edges []*Edge
	/**
	 * The suffix link as described in Ukkonen's paper.
	 * if str is the string denoted by the path from the root to this, this.suffix
	 * is the node denoted by the path that corresponds to str without the first rune.
	 */
	Suffix *Node
}

/*
 * getData returns the first numElements elements from the ones associated to this node.
 *
 * Gets data from the payload of both this node and its children, the string representation
 * of the path to this node is a substring of the one of the children nodes.
 *
 * @param numElements the number of results to return. Use <=0 to get all
 * @return the first numElements associated to this node and children
 */
func (n *Node) getData(numElements int) (ret []int) {

	if numElements > 0 {
		if numElements > len(n.Data) {
			numElements -= len(n.Data)
			ret = n.Data
		} else {
			ret = n.Data[:numElements]
			return
		}
	} else {
		ret = n.Data
	}

	// need to get more matches from child nodes. This is what may waste time
	for _, edge := range n.Edges {
		Data := edge.Node.getData(numElements)
	NEXTIDX:
		for _, idx := range Data {
			for _, v := range ret {
				if v == idx {
					continue NEXTIDX
				}
			}

			if numElements > 0 {
				numElements--
			}
			ret = append(ret, idx)
		}

		if numElements == 0 {
			break
		}
	}

	return
}

// addRef adds the given index to the set of indexes associated with this
func (n *Node) addRef(index int) {
	if n.contains(index) {
		return
	}
	n.addIndex(index)
	// add this reference to all the suffixes as well
	iter := n.Suffix
	for iter != nil {
		if iter.contains(index) {
			break
		}
		iter.addRef(index)
		iter = iter.Suffix
	}
}

func (n *Node) contains(index int) bool {
	i := sort.SearchInts(n.Data, index)
	return i < len(n.Data) && n.Data[i] == index
}

func (n *Node) addEdge(r rune, e *Edge) {
	if idx := n.search(r); idx == -1 {
		n.Edges = append(n.Edges, e)
		sort.Slice(n.Edges, func(i, j int) bool { return n.Edges[i].Label[0] < n.Edges[j].Label[0] })
	} else {
		n.Edges[idx] = e
	}

}

func (n *Node) getEdge(r rune) *Edge {
	idx := n.search(r)
	if idx < 0 {
		return nil
	}
	return n.Edges[idx]
}

func (n *Node) search(r rune) int {
	idx := sort.Search(len(n.Edges), func(i int) bool { return n.Edges[i].Label[0] >= r })
	if idx < len(n.Edges) && n.Edges[idx].Label[0] == r {
		return idx
	}

	return -1
}

func (n *Node) addIndex(idx int) {
	n.Data = append(n.Data, idx)
}

func newNode() *Node {
	return &Node{}
}
