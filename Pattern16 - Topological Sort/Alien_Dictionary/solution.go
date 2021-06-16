package Alien_Dictionary

import (
	"strings"
)

/*
* Alien Dictionary
There is a dictionary containing words from an alien language for which we don’t know the ordering of the alphabets.
Write a method to find the correct order of the alphabets in the alien language.
It is given that the input is a valid dictionary and there exists an ordering among its alphabets.

Input: Words: ["ba", "bc", "ac", "cab"]
Output: bac
Explanation: Given that the words are sorted lexicographically by the rules of the alien language, so
from the given words we can conclude the following ordering among its characters:

1. From "ba" and "bc", we can conclude that 'a' comes before 'c'.
2. From "bc" and "ac", we can conclude that 'b' comes before 'a'

From the above two points, we can conclude that the correct character order is: "bac"

Input: Words: ["cab", "aaa", "aab"]
Output: cab
Explanation: From the given words we can conclude the following ordering among its characters:

1. From "cab" and "aaa", we can conclude that 'c' comes before 'a'.
2. From "aaa" and "aab", we can conclude that 'a' comes before 'b'

From the above two points, we can conclude that the correct character order is: "cab"

Input: Words: ["ywx", "wz", "xww", "xz", "zyy", "zwz"]
Output: ywxz
Explanation: From the given words we can conclude the following ordering among its characters:

1. From "ywx" and "wz", we can conclude that 'y' comes before 'w'.
2. From "wz" and "xww", we can conclude that 'w' comes before 'x'.
3. From "xww" and "xz", we can conclude that 'w' comes before 'z'
4. From "xz" and "zyy", we can conclude that 'x' comes before 'z'
5. From "zyy" and "zwz", we can conclude that 'y' comes before 'w'

From the above five points, we can conclude that the correct character order is: "ywxz"
*/

func findOrder(words []string) string {
	var (
		graph    = make(map[rune][]rune)
		inDegree = make(map[rune]int)
		queue    = make([]rune, 0)
		res      = strings.Builder{}
	)

	n := len(words)

	for _, s := range words {
		for _, c := range s {
			graph[c] = make([]rune, 0)
			inDegree[c] = 0
		}
	}

	for i := 0; i < n-1; i++ {
		w1, w2 := words[i], words[i+1]
		for j := 0; j < min(len(w1), len(w2)); j++ {
			var parent, child = rune(w1[j]), rune(w2[j])
			if parent != child {
				graph[parent] = append(graph[parent], child)
				inDegree[child]++
				break
			}
		}
	}

	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) != 0 {
		vertex := queue[0]
		queue = queue[1:]
		res.WriteRune(vertex)
		for _, v := range graph[vertex] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return res.String()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
