package util

func isPowerOfTwo(x int64) bool {
	return x != 0 && (x&(x-1) == 0)
}
