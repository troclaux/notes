package main

import (
	"testing"
)

func createSampleTree() *Node {
	root := &Node{Data: 1}
	root.Left = &Node{Data: 2}
	root.Right = &Node{Data: 3}
	root.Left.Left = &Node{Data: 4}
	root.Left.Right = &Node{Data: 5}
	root.Right.Left = &Node{Data: 6}
	root.Right.Right = &Node{Data: 7}
	return root
}

func TestPrintTree(t *testing.T) {
	testTree := createSampleTree()
	result := stringfyTree(testTree)
	expected := `1 
2 3 
4 5 6 7 
`

	if result != expected {
		t.Errorf("Expected (len=%d):\n%q\nGot (len=%d):\n%q", len(expected), expected, len(result), result)
	}

}
