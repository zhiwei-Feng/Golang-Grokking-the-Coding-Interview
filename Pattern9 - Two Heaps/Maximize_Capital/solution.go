package Maximize_Capital

import (
	"container/heap"
)

/*
Given a set of investment projects with their respective profits, we need to find the most profitable projects.
We are given an initial capital and are allowed to invest only in a fixed number of projects. Our goal is to choose projects that give us the maximum profit.
Write a function that returns the maximum total capital after selecting the most profitable projects.

We can start an investment project only when we have the required capital. Once a project is selected, we can assume that its profit has become our capital.
ref: https://leetcode-cn.com/problems/ipo/

Input: Project Capitals=[0,1,2], Project Profits=[1,2,3], Initial Capital=1, Number of Projects=2
Output: 6
Explanation:

1. With initial capital of ‘1’, we will start the second project which will give us profit of ‘2’.
Once we selected our first project, our total capital will become 3 (profit + initial capital).
2. With ‘3’ capital, we will select the third project, which will give us ‘3’ profit.

After the completion of the two projects, our total capital will be 6 (1+2+3).

Input: Project Capitals=[0,1,2,3], Project Profits=[1,2,3,5], Initial Capital=0, Number of Projects=3
Output: 8
Explanation:

1. With ‘0’ capital, we can only select the first project, bringing out capital to 1.
2. Next, we will select the second project, which will bring our capital to 3.
3. Next, we will select the fourth project, giving us a profit of 5.

After selecting the three projects, our total capital will be 8 (1+2+5).
*/

// k: proj nums, w: init capital
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var (
		minCapital = &MinCapitalHeap{}
		maxProfit  = &MaxProfitHeap{}
	)

	heap.Init(minCapital)
	heap.Init(maxProfit)

	for i := 0; i < len(profits); i++ {
		heap.Push(minCapital, Project{profits[i], capital[i]})
	}

	for i := 0; i < k; i++ {
		for minCapital.Len() != 0 && minCapital.Top().capital <= w {
			heap.Push(maxProfit, heap.Pop(minCapital))
		}
		if maxProfit.Len() == 0 {
			break
		}
		w += heap.Pop(maxProfit).(Project).profit
	}

	return w
}

type Project struct {
	profit  int
	capital int
}

type MinCapitalHeap []Project
type MaxProfitHeap []Project

func (m MinCapitalHeap) Len() int {
	return len(m)
}

func (m MinCapitalHeap) Less(i, j int) bool {
	return m[i].capital < m[j].capital
}

func (m MinCapitalHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinCapitalHeap) Push(x interface{}) {
	*m = append(*m, x.(Project))
}

func (m *MinCapitalHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func (m MinCapitalHeap) Top() Project {
	return m[0]
}

func (m MaxProfitHeap) Len() int {
	return len(m)
}

func (m MaxProfitHeap) Less(i, j int) bool {
	return m[i].profit > m[j].profit
}

func (m MaxProfitHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxProfitHeap) Push(x interface{}) {
	*m = append(*m, x.(Project))
}

func (m *MaxProfitHeap) Pop() interface{} {
	old := *m
	n := len(*m)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func (m MaxProfitHeap) Top() Project {
	return m[0]
}
