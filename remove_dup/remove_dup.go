package main

func removeDup(ds []string) []string {
	switch len(ds) {
	case 0:
		return []string{}
	case 1:
		return ds
	default:
		var dm = make(map[string]bool)
		var res []string

		for _, element := range ds {
			dm[element] = true
		}

		for k := range dm {
			res = append(res, k)
		}
		return res
	}
}
