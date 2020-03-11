package rangeint

import (
	"fmt"
	"strconv"
)

func solution(in []int) (ret string) {
	if len(in) == 0 {
		return ""
	}

	ret = strconv.Itoa(in[0])

	click := 0
	for i := 1; i < len(in); i++ {
		if in[i]-1 != in[i-1] {
			if click == 0 {
				ret += ","
				ret += strconv.Itoa(in[i])
			} else if click == 1 {
				ret += ","
				ret += strconv.Itoa(in[i-1])
				ret += ","
				ret += strconv.Itoa(in[i])
				click = 0
			} else {
				ret += "-"
				ret += strconv.Itoa(in[i-1])
				ret += ","
				ret += strconv.Itoa(in[i])
				click = 0
			}
		} else {
			click++
			if i == len(in)-1 {
				if click == 1 {
					ret += ","
					ret += strconv.Itoa(in[i])
					break
				}
				ret += "-"
				ret += strconv.Itoa(in[i])
			}
		}
	}

	return
}

func solution2(in []int) (ret string) {
	if len(in) == 0 {
		return ""
	}

	ret = fmt.Sprintf("%d", in[0])

	click := 0
	for i := 1; i < len(in); i++ {
		if in[i]-1 != in[i-1] {
			if click == 0 {
				ret += fmt.Sprintf(",%d", in[i])
			} else if click == 1 {
				ret += fmt.Sprintf(",%d,%d", in[i-1], in[i])
				click = 0
			} else {
				ret += fmt.Sprintf("-%d,%d", in[i-1], in[i])
				click = 0
			}
		} else {
			click++
			if i == len(in)-1 {
				if click == 1 {
					ret += fmt.Sprintf(",%d", in[i])
				} else {
					ret += fmt.Sprintf("-%d", in[i])
				}
				break
			}
		}
	}

	return
}
