package hit2

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
