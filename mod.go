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

func Sort3Int(a, b, c int) (v1, v2, v3 int) {
	switch {
	case c >= b && b >= a:
		v1, v2, v3 = a, b, c
	case c >= b && c >= a && b <= a:
		v1, v2, v3 = b, a, c
	case b >= a && a >= c:
		v1, v2, v3 = c, a, b
	case b >= a && b >= c && a <= c:
		v1, v2, v3 = a, c, b
	case a >= b && b >= c:
		v1, v2, v3 = c, b, a
	default:
		v1, v2, v3 = b, c, a
	}
	return
}
