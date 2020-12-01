package main

import "fmt"

type tree struct {
	value int
	left, right *tree
}
func main()  {
	values := []int{1, 3, 2, 5, -1, 18, 100, -8}
	Sort(values)
	//fmt.Println(Sort(values))
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(appendValues(values[:0], root))
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t // return &tree{value: value}
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}