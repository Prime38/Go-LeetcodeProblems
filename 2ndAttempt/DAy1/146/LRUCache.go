package main
type node struct {
	key int
	val int
	next *node
	prev *node
}
type LRUCache struct {
	capacity int
	head *node
	tail *node
	cache map[int]*node

}
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node),
	}
}

func (this *LRUCache) Get(key int) int {
	value,ok:=this.cache[key]
	if !ok{
		return -1
	}
	this.Pop(value)
	this.Push(value)
	return value.val

}
func (this * LRUCache) Push(node *node){
	node.prev=nil
	node.next=this.head
	if node.next!=nil{
		node.next.prev=node
	}
	this.head=node
	if this.tail==nil{
		this.tail=node
	}

}

func (this * LRUCache) Pop(node *node){
	if node==this.head{
		this.head=node.next
	}
	if node==this.tail{
		this.tail=node.prev
	}
	if node.next!=nil{
		node.next.prev=node.prev
	}
	if node.prev!=nil{
		node.prev.next=node.next
	}

}

func (this *LRUCache) Put(key int, value int)  {
	man,ok:=this.cache[key]
	if !ok{
		man=&node{val:value,key:key}
		this.cache[key]=man
	} else{
		man.val=value
		this.Pop(man)
	}
	this.Push(man)
	if len(this.cache)>this.capacity{
		man=this.tail
		if man!=nil{
			this.Pop(man)
			delete(this.cache,man.key)
		}
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main(){

}
