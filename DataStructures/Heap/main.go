package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func (iheap *IntegerHeap) Push(val any) {
	*iheap = append(*iheap, val.(int))
}

func (iheap *IntegerHeap) Pop() any {
	var previous IntegerHeap = *iheap
	n := len(previous)
	val := previous[n-1]
	*iheap = previous[:n-1]
	return val
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{40, 1, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
