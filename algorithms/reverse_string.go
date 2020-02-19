package Algorithm

import "strings"

func ReverseString(str string) string {
	var res string
	for i := len(str)-1; i>=0; i-- {
		res = res + string(str[i])
	}
	return res
}

func ReverseStringBuilder(str string) string {
	var sb strings.Builder
	for i := len(str)-1; i>=0; i-- {
		sb.WriteByte(str[i])
	}
	return sb.String()
}

func ReverseString_2(str string) string {
	var res string
	for i := 0; i< len(str); i++ {
		res = string(str[i]) + res
	}
	return res
}

func ReverseString_3(str string) string {
	var res string
	for _, val := range str {
		res = string(val) + res
	}
	return res
}

