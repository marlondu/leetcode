package problem1_100

import "math"

// 数字反转
// 321
// 123

func reverse(x int) int {
	// rs = 0
	// 321 % 10 == 1 rs = (rs * 10) + 321%10
	// 321 / 10 % 10 == 2 ((0 * 10) + 1) * 10 + 32%2
	// 321 / 100 % 10 == 3 (((0 * 10) + 1) * 10 + 32%2) * 10 + 3%10
	rs := 0
	for x != 0 {
		var pop = x%10
		x = x/10
		// 假如rs是正数，如果rs * 10 + x%10溢出，一定有 rs > math.Max/10, math.MaxInt%10 == 7
		if rs > math.MaxInt32/10 || (rs == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		// -2147483648到2147483647
		if rs < math.MinInt32/10 || (rs == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rs = rs * 10 + pop
	}
	return rs
}
