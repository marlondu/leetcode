package problem1_100

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 题解
// 这道题按照我们正常的手算方法即可。
// 从地位开始相加，进位carry == 0 或 1, 当然两个链表长度可能不同

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	l3 := head
	carry := 0
	for l1 != nil || l2 != nil {
		var sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		l3.Val = sum % 10
		carry = sum / 10
		if carry == 1 || l1 != nil || l2 != nil {
			l3.Next = new(ListNode)
			l3 = l3.Next
		}
	}
	if carry == 1 {
		l3.Val = 1
	}
	return head
}
