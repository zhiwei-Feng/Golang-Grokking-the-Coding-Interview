package introduction

/*
Given an array of n-1n−1 integers in the range from 11 to nn, find the one number that is missing from the array.

Input: 1, 5, 2, 6, 4
Answer: 3

Following are some important properties of XOR to remember:

Taking XOR of a number with itself returns 0, e.g.,

1 ^ 1 = 0
29 ^ 29 = 0
Taking XOR of a number with 0 returns the same number, e.g.,

1 ^ 0 = 1
31 ^ 0 = 31
XOR is Associative & Commutative, which means:

(a ^ b) ^ c = a ^ (b ^ c)
a ^ b = b ^ a

ref: https://leetcode-cn.com/problems/missing-number
*/

func missingNumber(nums []int) int {
	n := len(nums) + 1
	x1 := 1
	for i := 2; i <= n; i++ {
		x1 ^= i
	}
	x2 := nums[0]
	for i := 1; i < len(nums); i++ {
		x2 ^= nums[i]
	}
	return x1 ^ x2
}
