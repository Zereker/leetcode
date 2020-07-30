package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var re strings.Builder
	traverse(root, &re)
	return re.String()
}

func traverse(root *TreeNode, re *strings.Builder) {
	if root == nil {
		re.WriteString("#,")
		return
	}

	traverse(root.Left, re)
	traverse(root.Right, re)

	v := strconv.Itoa(root.Val)
	re.WriteString(v)
	re.WriteString(",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(strings.TrimRight(data, ","), ",")

	return deTraverse(&nodes)
}

func deTraverse(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	n := (*nodes)[len(*nodes)-1]
	(*nodes) = (*nodes)[:len(*nodes)-1]

	if n == "#" {
		return nil
	}

	v, _ := strconv.Atoi(n)
	root := &TreeNode{
		Val: v,
	}

	root.Right = deTraverse(nodes)
	root.Left = deTraverse(nodes)
	return root
}