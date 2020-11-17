package main

import (
	"fmt"
)
type tree struct {
	value int
	left, right *tree
}

func main() {
	Mysort([]int{5, 3, 1})

}

func Mysort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)

	fmt.Println(*root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

// 将每一个 slice的值,压入 tree的左边和右边,形成一个二叉树
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}