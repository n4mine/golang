package util

// split m per n
func splitN(m, n int) [][]int {
	var res [][]int

	if n <= 0 {
		return [][]int{[]int{0, m}}
	}

	for i := 0; i < m; i = i + n {
		var start, end int
		start = i
		end = i + n

		if end >= m {
			end = m
		}

		res = append(res, []int{start, end})

	}
	return res
}
