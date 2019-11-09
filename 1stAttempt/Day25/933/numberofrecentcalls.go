package main

type RecentCounter struct {
	t []int

}


func Constructor() RecentCounter {
	return  RecentCounter{
		t:make([]int,0,1000),
	}
}


func (this *RecentCounter) Ping(t int) int {
	this.t=append(this.t,t)
	if len(this.t)==0{
		return 0
	}
	for this.t[0]+3000<t{
		this.t=this.t[1:]

	}
	return len(this.t)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
