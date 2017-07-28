package util

func splitI(target int) []int {
	var res []int

	if target < 10 {
		return append(res, target)
	}

	res = append(res, target%10)
	return append(res, splitI(target/10)...)
}
