package Algorithm

import "fmt"

func decToBase(dec int, base int)  string{
	var res string
	for dec > 0 {
		rem := dec % base
		res = fmt.Sprintf("%X%s", rem, res)
		dec = dec/base
	}
	return res
}

func decToBase_2(dec int, base int)  string{
	var res string
	const charset = "0123456789ABCDEF"
	for dec > 0 {
		rem := dec % base
		res = fmt.Sprintf("%s%s", string(charset[rem]), res)
		dec = dec/base
	}
	return res
}
