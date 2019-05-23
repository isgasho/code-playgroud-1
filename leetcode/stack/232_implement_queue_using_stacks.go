package main

/*
  解题思路：栈实现队列，做两次栈操作就是队列，那么一个in栈 一个out栈就能解决。每当要pop时候，如果out有值就直接pop，如果没有就从in倒进来再处理。
  时间复杂度：O(n)
 */

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.stackOut) > 0 {
		item := this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[0 : len(this.stackOut)-1]
		return item
	}
	for len(this.stackIn) > 1 {
		this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
		this.stackIn = this.stackIn[:len(this.stackIn)-1]
	}
	if len(this.stackIn) == 1 {
		item := this.stackIn[0]
		this.stackIn = []int{}
		return item
	}
	return 0
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stackOut) > 0 {
		return this.stackOut[len(this.stackOut)-1]
	}
	if len(this.stackIn) > 0 {
		return this.stackIn[0]
	}
	return 0
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stackIn) == 0 && len(this.stackOut) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
