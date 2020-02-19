package Algorithm

func fibonacci(num int) int {
	if num < 2 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func fibonacci_2(num int) int {
	if num < 2 {
		return num
	}
	nMinus2 := 0
	nMinus1 := 1
	var cur int
	for n := 2; n <= num; n++{
		cur = nMinus2 + nMinus1
		nMinus2 = nMinus1
		nMinus1 = cur
	}
	return cur
}
