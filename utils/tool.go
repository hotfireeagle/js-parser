package utils

import "unicode"

// 判断是否是ECMAScript规范下合法的变量标识符首字母
func CheckIsJSWordPrefix(n byte) bool {
	special := map[byte]bool{
		'_': true,
		'$': true,
	}
	return unicode.IsLetter(rune(n)) || special[n]
}

// 判断是否是合法的后缀字符
func CheckIsJSWordSuffix(n byte) bool {
	return CheckIsJSWordPrefix(n) || unicode.IsDigit(rune(n))
}
