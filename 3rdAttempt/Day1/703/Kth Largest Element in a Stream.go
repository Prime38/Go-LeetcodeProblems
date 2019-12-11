package main

type IntHeap []int
func (h IntHeap) Len()int{
	return len(h)
}
func (h IntHeap) Less(i,j int) bool{
	return h[i]<h[j]
}
func (h IntHeap) Swap(i,j int){
	h[i],h[j]=h[j],h[i]
}
func (h *IntHeap) Push(x interface{}){
	*h=append(*h,x.(int))
}
func (h *IntHeap) Pop()interface{}{
	old:=*h
	n:=len(old)
	x:=old[n-1]
	*h=old[0:n-1]
	return x
}
type KthLargest struct {
	Nums *IntHeap
	k int
}
func Constructor(k int, nums []int) KthLargest {
	h:=&IntHeap{}
	heap.Init(h)
	for i:=0;i<len(nums);i++{
		heap.Push(h,nums[i])
	}
	for len(*h)>k{
		heap.Pop(h)
	}
	return KthLargest{h,k}
}
func (this *KthLargest) Add(val int) int {
	if len(*this.Nums)<this.k{
		heap.Push(this.Nums,val)
	} else if val> (*this.Nums)[0]{
		heap.Push(this.Nums,val)
		heap.Pop(this.Nums)
	}
	return (*this.Nums)[0]
}
/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */