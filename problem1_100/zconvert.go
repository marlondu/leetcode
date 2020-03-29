package problem1_100

import (
	"fmt"
)

// 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
//
//比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//
//L   C   I   R
//E T O E S I I G
//E   D   H   N
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);
//示例 1:
//
//输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
//示例 2:
//
//输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zigzag-conversion
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 题解
// row
// --------------------------------> x
// |(0,0)
// |(0,1)
// |(0,2)(1,2)
// |(0,3)
// V
// y

//
const (
	DOWN = 0 // 垂直向下
	UP   = 1 // 斜向上
)

func convert(s string, rows int) string {
	direct := DOWN // 路线方向
	i, j := 0, 0   // 横纵坐标
	// key : row
	positions := make([][]int, rows)
	for n := 0; n < len(s); n++ {
		positions[j] = append(positions[j], n)
		fmt.Printf("[%d, %d] %c %d\n", i, j, s[n], len(positions[j]))
		if direct == DOWN && j < (rows-1) {
			j++
		} else if direct == UP && j > 0 {
			i++
			j--
		}
		if j == (rows - 1) {
			direct = UP
		} else if j == 0 {
			direct = DOWN
		}
	}
	chars := make([]byte, len(s))
	n := 0
	for i = 0; i < rows; i++ {
		for j = 0; j < len(positions[i]); j++ {
			chars[n] = s[positions[i][j]]
			n++
		}
	}
	return string(chars)
}
