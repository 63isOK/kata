package hit2

// Hit2 is a brute force hit
func Hit2(arr []int, target int) []int {
	var i, j int
	var one, two int
	for i, one = range arr {
		for j, two = range arr[i+1:] {
			if one+two == target {
				goto FOUND
			}
		}
	}

FOUND:
	return []int{i, j + i + 1}
}

func Hit2HashMap(arr []int, target int) []int {
	hash := make(map[int]int, len(arr))
	for i, v := range arr {
		hash[v] = i
	}

	for i, v := range arr {
		if j, ok := hash[target-v]; ok {
			return []int{i, j}
		}
	}

	return []int{}
}
