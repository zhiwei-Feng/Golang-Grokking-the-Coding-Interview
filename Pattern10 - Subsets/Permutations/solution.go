package Permutations

/*
Given a set of distinct numbers, find all of its permutations.

Permutation is defined as the re-arranging of the elements of the set. For example, {1, 2, 3} has the following six permutations:

{1, 2, 3}
{1, 3, 2}
{2, 1, 3}
{2, 3, 1}
{3, 1, 2}
{3, 2, 1}
If a set has ‘n’ distinct elements it will have n!n! permutations.

Input: [1,3,5]
Output: [1,3,5], [1,5,3], [3,1,5], [3,5,1], [5,1,3], [5,3,1]

ref: https://leetcode-cn.com/problems/permutations/
*/

func permute(nums []int) [][]int {
	// DFS version
	var ans = make([][]int, 0)
	var path = make(map[int]int)
	//dfs
	var dfs func(int, []int)
	dfs = func(num int, curPath []int) {
		curPath = append(curPath, num)
		if len(curPath) == len(nums) {
			ans = append(ans, curPath)
			return
		}
		path[num] = 1
		for _, v := range nums {
			if _, ok := path[v]; !ok {
				dfs(v, curPath)
			}
		}
		// 回溯
		delete(path, num)
	}

	for _, v := range nums {
		dfs(v, []int{})
	}

	return ans
}

//func permuteBFSVersion(nums []int) [][]int {
//    // BFS version
//    // 类似于课程提供的算法，即每次对当前的排列的各个位置插入当前数
//    // 直到每个排列完成
//    var ans = make([][]int, 0)
//    var permutations = make([][]int, 0)
//    permutations = append(permutations, []int{})
//    for _, num := range nums {
//        n := len(permutations)
//        for i := 0; i < n; i++ {
//            oldPermutations := permutations[0]
//            permutations = permutations[1:]
//            for j := 0; j <= len(oldPermutations); j++ {
//                newPermutations := make([]int, len(oldPermutations))
//                copy(newPermutations, oldPermutations)
//                newPermutations = insertToSlice(newPermutations, j, num)
//                if len(newPermutations) == len(nums) {
//                    ans = append(ans, newPermutations)
//                } else {
//                    permutations = append(permutations, newPermutations)
//                }
//            }
//        }
//    }
//
//    return ans
//}
//
//func insertToSlice(slice []int, ind, num int) []int {
//    if ind == 0 {
//        return append([]int{num}, slice...)
//    } else if ind == len(slice) {
//        return append(slice, num)
//    }
//
//    slice = append(slice, 0)
//    copy(slice[ind+1:], slice[ind:])
//    slice[ind] = num
//    return slice
//}
