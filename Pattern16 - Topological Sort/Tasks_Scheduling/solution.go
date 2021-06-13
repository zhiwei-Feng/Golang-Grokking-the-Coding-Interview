package Tasks_Scheduling

/*
* Tasks Scheduling
There are ‘N’ tasks, labeled from ‘0’ to ‘N-1’.
Each task can have some prerequisite tasks which need to be completed before it can be scheduled.
Given the number of tasks and a list of prerequisite pairs, find out if it is possible to schedule all the tasks.

Input: Tasks=3, Prerequisites=[0, 1], [1, 2]
Output: true
Explanation: To execute task '1', task '0' needs to finish first. Similarly, task '1' needs to finish
before '2' can be scheduled. A possible sceduling of tasks is: [0, 1, 2]

Input: Tasks=3, Prerequisites=[0, 1], [1, 2], [2, 0]
Output: false
Explanation: The tasks have cyclic dependency, therefore they cannot be sceduled.

Input: Tasks=6, Prerequisites=[2, 5], [0, 5], [0, 4], [1, 4], [3, 2], [1, 3]
Output: true
Explanation: A possible scheduling of tasks is: [0 1 4 3 2 5]
*/

func isSchedulingPossible(tasks int, prerequisites [][]int) bool {
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

	return len(topoRes) == tasks
}
