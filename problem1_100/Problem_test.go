package problem1_100

import (
	"fmt"
	"testing"
)

// 01
func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums, 18)
	if res == nil {
		t.Errorf("result is nil")
	} else {
		t.Log(fmt.Sprintf("result:%v\n", res))
	}
}

// 02
func TestAddTwoNumbers(t *testing.T) {
	l1, l2 := new(ListNode), new(ListNode)
	h1, h2 := l1, l2
	l1.Val, l1.Next = 2, new(ListNode)
	l1 = l1.Next
	l1.Val, l1.Next = 4, new(ListNode)
	l1 = l1.Next
	l1.Val = 3

	l2.Val, l2.Next = 5, new(ListNode)
	l2 = l2.Next
	l2.Val, l2.Next = 6, new(ListNode)
	l2 = l2.Next
	l2.Val = 4

	l3 := addTwoNumbers(h1, h2)
	for l3 != nil {
		t.Logf("-->%d", l3.Val)
		l3 = l3.Next
	}
}

// 03
func TestLongestLengthOfSubstring(t *testing.T) {
	s := "abcabcbb"
	t.Logf("string: %s, longestLengthOfSubstring: %d" , s, longestLengthOfSubstring(s))

	s = "bbbbbb"
	t.Logf("string: %s, longestLengthOfSubstring: %d" , s, longestLengthOfSubstring(s))

	s = "pwwkew"
	t.Logf("string: %s, longestLengthOfSubstring: %d" , s, longestLengthOfSubstring(s))
}

// 04
func TestMedianOfTwoArrays(t *testing.T) {
	a := []int{1, 2, 5, 7,  8, 11, 13}
	b := []int{3, 4, 6, 9, 10, 12}
	t.Log(findMedianSortedArrays(a, b))

	a = []int{1, 2}
	b = []int{3}
	t.Log(findMedianSortedArrays(a, b))

	a = []int{1, 2}
	b = []int{1, 2}
	t.Log(findMedianSortedArrays(a, b))
}

// 05
func TestLonestPalindrome(t *testing.T) {
	s := "babad"
	t.Log(longestPalindrome(s))

	s = "cbbd"
	t.Log(longestPalindrome(s))

	s = "ccc"
	t.Log(longestPalindrome(s))

	s = "aaabaaaa"
	t.Log(longestPalindrome(s))
}

// 06
func TestZConvert(t *testing.T)  {
	t.Log(convert("LEETCODEISHIRING", 4))
}

func TestNoReverse(t *testing.T) {
	t.Log(reverse(-123))
}
