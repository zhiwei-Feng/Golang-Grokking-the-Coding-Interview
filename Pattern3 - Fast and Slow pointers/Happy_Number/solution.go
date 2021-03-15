package Happy_Number

/*
Any number will be called a happy number if, after repeatedly replacing it with a number equal to the sum of the square of all of its digits, leads us to number ‘1’.
All other (not-happy) numbers will never reach ‘1’. Instead, they will be stuck in a cycle of numbers which does not include ‘1’.

Input: 23
Output: true (23 is a happy number)
Explanations: Here are the steps to find out that 23 is a happy number:

Input: 12
Output: false (12 is not a happy number)
Explanations: Here are the steps to find out that 12 is not a happy number:
*/

func squareSum(num int) int {
	sum := 0
	for num > 0 {
		curNum := num % 10
		sum += curNum * curNum
		num /= 10
	}
	return sum
}

func findHappyNumber(num int) bool {
	var (
		fast, slow = num, num
	)

	// 1. 有环，则最终fast==slow
	// 2. 无环，则最终fast==slow==1

	fast = squareSum(squareSum(fast))
	slow = squareSum(slow)
	for fast != slow {
		fast = squareSum(squareSum(fast))
		slow = squareSum(slow)
	}

	if fast == 1 {
		return true
	} else {
		return false
	}
}
