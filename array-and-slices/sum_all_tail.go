package arrayandslices

func SumAllTails(numbersToSum ...[]int) []int {
	var res []int

	for _, nums := range numbersToSum {
		if len(nums) <= 1 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(nums[1:]))
		}
	}

	return res
}
