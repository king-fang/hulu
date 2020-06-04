package strs

// 字符串长度
func RuneLen(s string) int {
	bt := []rune(s)
	return len(bt)
}
