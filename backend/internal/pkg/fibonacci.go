package pkg

func FibonacciSequence(n int) []int64 {
	size := n - 1
	f := make([]int64, size+1, size+2)
	if size < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= size; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}
