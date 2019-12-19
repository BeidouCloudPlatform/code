package radix_tree_2

import (
	"sort"
	"strings"
)

type Edges []Edge

// Edge node
type Edge struct {
	label byte
	node  *Node
}

func (edges Edges) Len() int {
	return len(edges)
}

func (edges Edges) Less(i, j int) bool {
	return edges[i].label < edges[j].label
}

func (edges Edges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}

func (edges Edges) Sort() {
	sort.Sort(edges)
}

type Node struct {
	// store leaf
	leaf   *LeafNode
	prefix string
	edges  Edges
}

type LeafNode struct {
	key   string
	value interface{}
}

func (node *Node) isLeaf() bool {
	return node.leaf != nil
}

func (node *Node) getEdge(label byte) *Node {
	number := len(node.edges)
	idx := sort.Search(number, func(i int) bool {
		return node.edges[i].label >= label
	})
	
	if idx < number && node.edges[idx].label == label {
		return node.edges[idx].node
	}
	
	return nil
}

func (node *Node) addEdge(edge Edge) {
	node.edges = append(node.edges, edge)
	node.edges.Sort()
}

func (node *Node) updateEdge(label byte, newNode *Node) {
	number := len(node.edges)
	idx := sort.Search(number, func(i int) bool {
		return node.edges[i].label >= label
	})
	
	if idx < number && node.edges[idx].label == label {
		node.edges[idx].node = newNode
		return
	}
	
	panic("missing edge")
}


type Tree struct {
	root *Node
	size int
}

type WalkFn func(key string, value interface{}) bool

func (tree *Tree) Insert(key string, value interface{}) (interface{}, bool) {
	var parent *Node
	search := key
	rootNode := tree.root

	for {
		if len(search) == 0 { // substring key util key is ""
			if rootNode.isLeaf() {
				old := rootNode.leaf.value
				rootNode.leaf.value = value
				return old, true
			}

			rootNode.leaf = &LeafNode{
				key:   search,
				value: value,
			}
			tree.size++
			return nil, false
		}

		parent = rootNode
		rootNode = rootNode.getEdge(search[0])
		if rootNode == nil { // No edge, create one
			edge := Edge{
				label: search[0],
				node: &Node{
					leaf: &LeafNode{
						key:   search,
						value: value,
					},
					prefix: search,
				},
			}
			parent.addEdge(edge)
			tree.size++
			return nil, false
		}

		commonPrefix := longestPrefix(search, rootNode.prefix)
		if commonPrefix == len(rootNode.prefix) {
			search = search[commonPrefix:]
			continue
		}

		// Split the node
		tree.size++
		child := &Node{
			prefix: search[:commonPrefix],
		}
		parent.updateEdge(search[0], child)

		// Restore the existing node
		child.addEdge(Edge{
			label: rootNode.prefix[commonPrefix],
			node:  rootNode,
		})
		rootNode.prefix = rootNode.prefix[commonPrefix:]

		// if the new key is a subnet, add it into this node
		leaf := &LeafNode{
			key:   search,
			value: value,
		}
		search = search[commonPrefix:]
		if len(search) == 0 {
			child.leaf = leaf
			return nil, false
		}
		child.addEdge(Edge{
			label: search[0],
			node: &Node{
				leaf:   leaf,
				prefix: search,
			},
		})
		return nil, false
	}
}

func (tree *Tree) Get(key string) (interface{}, bool) {
	node := tree.root
	search := key
	for {
		if len(search) == 0 {
			if node.isLeaf() {
				return node.leaf.value, true
			}
		}

		node = node.getEdge(search[0])
		if node == nil {
			break
		}
		if strings.HasPrefix(search, node.prefix) {
			search = search[len(node.prefix):]
		} else {
			break
		}
	}

	return nil, false
}

func (tree *Tree) Delete(key string) (interface{}, bool) {
	search := key
	rootNode := tree.root
	var parent *Node
	for {
		if len(search) == 0 {
			if !rootNode.isLeaf() {
				break
			}
			goto DELETE
		}

		parent = rootNode
		rootNode = rootNode.getEdge(search[0])
		if rootNode == nil {
			break
		}
		if strings.HasPrefix(search, rootNode.prefix) {
			search = search[len(rootNode.prefix):]
		} else {
			break
		}
	}

	return nil, false

DELETE:
	leaf := rootNode.leaf
	rootNode.leaf = nil
	tree.size--
	if parent != nil && len(rootNode.edges) == 0 {

	}

	// Check if delete this node from the parent

	// Check if merge this node

	// Check if merge the parent's other node
	if parent != nil {

	}

	return leaf.value, true
}

func (tree *Tree) Len() int {
	return tree.size
}

func (tree *Tree) Walk(fn WalkFn) {
	recursiveWalk(tree.root, fn)
}


func longestPrefix(k1 string, k2 string) int {
	max := len(k1)
	if l := len(k2); l < max {
		max = l
	}
	
	var i int
	for i = 0; i < max; i++ {
		if k1[i] != k2[i] {
			break
		}
	}
	
	return i
}

/*func (tree *Tree) Min() (string, interface{}, bool)  {

}

func (tree *Tree) Max() (string, interface{}, bool)  {

}*/

func New() *Tree {
	return NewFromMap(nil)
}

func NewFromMap(dictionary map[string]interface{}) *Tree {
	tree := &Tree{
		root: &Node{},
	}
	for key, value := range dictionary {
		tree.Insert(key, value)
	}
	
	return tree
}

func recursiveWalk(node *Node, fn WalkFn) bool {
	if node.leaf != nil && fn(node.leaf.key, node.leaf.value) {
		return true
	}
	
	for _, edge := range node.edges {
		if recursiveWalk(edge.node, fn) {
			return true
		}
	}
	
	return false
}
