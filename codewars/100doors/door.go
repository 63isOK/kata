package door

func doorFor() []bool {
	ret := [101]bool{}

	for i := 1; i < 101; i++ {
		for j := i; j < 101; j += i {
			ret[j] = !ret[j]
		}
	}

	return ret[1:]
}

func doorFactor() []bool {
	ret := make([]bool, 100)

	for i := 1; i <= 10; i++ {
		ret[i*i-1] = true
	}

	return ret
}
