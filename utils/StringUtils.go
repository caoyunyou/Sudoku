package utils

// StrLen 计算字符串的字符个数
func StrLen(str string) int {
	return len([]rune(str))
}
