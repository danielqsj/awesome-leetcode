package numberComplement

func numberComplement(num int) int {
	result := 0
	for i, j := num, 1; i > 0; i, j = i>>1, j*2 {
		if i&1 == 0 {
			result = result + j
		}
	}
	return result
}
