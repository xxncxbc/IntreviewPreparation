package ReverseLinkedList

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddNode(n *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val, Next: n}
	return newNode
}

func GenerateList(vals ...int) *ListNode {
	var node *ListNode
	for i := len(vals) - 1; i >= 0; i-- {
		node = AddNode(node, vals[i])
	}
	return node
}

func ReverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func StringList(head *ListNode) string {
	builder := strings.Builder{}
	for head != nil {
		i := strconv.Itoa(head.Val)
		builder.WriteString(i)
		builder.WriteString("->")
		head = head.Next
	}
	builder.WriteString("nil")
	return builder.String()
}
