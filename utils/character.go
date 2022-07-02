package utils

// 判断当前字符是否变量标识符的一部分
func IsIdentifierPart(cp byte) bool {
	return (cp == 0x24) || (cp == 0x5F) || // $ _
		(cp >= 0x41 && cp <= 0x5A) || // A..Z
		(cp >= 0x61 && cp <= 0x7A) || // a..z
		(cp >= 0x30 && cp <= 0x39) // 0..9
}

// 判断是否变量标识符的开头部分
func IsIdentifierBegin(cp byte) bool {
	return (cp == 0x24) || (cp == 0x5F) || // $ _
		(cp >= 0x41 && cp <= 0x5A) || // A..Z
		(cp >= 0x61 && cp <= 0x7A) // a..z
}

// 判断是否是十进制数字
func IsDecimalDigit(cp byte) bool {
	return (cp >= 0x30 && cp <= 0x39) // 0..9
}

func IsNumberBegin(cp byte) bool {
	return IsDecimalDigit(cp) || cp == '.'
}

func IsNumberPart(cp byte) bool {
	return IsNumberBegin(cp)
}
