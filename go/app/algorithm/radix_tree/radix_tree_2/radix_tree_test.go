package radix_tree_2

import (
	"crypto/rand"
	"fmt"
	"testing"
)

/**
https://github.com/armon/go-radix/blob/master/radix_test.go
*/
func TestRadix2(test *testing.T) {
	dictionary := make(map[string]interface{})
	/*for i := 0; i < 1000; i++ {
		id := generateUUID()
		dictionary[id] = i
	}*/
	dictionary["h"] = 1
	dictionary["he"] = 2

	tree := NewFromMap2(dictionary)
	if tree.Len() != len(dictionary) {
		test.Fatalf("wrong length: want %d got %d", len(dictionary), tree.Len())
	}

	/*tree.Walk(func(key string, value interface{}) bool {
		println(key)
		return false
	})*/

	/*for key, value := range dictionary {
		out, ok := tree.Get(key)
		if !ok {
			test.Fatalf("missing key: %v", key)
		}

		if out != value {
			test.Fatalf("value mismatch, want %v got %v", value, out)
		}
	}*/

	/*outMin, _, _ := tree.Min()
	if outMin != min {
		test.Fatalf("bad min want %v got %v", min, outMin)
	}
	outMax, _, _ := tree.Max()
	if outMax != max {
		test.Fatalf("bad max want %v got %v", max, outMax)
	}*/
}
func NewFromMap2(dictionary map[string]interface{}) *Tree {
	tree := &Tree{
		root: &Node{},
	}
	for key, value := range dictionary {
		tree.Insert2(key, value)
	}
	
	return tree
}
func (tree *Tree) Insert2(key string, value interface{}) (interface{}, bool) {
	var parent *Node
	search := key
	rootNode := tree.root
	for  {
		if len(search) == 0 {
		
		}
		
		parent = rootNode
		rootNode = rootNode.getEdge(search[0])
		if rootNode == nil { // rootNode has no edge, create one edge
			parent.addEdge(Edge{
				label: search[0],
				node:  &Node{
					leaf:   &LeafNode{
						key:   search,
						value: value,
					},
					prefix: search,
					edges:  nil,
				},
			})
			tree.size++
			return nil, false
		}
		
		commonPrefix := longestPrefix(search, rootNode.prefix)
		if commonPrefix == len(rootNode.prefix) {
			search = search[commonPrefix:]
			continue
		}
	}
	
	return nil, false
}

func TestRoot(test *testing.T) {
	tree := New()
	_, ok := tree.Delete("")
	if ok {
		test.Fatalf("bad")
	}
	_, ok = tree.Insert("", true)
	if ok {
		test.Fatalf("bad")
	}
	value, ok := tree.Get("")
	if !ok || value != true {
		test.Fatalf("bad: %v", value)
	}
	value, ok = tree.Delete("")
	if !ok || value != true {
		test.Fatalf("bad: %v", value)
	}
}

func TestDelete(test *testing.T) {
	tree := New()
	search := []string{"", "a", "ab"}
	for _, value := range search {
		tree.Insert(value, true)
	}
	for _, value := range search {
		_, ok := tree.Delete(value)
		if !ok {
			test.Fatalf("bad %v", search)
		}
	}
}

func generateUUID() string {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		panic(fmt.Errorf("failed to read random bytes: %v", err))
	}

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:16])
}
