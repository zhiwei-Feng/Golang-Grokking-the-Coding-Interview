package Topological_Sort

/*
* Topological Sort
Topological Sort of a directed graph (a graph with unidirectional edges) is a linear ordering of its vertices
such that for every directed edge (U, V) from vertex U to vertex V, U comes before V in the ordering.

Given a directed graph, find the topological ordering of its vertices.

Input: Vertices=4, Edges=[3, 2], [3, 0], [2, 0], [2, 1]
Output: Following are the two valid topological sorts for the given graph:
1) 3, 2, 0, 1
2) 3, 2, 1, 0

Input: Vertices=5, Edges=[4, 2], [4, 3], [2, 0], [2, 1], [3, 1]
Output: Following are all valid topological sorts for the given graph:
1) 4, 2, 3, 0, 1
2) 4, 3, 2, 0, 1
3) 4, 3, 2, 1, 0
4) 4, 2, 3, 1, 0
5) 4, 2, 0, 3, 1

Input: Vertices=7, Edges=[6, 4], [6, 2], [5, 3], [5, 4], [3, 0], [3, 1], [3, 2], [4, 1]
Output: Following are all valid topological sorts for the given graph:
1) 5, 6, 3, 4, 0, 1, 2
2) 6, 5, 3, 4, 0, 1, 2
3) 5, 6, 4, 3, 0, 2, 1
4) 6, 5, 4, 3, 0, 1, 2
5) 5, 6, 3, 4, 0, 2, 1
6) 5, 6, 3, 4, 1, 2, 0

There are other valid topological ordering of the graph too.
*/

func topologicSort(vertices int, edges [][]int) []int {
	var sortedOrder = make([]int, 0)
	if vertices == 0 {
		return sortedOrder
	}

	// a. Initialize the graph
	inDegree := make(map[int]int)
	graph := make(map[int][]int)
	for i := 0; i < vertices; i++ {
		inDegree[i] = 0
		graph[i] = make([]int, 0)
	}

	// b. Build the graph
	for _, v := range edges {
		parent, child := v[0], v[1]
		graph[parent] = append(graph[parent], child)
		inDegree[child]++
	}

	// c. Find all sources i.e., all vertices with 0 in-degrees
	sources := make([]int, 0)
	for k, v := range inDegree {
		if v == 0 {
			sources = append(sources, k)
		}
	}

	// d. For each source, add it to the sortedOrder and subtract one from all of its children's
	// in-degrees if a child's in-degree becomes zero, add it to the sources queue
	for len(sources) != 0 {
		vertex := sources[0]
		sources = sources[1:]
		sortedOrder = append(sortedOrder, vertex)
		children := graph[vertex]
		for _, child := range children {
			inDegree[child]--
			if inDegree[child] == 0 {
				sources = append(sources, child)
			}
		}
	}

	if len(sortedOrder) != vertices {
		return []int{}
	}
	return sortedOrder
}
