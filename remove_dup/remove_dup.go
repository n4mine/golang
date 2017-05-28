package util

func removeDup(ds []string) []string {
	switch len(ds) {
	case 0:
		return []string{}
	case 1:
		return ds
	default:
		var um = make(map[string]bool)
		var res []string

		for _, element := range ds {
			if _, ok := um[element]; !ok {
				um[element] = true
				res = append(res, element)
			}
		}

		return res
	}
}
