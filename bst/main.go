package main

import (
	"errors"
	"fmt"
)

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
	bst.CreateTree([]int{9, 4, 6, 20, 170, 15, 1})
	bst.inorder(bst.root)
	_, err := bst.Lookup(100)
	if err == nil {
		fmt.Println("The node was found")
	} else {
		fmt.Printf("%v", err)
	}
}

func (bst *BST) CreateTree(a []int) {
	for _, num := range a {
		bst.insert(num)
	}
}

func (bst *BST) Lookup(val int) (*Node, error) {
	node := bst.root
	for node != nil {
		if node.value == val {
			return node, nil
		} else if val < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil, errors.New("not found")
}

func (bst *BST) insert(i int) (*Node, error) {
	newnode := Node{nil, nil, i}
	n := bst.root
	if n == nil {
		bst.root = &newnode
		return bst.root, nil
	}

	node := bst.root
	for {
		if i == node.value {
			return nil, errors.New("cannot add duplicate to bst")
		}
		if i < node.value {
			if node.left != nil {
				node = node.left
			} else {
				node.left = &newnode
				return bst.root, nil
			}
		} else if i > node.value {
			if node.right != nil {
				node = node.right
			} else {
				node.right = &newnode
				return bst.root, nil
			}
		}
	}
}

func (bst *BST) inorder(n *Node) {
	if n != nil {
		bst.inorder(n.left)
		fmt.Println(n.value)
		bst.inorder(n.right)
	}
}
