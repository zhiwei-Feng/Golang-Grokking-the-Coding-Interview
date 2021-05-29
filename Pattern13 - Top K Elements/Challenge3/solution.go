package challenge3

import "container/heap"

/*
* Frequency Stack
Design a class that simulates a Stack data structure, implementing the following two operations:
1. push(int num): Pushes the number ‘num’ on the stack.
2. pop(): Returns the most frequent number in the stack. If there is a tie, return the number which was pushed later.

After following push operations: push(1), push(2), push(3), push(2), push(1), push(2), push(5)

1. pop() should return 2, as it is the most frequent number
2. Next pop() should return 1
3. Next pop() should return 2

ref: https://leetcode-cn.com/problems/maximum-frequency-stack/
*/

type FreqStack struct {
	Stack []int
	Heap  FreqHeap
	Map   map[int]*Frequency
}

func Constructor() FreqStack {
	fs := FreqStack{}
	heap.Init(&fs.Heap)
	fs.Map = make(map[int]*Frequency)
	return fs
}

func (this *FreqStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if f, ok := this.Map[val]; ok {
		f.LastInd = len(this.Stack) - 1
		f.Times++
	} else {
		this.Map[val] = &Frequency{val, 1, len(this.Stack) - 1}
	}
	heap.Push(&this.Heap, *this.Map[val])
}

func (this *FreqStack) Pop() int {
	x := heap.Pop(&this.Heap).(Frequency).Val
	this.Map[x].Times--
	if this.Map[x].Times == 0 {
		delete(this.Map, x)
	}
	return x
}

type Frequency struct {
	Val     int
	Times   int
	LastInd int
}
type FreqHeap []Frequency

func (h FreqHeap) Len() int { return len(h) }
func (h FreqHeap) Less(i, j int) bool {
	if h[i].Times != h[j].Times {
		return h[i].Times > h[j].Times
	} else {
		return h[i].LastInd > h[j].LastInd
	}
}
func (h FreqHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *FreqHeap) Push(x interface{}) { *h = append(*h, x.(Frequency)) }
func (h *FreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
