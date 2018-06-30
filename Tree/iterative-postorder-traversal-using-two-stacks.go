package main

import "fmt"

const SIZE = 100

type node struct {
	data int
	left *node
	right *node
}

type btree struct {
	root *node
}

type stack struct {
	items [SIZE]*node
	top int
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 
func (stk *stack) push(node *node) {
	stk.top = stk.top + 1
	stk.items[stk.top] = node
}

func (stk *stack) pop() *node {
	node := stk.items[stk.top]
	stk.top = stk.top - 1
	return node
}

func (bt *btree) postOrderIterative(node *node) {
	s1 := stack{top: -1}
	s2 := stack{top: -1}
	s1.push(node)

	for s1.top >= 0 {
		node = s1.items[s1.top]
		s2.push(s1.pop())
		if node.left != nil {
			s1.push(node.left)
		}
		if node.right != nil {
			s1.push(node.right)
		}
	}

	for i := s2.top; i >= 0; i-- {
		fmt.Print(s2.items[i].data, " ")
	}
}

func main() {
	bt := btree{root: &node{data: 1}}

	bt.root.left        = &node{data: 2}
	bt.root.right       = &node{data: 3} 
	bt.root.left.left   = &node{data: 4}
	bt.root.left.right  = &node{data: 5}
	bt.root.right.left  = &node{data: 6}
	bt.root.right.right = &node{data: 7}

	bt.postOrderIterative(bt.root)
	fmt.Println()
}