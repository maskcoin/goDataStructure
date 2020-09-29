package stack

import "container/list"

type Stack struct {
	List *list.List // 当作栈的结构
}

func NewStack() *Stack {
	return &Stack{List: list.New()}
}

func (stack *Stack) Push(value interface{}) {
	stack.List.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	ret := stack.List.Back()
	if ret != nil {
		stack.List.Remove(ret)
	}
	return ret.Value
}

func (stack *Stack) Top() interface{} {
	return stack.List.Back().Value
}

func (stack *Stack) Len() int {
	return stack.List.Len()
}

func (stack *Stack) Empty() bool {
	return stack.List.Len() == 0
}
