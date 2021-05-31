package kthsmallestnumberinasortedmatrix

func kthSmallestByBinarySearch(matrix [][]int, k int) int {
	var (
		n     = len(matrix)
		start = matrix[0][0]
		end   = matrix[n-1][n-1]
	)
	for start < end {
		mid := start + (end-start)/2
		pair := Pair{matrix[0][0], matrix[n-1][n-1]}
		count := countLessEq(matrix, mid, &pair)
		if count == k {
			return pair.Small
		}

		if count < k {
			start = pair.Large
		} else {
			end = pair.Small
		}
	}

	return start
}

func countLessEq(matrix [][]int, mid int, pair *Pair) int {
	var (
		count = 0
		n     = len(matrix)
		row   = n - 1
		col   = 0
	)
	for row >= 0 && col < n {
		if matrix[row][col] > mid {
			pair.Large = min(pair.Large, matrix[row][col])
			row--
		} else {
			pair.Small = max(pair.Small, matrix[row][col])
			col++
			count += row + 1
		}
	}

	return count
}

type Pair struct {
	Small int
	Large int
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
