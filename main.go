package main

import (
	"fmt"
	"github.com/marlondu/leetcode/util"
	"math"
	"strings"
)

func main() {
	var s = "au"
	l1 := lengthOfLongestSubstring(s)
	l2 := lengthOfLongestSubstringV3(s)
	fmt.Printf("s:%s longest:%d, longestV2:%d\n", s, l1, l2)

	s = "aniewqmoxkjwpymqorluxedvywhcoghotpusfgiestckrpaigocfufbubiyrrffmwaeeimidfnnzcphkflpbqsvtdwludsg"
	l1 = lengthOfLongestSubstring(s)
	l2 = lengthOfLongestSubstringV3(s)
	fmt.Printf("s:%s longest:%d, longestV2:%d\n", s, l1, l2)

	s = "abcabcbb"
	l1 = lengthOfLongestSubstring(s)
	l2 = lengthOfLongestSubstringV3(s)
	fmt.Printf("s:%s longest:%d, longestV2:%d\n", s, l1, l2)

	s = "bbbbbb"
	l1 = lengthOfLongestSubstring(s)
	l2 = lengthOfLongestSubstringV3(s)
	fmt.Printf("s:%s longest:%d, longestV2:%d\n", s, l1, l2)

	s = "au"
	l1 = lengthOfLongestSubstring(s)
	l2 = lengthOfLongestSubstringV3(s)
	fmt.Printf("s:%s longest:%d, longestV2:%d\n", s, l1, l2)
}

// 滑动窗口
func lengthOfLongestSubstringV3(s string) int {
	var maxLen = 0
	var indexes [128]int // 保存字符的索引,从1开始
	for i,j := 0, 0; j < len(s); j++ {
		// i j
		// 0 1 2 3 4 5
		// s d h h i
		i = int(math.Max(float64(indexes[s[j]]), float64(i)))
		maxLen = int(math.Max(float64(maxLen), float64(j - i + 1)))
		indexes[s[j]] = j + 1 // 当s[j]在当前窗口已经存在，则直接讲i置为 j + 1
	}
	return maxLen
}

// i j
// 0 1 2 3 4 5
// s d h h i
func lengthOfLonestSubstringV2(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var loIndex, hiIndex = 0, 1
	var maxLen = hiIndex - loIndex
	for loIndex < (len(s) - maxLen) && hiIndex < len(s) {
		// 先让hiIndex走
		if contains(s[loIndex:hiIndex], s[hiIndex]) {
			// 低座标前近,直到不含重复字符
			for contains(s[loIndex:hiIndex], s[hiIndex]) && loIndex < hiIndex {
				loIndex += 1
			}
		}
		hiIndex += 1
		var max = hiIndex - loIndex
		if maxLen < max {
			maxLen = max
		}
	}
	return maxLen;
}

func contains(s string, b byte) bool {
	return strings.IndexByte(s, b) >= 0
}

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	l := len(s) // 字符串长度
	// 因为求最大字串，从最大字串开始检查
	for ln := l; ln > 0; ln-- {
		// 从头依次检查长度为ln的字串
		for i := 0; i <= (l - ln); i++ {
			//fmt.Printf("length:%d, substring: %s\n", ln, s[i:ln+i])
			if !checkIfRepeat(s[i : ln+i]) {
				fmt.Println(s[i: ln + i])
				return ln
			}
		}
	}
	return 0
}

func checkIfRepeat(s string) bool {
	bitMap := util.NewBitMap(128)
	for i := 0; i < len(s); i++ {
		if bitMap.Has(int(s[i])) {
			return true
		} else {
			bitMap.Add(int(s[i]))
		}
	}
	return false
}
