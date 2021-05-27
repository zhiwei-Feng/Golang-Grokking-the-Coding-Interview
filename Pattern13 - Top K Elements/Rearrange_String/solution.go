package rearrangestring

import (
	"container/heap"
	"strings"
)

/*
*Given a string, find if its letters can be rearranged in such a way that no two same characters come next to each other.

Input: "aappp"
Output: "papap"
Explanation: In "papap", none of the repeating characters come next to each other.

Input: "Programming"
Output: "rgmrgmPiano" or "gmringmrPoa" or "gmrPagimnor", etc.
Explanation: None of the repeating characters come next to each other.

Input: "aapa"
Output: ""
Explanation: In all arrangements of "aapa", atleast two 'a' will come together e.g., "apaa", "paaa".

ref: https://leetcode-cn.com/problems/reorganize-string
*/

func reorganizeString(s string) string {
	var (
		h          Heap
		rearranged strings.Builder
		mem        = make(map[rune]int)
	)

	heap.Init(&h)
	for _, v := range s {
		mem[v]++
	}

	for k, v := range mem {
		heap.Push(&h, Character{k, v})
	}

	prev := Character{-1, -1}
	for h.Len() > 0 {
		top := heap.Pop(&h).(Character)

		if prev.Times > 0 {
			heap.Push(&h, prev)
		}

		rearranged.WriteRune(top.Val)
		top.Times--
		prev = top
	}

	if rearranged.Len() == len(s) {
		return rearranged.String()
	} else {
		return ""
	}
}

type Character struct {
	Val   rune
	Times int
}

type Heap []Character

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].Times > h[j].Times }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Character)) }
func (h *Heap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
