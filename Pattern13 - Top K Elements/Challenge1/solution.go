package challenge1

import (
	"container/heap"
	"strings"
)

/*
* Rearrange String K Distance Apart
*
* Given a string and a number ‘K’, find if the string can be rearranged such that
* the same characters are at least ‘K’ distance apart from each other.

Input: "mmpp", K=2
Output: "mpmp" or "pmpm"
Explanation: All same characters are 2 distance apart.

Input: "Programming", K=3
Output: "rgmPrgmiano" or "gmringmrPoa" or "gmrPagimnor" and a few more
Explanation: All same characters are 3 distance apart.

Input: "aab", K=2
Output: "aba"
Explanation: All same characters are 2 distance apart.

Input: "aappa", K=3
Output: ""
Explanation: We cannot find an arrangement of the string where any two 'a' are 3 distance apart.

ref: https://leetcode-cn.com/problems/rearrange-string-k-distance-apart/
*/

func rearrangeString(s string, k int) string {
	if k <= 1 {
		return s
	}
	var (
		h     Heap
		res   strings.Builder
		mem   = make(map[rune]int)
		cache = make([]Char, 0)
	)

	for _, v := range s {
		mem[v]++
	}

	heap.Init(&h)
	for k, v := range mem {
		heap.Push(&h, Char{k, v})
	}

	for h.Len() > 0 {
		top := heap.Pop(&h).(Char)
		res.WriteRune(top.Val)
		top.Times--

		cache = append(cache, top)

		if len(cache) == k {
			entry := cache[0]
			cache = cache[1:]
			if entry.Times > 0 {
				heap.Push(&h, entry)
			}
		}

	}

	if res.Len() == len(s) {
		return res.String()
	} else {
		return ""
	}
}

type Char struct {
	Val   rune
	Times int
}
type Heap []Char

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].Times > h[j].Times }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Char)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
