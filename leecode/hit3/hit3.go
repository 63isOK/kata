package hit3

import (
	"fmt"
	"sort"
)

func hit3(arr []int) [][]int {
	ret := make([][]int, 0)
	if len(arr) < 3 {
		return ret
	}

	hash := make(map[int]int, len(arr))
	for _, v := range arr {
		hash[v]++
	}

	unique := make(map[string]struct{})

	for i, one := range arr {
		for _, two := range arr[i+1:] {
			three := 0 - one - two
			if _, ok := hash[three]; ok {

				// unique
				v1, v2, v3 := one, two, three
				if v1 > v2 {
					v1, v2 = v2, v1
				}
				if v2 > v3 {
					v2, v3 = v3, v2
				}
				if v1 > v2 {
					v1, v2 = v2, v1
				}
				tag := fmt.Sprintf("%d,%d,%d", v1, v2, v3)
				// tag := strconv.Itoa(v1) + "," + strconv.Itoa(v2) + "," + strconv.Itoa(v3)
				if _, ok := unique[tag]; ok {
					continue
				}

				// three != one && three != two -> ok
				// one === two === three && hash[0]>=3 -> ok
				// one == three && hash[one] >= 2 -> ok
				// two == three && hash[two] >= 2 -> ok
				// other -> fail
				if three != one && three != two {
				} else if 0 == one && 0 == two && 0 == three {
					if hash[0] < 3 {
						continue
					}
				} else if one == three && hash[one] >= 2 {
				} else if two == three && hash[two] >= 2 {
				} else {
					continue
				}

				ret = append(ret, []int{v1, v2, v3})
				unique[tag] = struct{}{}
			}
		}
	}

	sort.SliceStable(ret, func(i, j int) bool {
		x1, y1, z1 := ret[i][0], ret[i][1], ret[i][2]
		x2, y2, z2 := ret[j][0], ret[j][1], ret[j][2]

		return x1 < x2 || (x1 == x2 && y1 < y2) ||
			(x1 == x2 && y1 == y2 && z1 <= z2)
	})

	return ret
}

func hit32(nums []int) [][]int {

	ret := make([][]int, 0)
	if len(nums) < 3 {
		return ret
	}

	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}

	unique := make(map[string]struct{})

	for i, one := range nums {
		for _, two := range nums[i+1:] {
			three := 0 - one - two
			if _, ok := hash[three]; ok {

				// unique
				v1, v2, v3 := one, two, three
				if v1 > v2 {
					v1, v2 = v2, v1
				}
				if v2 > v3 {
					v2, v3 = v3, v2
				}
				if v1 > v2 {
					v1, v2 = v2, v1
				}

				tag := fmt.Sprintf("%d,%d,%d", v1, v2, v3)
				if _, ok := unique[tag]; ok {
					continue
				}

				// check
				if three != one && three != two {
				} else if three == one && three == two && hash[three] >= 3 {
				} else if three == one && hash[three] >= 2 {
				} else if three == two && hash[three] >= 2 {
				} else {
					continue
				}

				// match
				ret = append(ret, []int{v1, v2, v3})
				unique[tag] = struct{}{}
			}
		}
	}

	return ret
}

// func main() {
// 	f, err := os.Create("cpu.prof")
// 	if err != nil {
// 		fmt.Println("access cpu.prof failed.")
// 		return
// 	}
// 	defer f.Close()

// 	if err = pprof.StartCPUProfile(f); err != nil {
// 		fmt.Println("start cpu profile failed.")
// 		return
// 	}
// 	defer pprof.StopCPUProfile()

// 	input_arr := []int{3, 4, 2, -1, -1, 0, -2, 1, 0, 0, 0}
// 	fmt.Println(hit3(input_arr))
// }
