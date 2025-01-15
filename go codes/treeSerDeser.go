package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
}

func Constructor1() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	s := ""
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode == nil {
			s = s + "#,"
		} else {
			s += fmt.Sprintf("%d,", curNode.Val) // Convert int to string using fmt.Sprintf
			queue = append(queue, curNode.Left)
			queue = append(queue, curNode.Right)
		}
	}
	return s
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	nodes := strings.Split(data, ",")
	rootVal := parseInt(nodes[0])
	root := &TreeNode{Val: rootVal}

	queue := []*TreeNode{root}

	i := 1

	for i < len(nodes) && len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if nodes[i] != "#" {
			leftVal := parseInt(nodes[i])
			current.Left = &TreeNode{Val: leftVal}
			queue = append(queue, current.Left)
		}
		i++
		if i >= len(nodes) {
			break
		}
		if nodes[i] != "#" {
			rightVal := parseInt(nodes[i])
			current.Right = &TreeNode{Val: rightVal}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}
func parseInt(s string) int {
	n := 0
	sign := 1

	for i, char := range s {
		if i == 0 && char == '-' {
			sign = -1
		} else if char >= '0' && char <= '9' {
			n = n*10 + int(char-'0')
		}
	}

	return n * sign
}
