package main

import (
	"container/heap"
	"fmt"
)
type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func lastStoneWeight(stones []int) int {
	h:=IntHeap(stones)
	heap.Init(&h)
	for len(h)>=2{
		a,b:=heap.Pop(&h).(int),heap.Pop(&h).(int)
		if a==b{
			continue
		}
		heap.Push(&h,a-b)
	}
	if len(h)==0{
		return 0
	}
	return stones[0]
}
func main(){
	stones:=[]int{2,7,4,1,8,1}
	fmt.Println(lastStoneWeight(stones))
}