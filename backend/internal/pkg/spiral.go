package pkg

func SpiralMatrix(rows int, cols int, orderSequence []int64) []int64 {
	size := rows * cols
	left := 0
	top := 0
	right := cols - 1
	bottom := rows - 1
	s := make([]int64, size)
	i := 0

	for left < right && i < size {
		var c int
		for c = left; c <= right; c++ {
			s[top*cols+c] = orderSequence[i]
			i++
		}
		top++
		if i >= size {
			break
		}

		var r int
		for r = top; r <= bottom; r++ {
			s[r*cols+right] = orderSequence[i]
			i++
		}
		right--

		for c = right; c >= left; c-- {
			s[bottom*cols+c] = orderSequence[i]
			i++
		}
		bottom--

		if i >= size {
			break
		}
		for r = bottom; r >= top; r-- {
			s[r*cols+left] = orderSequence[i]
			i++
		}
		left++
	}

	if cols%2 == 1 {
		for r := top; r <= bottom; r++ {
			s[r*cols+right] = orderSequence[i]
			i++
		}
	}
	return s
}
