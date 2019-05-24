package main

/*
  解题思路：利用栈的思维，遍历字符串，正向符号[{(，遇到正向符号，就把它的反向符号]})压栈。如果遍历到反向符号，然后推出栈顶对比。如果相等即该出回文。
  时间复杂度: O(n) 空间复杂度O(n)
 */

type Stack struct {
	data []rune
}

func newStack() Stack {
	return Stack{}
}

func (this *Stack) pop() rune {
	item := this.data[len(this.data)-1]
	this.data = this.data[0 : len(this.data)-1]
	return item
}

func (this *Stack) push(x rune) {
	this.data = append(this.data, x)
}

func (this *Stack) isEmpty() bool {
	if len(this.data) < 1 {
		return true
	}
	return false
}

func isValid(s string) bool {
	stack := newStack()
	for _, c := range s {
		if c == '(' {
			stack.push(')')
		} else if c == '{' {
			stack.push('}')
		} else if c == '[' {
			stack.push(']')
		} else {
			if stack.isEmpty() || stack.pop() != c {
				return false
			}
		}
	}
	return stack.isEmpty() == true
}
