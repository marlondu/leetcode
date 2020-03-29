package problem1_100

import "math"

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 题解 关键字 滑动窗口
// 这其实是一个滑动窗口的问题，
// 假如我有一个滑动窗口，窗口边界可以改变，而且这个窗口中框住的数字是不重复的，那么问题就变成求这个窗口的长度
// 假如我有一个窗口，左右边界分别是[i,j]. 窗口的长度是j - i + 1。起始i,j均为0
// 然后j向后移动，并检查字符s[j]是否在字符串s[i, j) (左闭右开)之中. 如果不存在继续向后移动(窗口长度 j - i + 1).
// 如果存在重复，加入字符s[j]在s[i,j)中的位置是j', 则i向前移动到位置 j' + 1
// 这时我们想到使用map来保存字符和索引的映射，来降低查找重复的时间复杂度到O(1)

func longestLengthOfSubstring(s string) int {
	var ans, l = 0, len(s)
	//indexes := make(map[uint8]int, l) // 保存字符索引
	// 如果我们假设我们的字符为ascii码，码值最大128,我们直接用128个整数组表示
	var indexes = [128]int{}
	for i,j := 0, 0; j < l; j++ {
		// 查看s[j]的索引在不在
		idx := indexes[s[j]]
		// 依照题解中所说，如果s[j]中在之前的字符串存在中，则i取 j' + 1
		// 因为默认indexes中默认值是0,不存在应该是-1才对，因此我们直接存j + 1, 取着方便
		i = int(math.Max(float64(idx), float64(i)))
		ans = int(math.Max(float64(ans), float64(j - i + 1)))
		indexes[s[j]] = j + 1
	}
	return ans
}
