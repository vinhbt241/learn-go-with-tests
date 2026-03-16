package arrayandslices

func SumAll(numbersToSum ...[]int) []int {
	var res []int

	for _, n := range numbersToSum {
		res = append(res, Sum(n))
	}

	return res
}
