package util

// split target by n
func splitByN(target, n int) [][2]int {
	var res [][2]int

	if n <= 0 {
		return [][2]int{[2]int{0, target}}
	}

	for i := 0; i < target; i = i + n {
		var start, end int
		start = i
		end = i + n

		if end >= target {
			end = target
		}

		res = append(res, [2]int{start, end})

	}
	return res
}
