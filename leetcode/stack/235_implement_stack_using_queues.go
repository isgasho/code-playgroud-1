package main

/*
  解题思路：这题跟队列那个差不多，拿另外一个队列做容器，互相倒入，导出队列剩余的最后那个元素就是栈顶。还有节省空间的做法
  时间复杂度：O(n)
 */

type MyStack struct {
	queueA []int
	queueB []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if len(this.queueA) > 0 {
		this.queueA = append(this.queueA, x)
	} else {
		this.queueB = append(this.queueB, x)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var item int
	if len(this.queueA) > 0 {
		for len(this.queueA) > 1 {
			item = this.queueA[0]
			this.queueA = this.queueA[1:len(this.queueA)]
			this.queueB = append(this.queueB, item)
		}
		item = this.queueA[0]
		this.queueA = []int{}
		return item
	}
	if len(this.queueB) > 0 {
		for len(this.queueB) > 1 {
			item = this.queueB[0]
			this.queueB = this.queueB[1:len(this.queueB)]
			this.queueA = append(this.queueA, item)
		}
		item = this.queueB[0]
		this.queueB = []int{}
		return item
	}
	return 0
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.queueA) > 0 {
		return this.queueA[len(this.queueA)-1]
	}
	if len(this.queueB) > 0 {
		return this.queueB[len(this.queueB)-1]
	}
	return 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.queueB) == 0 && len(this.queueA) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
