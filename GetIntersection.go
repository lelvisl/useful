package useful

/*

1 - те что есть только в левом массиве
2 - те что только в правом
3 - пересечение
4 - это те что в пересечение не попали
*/

func GetIntersection(a []string, b []string, mode byte) []string {
	m := make(map[string]byte)

	for _, k := range a {
		m[k] += 1
	}

	for _, k := range b {
		m[k] += 2
	}

	result := []string{}

	if mode == 4 {
		for k, v := range m {
			if v < 3 {
				result = append(result, k)
			}
		}
	} else {
		for k, v := range m {
			if v == mode {
				result = append(result, k)
			}
		}
	}

	return result
}
