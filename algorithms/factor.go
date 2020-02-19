package Algorithm

func factor(primes []int, num int) []int {
	var factor []int

	for _, val := range primes {
		for num % val == 0 {
			factor = append(factor, val)
			num = num / val
		}
	}
	if num > 1 {
		factor = append(factor, num)
	}

	return factor
}