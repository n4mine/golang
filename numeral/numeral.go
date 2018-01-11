package numeral

// IsPowerOfTwo
func IsPowerOfTwo(x int64) bool {
	return x != 0 && (x&(x-1) == 0)
}

// SplitByN split target by n
func SplitByN(target, n int) [][2]int {
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

// SplitI
func SplitI(target int) []int {
	var res []int

	if target < 10 {
		return append(res, target)
	}

	res = append(res, target%10)
	return append(res, SplitI(target/10)...)
}
