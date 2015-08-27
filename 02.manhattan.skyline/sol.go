package solution

func Solution(H []int) int {
	in := false
	levels := make([]int, 0, len(H))
	res := 0
	curlev := -1
	for lev := range H {
		if lev > curlev {
			levels = append(levels, lev)
		}
		if lev < curlev {
			levels, in = isIn(levels, lev)
			if !in {
				res += 1
			}
		}
	}
	return res
}

func isIn(slice []int, lev int) ([]int, bool) {
	in := false
	cutIdx := len(slice) - 1
	for ; cutIdx >= 0; cutIdx-- {
		curlev := slice[cutIdx]
		if curlev == lev {
			in = true
			break
		}
		if curlev < lev {
			break
		}
	}
	slice = slice[:cutIdx]
	return slice, in
}
