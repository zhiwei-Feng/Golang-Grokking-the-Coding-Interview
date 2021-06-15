package All_Tasks_Scheduling_Orders

import "fmt"

/*
* All Tasks Scheduling Orders
There are ‘N’ tasks, labeled from ‘0’ to ‘N-1’.
Each task can have some prerequisite tasks which need to be completed before it can be scheduled.
Given the number of tasks and a list of prerequisite pairs, write a method to print all possible ordering of tasks meeting all prerequisites.

Input: Tasks=3, Prerequisites=[0, 1], [1, 2]
Output: [0, 1, 2]
Explanation: There is only possible ordering of the tasks.

Input: Tasks=4, Prerequisites=[3, 2], [3, 0], [2, 0], [2, 1]
Output:
1) [3, 2, 0, 1]
2) [3, 2, 1, 0]
Explanation: There are two possible orderings of the tasks meeting all prerequisites.

Input: Tasks=6, Prerequisites=[2, 5], [0, 5], [0, 4], [1, 4], [3, 2], [1, 3]
Output:
1) [0, 1, 4, 3, 2, 5]
2) [0, 1, 3, 4, 2, 5]
3) [0, 1, 3, 2, 4, 5]
4) [0, 1, 3, 2, 5, 4]
5) [1, 0, 3, 4, 2, 5]
6) [1, 0, 3, 2, 4, 5]
7) [1, 0, 3, 2, 5, 4]
8) [1, 0, 4, 3, 2, 5]
9) [1, 3, 0, 2, 4, 5]
10) [1, 3, 0, 2, 5, 4]
11) [1, 3, 0, 4, 2, 5]
12) [1, 3, 2, 0, 5, 4]
13) [1, 3, 2, 0, 4, 5]
*/

func printOrders(tasks int, prerequisites [][]int) {
	var (
		graph       = make(map[int][]int)
		inDegree    = make(map[int]int)
		queue       = make([]int, 0)
		sortedOrder = make([]int, 0)
	)

	for i := 0; i < tasks; i++ {
		graph[i] = make([]int, 0)
		inDegree[i] = 0
	}

	for _, v := range prerequisites {
		graph[v[0]] = append(graph[v[0]], v[1])
		inDegree[v[1]]++
	}

	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	printAllTopologicalSorts(graph, inDegree, queue, sortedOrder)
}

func printAllTopologicalSorts(graph map[int][]int, inDegree map[int]int, sources []int, sortedOrder []int) {
	if len(sources) != 0 {
		for _, vertex := range sources {
			sortedOrder = append(sortedOrder, vertex)
			var sourcesForNextCall = make([]int, len(sources))
			copy(sourcesForNextCall, sources)
			// only remove the current source, all other sources should remain in the queue for the next call
			eraseInd := 0
			for i, v := range sourcesForNextCall {
				if v == vertex {
					eraseInd = i
					break
				}
			}
			sourcesForNextCall = append(sourcesForNextCall[:eraseInd], sourcesForNextCall[eraseInd+1:]...)

			children := graph[vertex]
			for _, child := range children {
				inDegree[child]--
				if inDegree[child] == 0 {
					// save the new source for the next call
					sourcesForNextCall = append(sourcesForNextCall, child)
				}
			}

			// recursive call to print other orderings from the remaining (and new) sources
			printAllTopologicalSorts(graph, inDegree, sourcesForNextCall, sortedOrder)

			// backtrack, remove the vertex from the sorted order and put all of its
			// children back to consider the next source instead of the current vertex
			for i, v := range sortedOrder {
				if v == vertex {
					eraseInd = i
					break
				}
			}
			sortedOrder = append(sortedOrder[:eraseInd], sortedOrder[eraseInd+1:]...)
			for _, child := range children {
				inDegree[child]++
			}
		}
	}

	// if sortedOrder doesn't contain all tasks, either we've a cyclic dependency between tasks, or
	// we have not processed all the tasks in this recursive call
	if len(sortedOrder) == len(inDegree) {
		for _, v := range sortedOrder {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
