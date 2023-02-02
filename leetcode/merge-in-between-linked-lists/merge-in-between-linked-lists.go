package main

import "fmt"

func main() {
	res := mergeInBetween(&ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	},
		3, 4, &ListNode{
			Val: 1000000,
			Next: &ListNode{
				Val: 1000001,
				Next: &ListNode{
					Val:  1000002,
					Next: nil,
				},
			},
		})

	fmt.Println(res)
}

// ListNode  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	pre := list1
	// 第一次遍历找到前置节点
	for i := 1; i < a; i++ {
		pre = pre.Next
	}

	post := pre

	// 第二次遍历找到后置节点
	for i := 0; i < b-a+2; i++ {
		post = post.Next
	}

	// 前置节点的下一个是list2
	pre.Next = list2

	for list2.Next != nil {
		list2 = list2.Next
	}

	// list2的最后一个节点的下一个是后置节点
	list2.Next = post

	return list1
}
