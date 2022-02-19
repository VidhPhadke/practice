package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

type BST struct {
	root *Node
}

func main() {
	var bst *BST = &BST{nil}
	bst.CreateTree([]int{10})
	bst.inorder(bst.root)

}

func (bst *BST) CreateTree(a []int) {
	for _, num := range a {
		bst.insert(num)
	}
}

func (bst *BST) insert(i int) {
	n := bst.root
	if n == nil {
		bst.root = &Node{nil, nil, i}
	}

}

func (bst *BST) inorder(n *Node) {
	if n != nil {
		bst.inorder(n.left)
		fmt.Println(n.value)
		bst.inorder(n.right)
	}
}
