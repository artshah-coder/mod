package mod

func Swap(a, b int) (int, int) {
	return b, a
}

func Sort2Int(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
