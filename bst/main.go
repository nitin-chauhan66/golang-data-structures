package main

import (
	"fmt"
	"strings"
)

type node struct {
	key   uint32
	left  *node
	right *node
}

type bst struct {
	root *node
	len  uint32
}

func (b bst) String() string {
	sb := &strings.Builder{}
	inOrderTraversal(sb, b.root)
	return sb.String()
}

func inOrderTraversal(sb *strings.Builder, n *node) {
	if n == nil {
		return
	}
	inOrderTraversal(sb, n.left)
	sb.WriteString(fmt.Sprintf("%d ", n.key))
	inOrderTraversal(sb, n.right)
}

func (b *bst) add(key uint32) {
	b.root = b.addByNode(key, b.root)
}

func (b *bst) addByNode(key uint32, root *node) *node {
	if root == nil {
		return &node{key: key}
	}

	if key < root.key {
		root.left = b.addByNode(key, root.left)
	} else if key > root.key {
		root.right = b.addByNode(key, root.right)
	}

	return root
}

func (b *bst) remove(key uint32) {
	b.removeByNode(key, b.root)
}

func (b *bst) removeByNode(key uint32, root *node) *node {
	if root == nil {
		return root
	}

	if key > root.key {
		root.right = b.removeByNode(key, root.right)
	} else if key < root.key {
		root.left = b.removeByNode(key, root.left)
	} else {
		if root.left == nil {
			return root.right
		} else {

			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}

			root.key = temp.key

			root.left = b.removeByNode(temp.key, root.left)
		}
	}
	return root
}

func (b bst) search(key uint32) (*node, bool) {
	return b.searchByNode(key, b.root)
}

func (b bst) searchByNode(key uint32, root *node) (*node, bool) {
	if root == nil {
		return nil, false
	}

	if key < root.key {
		return b.searchByNode(key, root.left)
	} else if key > root.key {
		return b.searchByNode(key, root.right)
	}

	return root, true
}

func main() {
	tree := &bst{}
	tree.add(12)
	tree.add(10)
	tree.add(11)
	tree.add(13)
	fmt.Println(tree)
	fmt.Println(tree.search(11))
	fmt.Println(tree.search(13))
	fmt.Println(tree.search(1))
	fmt.Println(tree.search(3))
	tree.remove(12)
	tree.remove(11)
	fmt.Println(tree)
	fmt.Println(tree.search(11))
}
