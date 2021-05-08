package Challenge3_Count_of_Structurally_Unique_Binary_Search_Trees

/*
已经在challenge2中实现
ref: https://leetcode-cn.com/problems/unique-binary-search-trees/
*/
func numTrees(n int) int {
	var bfs func(int, int) int
	var m = make(map[Range]int)
	bfs = func(start, end int) int {
		var res = 0
		if start == end {
			return 1
		}
		if v, ok := m[Range{start, end}]; ok {
			return v
		}
		for i := start; i <= end; i++ {
			if i == start {
				res += bfs(i+1, end)
				continue
			}
			if i == end {
				res += bfs(start, i-1)
				continue
			}
			left := bfs(start, i-1)
			right := bfs(i+1, end)
			res += left * right
		}
		m[Range{start, end}] = res
		return res
	}

	return bfs(1, n)
}

type Range struct {
	start int
	end   int
}
