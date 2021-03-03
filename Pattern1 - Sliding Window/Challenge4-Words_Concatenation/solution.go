package Challenge4_Words_Concatenation

/*
Given a string and a list of words,
find all the starting indices of substrings in the given string that are a concatenation of all the given words
exactly once without any overlapping of words. It is given that all words are of the same length.

Input: String="catfoxcat", Words=["cat", "fox"]
Output: [0, 3]
Explanation: The two substring containing both the words are "catfox" & "foxcat".

Input: String="catcatfoxfox", Words=["cat", "fox"]
Output: [3]
Explanation: The only substring containing both the words is "catfox".
*/

func findWordConcatenation(str string, words []string) []int {
	var (
		wordFreqMap = make(map[string]int)
		wordsCount  = len(words)
		wordLength  = len(words[0])
		startIndex  = make([]int, 0, 10) // store final return result
	)

	// Keep the frequency of every word in a Map.
	for _, v := range words {
		wordFreqMap[v]++
	}

	// 注意这个不是常规的sliding window pattern
	for i := 0; i <= len(str)-wordLength*wordsCount; i++ {
		wordsSeenMap := make(map[string]int) // key1: 保存看过的word的数量

		// ===begin=== 对从当前index i开始的长度为wordLength*wordsCount的string进行判定
		for j := 0; j < wordsCount; j++ {
			nextWordIndex := i + j*wordLength // 当前word的开始index
			word := str[nextWordIndex : nextWordIndex+wordLength]

			// 不满足条件的情况直接break
			if _, ok := wordFreqMap[word]; !ok {
				// 如果出现不在words数组中的word，则直接跳出
				break
			}
			wordsSeenMap[word]++
			if wordsSeenMap[word] > wordFreqMap[word] {
				// 如果words数组中某个word数量出现多次，也直接跳出
				break
			}

			// 如果遍历到该string的最后一个字符都没跳出循环，意味该string是满足条件的
			if j == wordsCount-1 {
				startIndex = append(startIndex, i)
			}
		}
		// ===end===
	}

	return startIndex
}
