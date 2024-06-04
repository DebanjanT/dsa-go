package main

import "fmt"

type bstnode struct {
	data  int
	left  *bstnode
	right *bstnode
}

var SearchIterCount int

func (n *bstnode) Insert(v int) {
	if v > n.data {
		// Right Side of Node
		// If node empty
		if n.right == nil {
			n.right = &bstnode{
				data: v,
			}
		} else {
			// If node not empty
			// Recursion call
			n.right.Insert(v)
		}
	} else {
		if n.left == nil {
			n.left = &bstnode{
				data: v,
			}
		} else {
			// Recursion call
			n.left.Insert(v)
		}
	}
}

func (n *bstnode) Search(v int) {
	fmt.Println("START")
	SearchIterCount++
	if n.data == v {
		fmt.Println("e0")
		fmt.Println("Found the value in the tree")
		return
	} else {
		fmt.Println("e1")
		if v < n.data && n.left != nil {
			fmt.Println("e2")
			n.left.Search(v)
		} else if v > n.data && n.right != nil {
			fmt.Println("e3")
			n.right.Search(v)
		} else {

			fmt.Println("Value Not Found ")
		}
	}
}

func GoBinarySearchTree() {
	tree := &bstnode{data: 50}
	tree.Insert(20)
	tree.Insert(200)
	tree.Insert(44)
	tree.Insert(901)
	tree.Insert(88)
	tree.Insert(10)
	tree.Insert(49)
	tree.Insert(589)
	tree.Insert(201)
	tree.Insert(120)
	tree.Insert(9)
	tree.Insert(23)
    tree.Search(201)
	fmt.Println("Search Count: ", SearchIterCount)
}
