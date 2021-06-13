package Tasks_Scheduling_Order

/*
* Tasks Scheduling Order
There are ‘N’ tasks, labeled from ‘0’ to ‘N-1’.
Each task can have some prerequisite tasks which need to be completed before it can be scheduled.
Given the number of tasks and a list of prerequisite pairs, write a method to find the ordering of tasks we should pick to finish all tasks.

Input: Tasks=3, Prerequisites=[0, 1], [1, 2]
Output: [0, 1, 2]
Explanation: To execute task '1', task '0' needs to finish first. Similarly, task '1' needs to finish
before '2' can be scheduled. A possible scheduling of tasks is: [0, 1, 2]

Input: Tasks=3, Prerequisites=[0, 1], [1, 2], [2, 0]
Output: []
Explanation: The tasks have cyclic dependency, therefore they cannot be scheduled.

Input: Tasks=6, Prerequisites=[2, 5], [0, 5], [0, 4], [1, 4], [3, 2], [1, 3]
Output: [0 1 4 3 2 5]
Explanation: A possible scheduling of tasks is: [0 1 4 3 2 5]
*/

func findOrder(tasks int, prerequisites [][]int) []int {
	var (
		inDegree = make(map[int]int)
		graph    = make([][]int, tasks)
		topoRes  = make([]int, 0)
		queue    = make([]int, 0)
	)

	for i := 0; i < tasks; i++ {
		inDegree[i] = 0
		graph[i] = make([]int, 0)
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

	for len(queue) != 0 {
		vertex := queue[0]
		queue = queue[1:]
		topoRes = append(topoRes, vertex)
		for _, v := range graph[vertex] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(topoRes) == tasks {
		return topoRes
	}
	return []int{}
}
