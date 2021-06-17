package Challenge2

/*
* Minimum Height Trees
We are given an undirected graph that has characteristics of a k-ary tree.
In such a graph, we can choose any node as the root to make a k-ary tree.
The root (or the tree) with the minimum height will be called Minimum Height Tree (MHT).
There can be multiple MHTs for a graph.
In this problem, we need to find all those roots which give us MHTs.
Write a method to find all MHTs of the given graph and return a list of their roots.

Input: vertices: 5, Edges: [[0, 1], [1, 2], [1, 3], [2, 4]]
Output:[1, 2]
Explanation: Choosing '1' or '2' as roots give us MHTs. In the below diagram, we can see that the
height of the trees with roots '1' or '2' is three which is minimum.

Input: vertices: 4, Edges: [[0, 1], [0, 2], [2, 3]]
Output:[0, 2]
Explanation: Choosing '0' or '2' as roots give us MHTs. In the below diagram, we can see that the
height of the trees with roots '0' or '2' is three which is minimum.

Input: vertices: 4, Edges: [[0, 1], [1, 2], [1, 3]]
Output:[1]

ref: https://leetcode-cn.com/problems/minimum-height-trees/
*/

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	var (
		inDegree = make(map[int]int)
		graph    = make(map[int][]int)
	)
	for i := 0; i < n; i++ {
		inDegree[i] = 0
		graph[i] = []int{}
	}

	for _, v := range edges {
		n1, n2 := v[0], v[1]
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
		inDegree[n1]++
		inDegree[n2]++
	}

	leaves := make([]int, 0)
	for k, v := range inDegree {
		if v == 1 {
			leaves = append(leaves, k)
		}
	}

	totalNodes := n
	for totalNodes > 2 {
		leavesSize := len(leaves)
		totalNodes -= leavesSize
		for i := 0; i < leavesSize; i++ {
			vertex := leaves[0]
			leaves = leaves[1:]
			children := graph[vertex]
			for _, v := range children {
				inDegree[v]--
				if inDegree[v] == 1 {
					leaves = append(leaves, v)
				}
			}
		}
	}

	return leaves
}
