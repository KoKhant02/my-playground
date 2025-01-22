package algorithm

func Reverse(x int) int {
	const (
		int32Min = -1 << 31  // -2147483648
		int32Max = 1<<31 - 1 // 2147483647
	)

	result := 0
	for x != 0 {
		//remainder
		digit := x % 10
		//quotient
		x /= 10
		//Positive
		if result > int32Max || result < int32Min {
			return 0
		}

		result = result*10 + digit
	}

	return result
}
