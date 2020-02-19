package Algorithm

func fizzBuzz(num int) string {
	return printFizzBuss(num)
}

func printFizzBuss(n int) string{
	switch {
	case isDivisible(n, 15):
		return "Fizz Buzz"
	case isDivisible(n, 3):
		return "Fizz"
	case isDivisible(n, 5):
		return "Buzz"
	default:
		return string(n)
	}
}

func isDivisible(num int, divider int) bool {
	if num % divider == 0 {
		return true
	}
	return false
}