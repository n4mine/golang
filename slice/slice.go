package slice

// Uniq 去重
func Uniq(ds []string) []string {
	switch len(ds) {
	case 0:
		return []string{}
	case 1:
		return ds
	default:
		var um = make(map[string]struct{})
		var res []string

		for _, element := range ds {
			if _, ok := um[element]; !ok {
				um[element] = struct{}{}
				res = append(res, element)
			}
		}

		return res
	}
}

// Union 并集
func Union(sa, sb []int) []int {
	if sa == nil || sb == nil {
		return nil
	}

	if len(sa) == 0 || len(sb) == 0 {
		return []int{}
	}

	var res []int
	var um = make(map[int]struct{})

	for _, a := range sa {
		if _, ok := um[a]; !ok {
			um[a] = struct{}{}
		}
	}

	for _, b := range sb {
		if _, ok := um[b]; ok {
			res = append(res, b)
		}
	}

	return res
}
