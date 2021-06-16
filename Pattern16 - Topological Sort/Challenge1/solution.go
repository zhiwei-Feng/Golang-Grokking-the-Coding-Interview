package Challenge1

/*
Given a sequence originalSeq and an array of sequences, write a method to find if originalSeq can be uniquely reconstructed from the array of sequences.
Unique reconstruction means that we need to find if originalSeq is the only sequence such that all sequences in the array are subsequences of it.

Input: originalSeq: [1, 2, 3, 4], seqs: [[1, 2], [2, 3], [3, 4]]
Output: true
Explanation: The sequences [1, 2], [2, 3], and [3, 4] can uniquely reconstruct
[1, 2, 3, 4], in other words, all the given sequences uniquely define the order of numbers
in the 'originalSeq'.

Input: originalSeq: [1, 2, 3, 4], seqs: [[1, 2], [2, 3], [2, 4]]
Output: false
Explanation: The sequences [1, 2], [2, 3], and [2, 4] cannot uniquely reconstruct
[1, 2, 3, 4]. There are two possible sequences we can construct from the given sequences:
1) [1, 2, 3, 4]
2) [1, 2, 4, 3]

Input: originalSeq: [3, 1, 4, 2, 5], seqs: [[3, 1, 5], [1, 4, 2, 5]]
Output: true
Explanation: The sequences [3, 1, 5] and [1, 4, 2, 5] can uniquely reconstruct
[3, 1, 4, 2, 5].

ref: https://leetcode-cn.com/problems/sequence-reconstruction/
*/

func sequenceReconstruction(org []int, seqs [][]int) bool {
	if len(org) <= 0 {
		return false
	}

	var (
		graph       = make(map[int][]int)
		inDegree    = make(map[int]int)
		queue       = make([]int, 0)
		sortedOrder = make([]int, 0)
	)
	for _, seq := range seqs {
		for i := 0; i < len(seq); i++ {
			graph[seq[i]] = make([]int, 0)
			inDegree[seq[i]] = 0
		}
	}

	for _, seq := range seqs {
		for i := 0; i < len(seq)-1; i++ {
			head, tail := seq[i], seq[i+1]
			graph[head] = append(graph[head], tail)
			inDegree[tail]++
		}
	}

	if len(inDegree) != len(org) {
		return false
	}

	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	if len(queue) >= 2 {
		return false
	}

	for len(queue) != 0 {
		if len(queue) > 1 {
			return false
		}
		if org[len(sortedOrder)] != queue[0] {
			return false
		}
		vertex := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, vertex)

		for _, v := range graph[vertex] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return len(sortedOrder) == len(org)
}
