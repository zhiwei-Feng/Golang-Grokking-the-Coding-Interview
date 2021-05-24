package frequencysort

import (
	"container/heap"
	"strings"
)

/*
Given a string, sort it based on the decreasing frequency of its characters.

Input: "Programming"
Output: "rrggmmPiano"
Explanation: 'r', 'g', and 'm' appeared twice, so they need to appear before any other character.

Input: "abcbab"
Output: "bbbaac"
Explanation: 'b' appeared three times, 'a' appeared twice, and 'c' appeared only once.

ref: https://leetcode-cn.com/problems/sort-characters-by-frequency/
*/

type Character struct {
	Value rune
	Freq  int
}

type CharHeap []Character

func (ch CharHeap) Len() int            { return len(ch) }
func (ch CharHeap) Less(i, j int) bool  { return ch[i].Freq > ch[j].Freq }
func (ch CharHeap) Swap(i, j int)       { ch[i], ch[j] = ch[j], ch[i] }
func (ch *CharHeap) Push(x interface{}) { *ch = append(*ch, x.(Character)) }
func (ch *CharHeap) Pop() interface{} {
	old := *ch
	n := len(old)
	x := old[n-1]
	*ch = old[:n-1]
	return x
}

func frequencySort(s string) string {
	var (
		res = strings.Builder{}
		ch  CharHeap
		m   = make(map[rune]int)
	)

	for _, v := range s {
		m[v]++
	}
	heap.Init(&ch)
	for key, value := range m {
		heap.Push(&ch, Character{key, value})
	}

	for ch.Len() > 0 {
		top := heap.Pop(&ch).(Character)
		for i := 0; i < top.Freq; i++ {
			res.WriteRune(top.Value)
		}
	}

	return res.String()
}
