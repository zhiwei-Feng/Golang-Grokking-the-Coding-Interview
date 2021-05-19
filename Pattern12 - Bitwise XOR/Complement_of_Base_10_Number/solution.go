package complementofbase10number

/*
Every non-negative integer N has a binary representation, for example, 8 can be represented as “1000” in binary and 7 as “0111” in binary.

The complement of a binary representation is the number in binary that we get when we change every 1 to a 0 and every 0 to a 1.
For example, the binary complement of “1010” is “0101”.

For a given positive number N in base-10, return the complement of its binary representation as a base-10 integer.

Input: 8
Output: 7
Explanation: 8 is 1000 in binary, its complement is 0111 in binary, which is 7 in base-10.

Input: 10
Output: 5
Explanation: 10 is 1010 in binary, its complement is 0101 in binary, which is 5 in base-10.

ref: https://leetcode-cn.com/problems/complement-of-base-10-integer/
*/

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	var (
		all_bit_one = 0
		root        = 1
		tmpN        = n
	)
	for tmpN > 0 {
		all_bit_one = all_bit_one + root
		root = root * 2
		tmpN = tmpN >> 1
	}

	return all_bit_one ^ n
}
