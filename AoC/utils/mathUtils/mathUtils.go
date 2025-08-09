package mathUtils

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func SumOfInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}
