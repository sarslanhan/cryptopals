package util

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func FixedXor(a []byte, b []byte) []byte {
	l := min(len(a), len(b))

	result := make([]byte, l)

	for i := 0; i < l; i++ {
		result[i] = a[i] ^ b[i]
	}

	return result
}
