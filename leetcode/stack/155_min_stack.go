package main

/*
  解题思路：最小栈问题，也是很常见。最关键就是如果实现O(1)操作获取最小值。其实需要两个list一个做栈存储，一个做最小值存储。
  时间复杂度: O(1), 空间复杂度O(n)
 */

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else if x > this.GetMin() {
		this.min = append(this.min, this.GetMin())
	} else {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[0 : len(this.data)-1]
	this.min = this.min[0 : len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
