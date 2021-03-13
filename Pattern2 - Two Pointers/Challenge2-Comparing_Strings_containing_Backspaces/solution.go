package Challenge2_Comparing_Strings_containing_Backspaces

/*
Given two strings containing backspaces (identified by the character ‘#’), check if the two strings are equal.

Input: str1="xy#z", str2="xzz#"
Output: true
Explanation: After applying backspaces the strings become "xz" and "xz" respectively.

Input: str1="xy#z", str2="xyz#"
Output: false
Explanation: After applying backspaces the strings become "xz" and "xy" respectively.

Input: str1="xp#", str2="xyz##"
Output: true
Explanation: After applying backspaces the strings become "x" and "x" respectively.
In "xyz##", the first '#' removes the character 'z' and the second '#' removes the character 'y'.

Input: str1="xywrrmp", str2="xywrrmu#p"
Output: true
Explanation: After applying backspaces the strings become "xywrrmp" and "xywrrmp" respectively.
*/

//func handleBackspace(str string) string {
//	// version1
//	var newStr = ""
//	for i := 0; i < len(str); i++ {
//		if str[i] != '#' {
//			newStr = newStr + string(str[i])
//		} else {
//			newStr = newStr[:len(newStr)-1] // remove last char
//		}
//	}
//	return newStr
//}

func getNextValidIndex(str string, index int) int {
	backspaceCount := 0
	for index >= 0 {
		if str[index] == '#' {
			backspaceCount++
		} else if backspaceCount > 0 {
			backspaceCount--
		} else {
			break
		}
		index--
	}
	return index
}

func compare(str1, str2 string) bool {
	// version 1
	// 直接将str1和str2通过处理#的方式来获得最终的string
	//finalStr1, finalStr2 := handleBackspace(str1), handleBackspace(str2)
	//return finalStr1 == finalStr2

	// version 2
	// 从str的末尾开始比较，如果是#，则跳过下一个valid字符
	i, j := len(str1)-1, len(str2)-1
	for i >= 0 || j >= 0 {
		nexti := getNextValidIndex(str1, i)
		nextj := getNextValidIndex(str2, j)
		if nexti < 0 && nextj < 0 {
			return true
		}
		if nexti < 0 || nextj < 0 {
			return false
		}
		if str1[nexti] != str2[nextj] {
			return false
		}
		i = nexti - 1
		j = nextj - 1
	}
	// str1和str2开头都是valid字符则可以到这里
	return true
}
