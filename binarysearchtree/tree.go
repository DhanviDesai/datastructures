package binarysearchtree

import "fmt"

var root *Node

func AddNode(key int) {
	newNode := createNode(key)
	root = insertNode(root, &newNode)
}

func InorderTraversal() []int {
	result := make([]int, countNodes(root))
	index := 0
	return inorderTraversal(root, result, &index)
}

func CountNodes() int {
	return countNodes(root)
}

func countNodes(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.leftChild) + countNodes(root.rightChild)
}

func PreoderTraversal() {
	preorderTraversal(root)
	fmt.Println()
}

func preorderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d\t", root.key)
	preorderTraversal(root.leftChild)
	preorderTraversal(root.rightChild)
}

func inorderTraversal(root *Node, result []int, index *int) []int {
	if root == nil {
		return result
	}
	inorderTraversal(root.leftChild, result, index)
	result[*index] = root.key
	*index++
	inorderTraversal(root.rightChild, result, index)
	return result
}

func insertNode(root *Node, newNode *Node) *Node {
	if root == nil {
		root = newNode
		return root
	} else if newNode.key < root.key {
		root.leftChild = insertNode(root.leftChild, newNode)
	} else {
		root.rightChild = insertNode(root.rightChild, newNode)
	}
	return root
}

func createNode(key int) Node {
	newNode := Node{
		key:        key,
		leftChild:  nil,
		rightChild: nil,
	}
	return newNode
}
