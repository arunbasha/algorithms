package Algorithm

func GCD(num1 int, num2 int) int {
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	return num1
}
