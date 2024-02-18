package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

var mockTree = &Node{
	Value: 5,
	Left: &Node{
		Value: 3,
		Left: &Node{
			Value: 1,
			Left:  nil,
			Right: nil,
		},
		Right: &Node{
			Value: 4,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &Node{
		Value: 7,
		Left: &Node{
			Value: 6,
			Left:  nil,
			Right: &Node{
				Value: 3,
				Left: &Node{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 4,
					Left:  nil,
					Right: &Node{
						Value: 1,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	},
}

func main() {
	fmt.Println(getBinaryTreeHeight(*mockTree))
}

func getBinaryTreeHeight(node Node) (height int) {
	if node.Right == nil && node.Left == nil {
		return 0
	}

	rightHeight := getHeight(*node.Right, height+1)
	leftHeight := getHeight(*node.Left, height+1)
	return greater(rightHeight, leftHeight)
}

func getHeight(node Node, height int) int {
	if rightNode := node.Right; rightNode != nil {
		return getHeight(*rightNode, height+1)
	}
	if leftNode := node.Left; leftNode != nil {
		return getHeight(*leftNode, height+1)
	}

	return height
}

func greater(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
