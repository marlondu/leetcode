package util

// BitMap bitmap struct
type BitMap struct {
	initCap int      // 位图初始容量
	words   []uint64 // 位图数组
}

// NewBitMap 位图长度
func NewBitMap(cap int) *BitMap {
	l, m := cap/64, cap%64
	if m > 0 {
		l++
	}
	return &BitMap{initCap: cap, words: make([]uint64, l)}
}

// Add 这只数字标记
func (bm *BitMap) Add(num int) {
	// 查找位置
	word, bit := num/64, num%64
	if word > len(bm.words) {
		bm.words = append(bm.words, 0)
	}
	bm.words[word] |= 1 << bit
}

// Has judge num if has exists
func (bm *BitMap) Has(num int) bool {
	word, bit := num/64, num%64
	return word < len(bm.words) && (bm.words[word]&(1<<bit)) != 0
}
