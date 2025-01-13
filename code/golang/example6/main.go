package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func stringfyTree(root *Node) string {
	if root == nil {
		return ""
	}

	result := ""
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			result += fmt.Sprintf("%d ", node.Data)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result += "\n"
	}

	return result
}
