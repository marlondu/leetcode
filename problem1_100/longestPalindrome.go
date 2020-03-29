package problem1_100

import "math"

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 题解 关键字 动态规划
// 假如我们知道一个字串 ”aba“ 是回文字串，则cabac也一定是回文字串
// 也就是 s[i + 1, j - 1] 是回文，且 s[i] == s[j] 则s[i,j]也一定是回文字串
// 这产生了一个直观的动态规划解法，我们首先初始化一字母和二字母的回文，然后找到所有三字母回文，并依此类推…
//
//复杂度分析
//
//时间复杂度：O(n^2)O(n
//2
// )，这里给出我们的运行时间复杂度为 O(n^2)O(n
//2
// ) 。
//
//空间复杂度：O(n^2)O(n
//2
// )，该方法使用 O(n^2)O(n
//2
// ) 的空间来存储表。
//
func longestPalindromeV3(s string) string {
	// 感觉就是中心扩展法
	return longestPalindromeV2(s)
}

func longestPalindrome(s string) string {
	return longestPalindromeV2(s)
}

// 遍历法，列举所有可能比较最长的
func longestPalindromeV1(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	ll := 1             // 最长回文长度
	lls := string(s[0]) // 保存最长回文字串
	// i 指针从左开始向右移动
	for i := 0; i < l-ll; i++ {
		j := l - 1 // j指针从右向左移动，如果中间的[i,j]字符串比ll小，则无需计算
		for ; j > i && (j+1-i) > ll; j-- {
			if s[j] == s[i] {
				if isPalindrome(s[i : j+1]) {
					if len(s[i:j+1]) > ll {
						lls = s[i : j+1]
						ll = j + 1 - i
					}
				}
			}
		}
	}
	return lls
}

// 判断是否是回文子串
func isPalindrome(s string) bool {
	l := len(s)
	// 0 1 2 3 : 4
	// 0 1 2 : 3
	var yes = true
	for i := 0; i <= (l-1)/2; i++ {
		if s[i] != s[l-1-i] {
			yes = false
			break
		}
	}
	return yes
}

// 中心扩展算法
// 事实上，只需使用恒定的空间，我们就可以在 O(n^2)O(n
//2
// ) 的时间内解决这个问题。
//
//我们观察到回文中心的两侧互为镜像。因此，回文可以从它的中心展开，并且只有 2n - 12n−1 个这样的中心。
//
//你可能会问，为什么会是 2n - 12n−1 个，而不是 nn 个中心？
//原因在于所含字母数为偶数的回文的中心可以处于两字母之间（例如 \textrm{“abba”}“abba” 的中心在两个 \textrm{‘b’}‘b’ 之间）。
//
func longestPalindromeV2(s string) string {
	if len(s) <= 1 {
		return s
	}
	var start, end = 0, 0
	for i := 0; i < len(s); i++ {
		var len1 = expandCenterRound(s, i, i)   // 回文字串长度奇数情况
		var len2 = expandCenterRound(s, i, i+1) // 回文字串长度偶数情况
		var l = int(math.Max(float64(len1), float64(len2)))
		if l > (end - start) {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}

// 返回扩展后的长度
func expandCenterRound(s string, left int, right int) int {
	var L, R = left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}
