package util

func unionSlice(sa, sb []int) []int {
	if sa == nil || sb == nil {
		return nil
	}

	if len(sa) == 0 || len(sb) == 0 {
		return []int{}
	}

	var res []int
	var um = make(map[int]bool)

	for _, a := range sa {
		if _, ok := um[a]; !ok {
			um[a] = true
		}
	}

	for _, b := range sb {
		if _, ok := um[b]; ok {
			res = append(res, b)
		}
	}

	return res
}
