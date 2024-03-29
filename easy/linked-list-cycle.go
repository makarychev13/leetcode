package leetcode

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := map[*ListNode]bool{}
	hasCycle := false

	for head != nil {
		if _, ok := m[head]; ok {
			hasCycle = true
			break
		}

		m[head] = true
		head = head.Next
	}

	return hasCycle
}

//более оптимальное, но менее интуитивное решение с затратами O(1) по памяти.
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast.Next != nil && fast.Next.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow == fast
}
