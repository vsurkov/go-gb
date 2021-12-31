package mySort

func insertsSort(slc []int) []int {
	for i := 1; i < len(slc); i++ {
		for j := i; j > 0 && slc[j-1] > slc[j]; j-- {
			slc[j-1], slc[j] = slc[j], slc[j-1]
		}
	}
	return slc
}
