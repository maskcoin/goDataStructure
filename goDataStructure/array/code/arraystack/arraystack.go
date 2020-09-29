package arraystack

type Stack interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type ArrayStack struct {
	dataSource  []interface{}
	capSize     int
	currentSize int
}

func NewArrayStack() *ArrayStack {
	stack := new(ArrayStack)
	stack.dataSource = make([]interface{}, 0, 10)
	stack.capSize = 10
	stack.currentSize = 0
	return stack
}

func (stack *ArrayStack) Clear() {
	stack.dataSource = make([]interface{}, 0, 10)
	stack.capSize = 10
	stack.currentSize = 0
}

func (stack *ArrayStack) Size() int {
	return stack.currentSize
}

func (stack *ArrayStack) Pop() interface{} {
	if !stack.IsEmpty() {
		lastItem := stack.dataSource[stack.currentSize-1]
		stack.dataSource = stack.dataSource[:stack.currentSize-1]
		stack.currentSize--
		return lastItem
	} else {
		return nil
	}
}

func (stack *ArrayStack) Push(data interface{}) {
	if !stack.IsFull() {
		stack.dataSource = append(stack.dataSource, data)
		stack.currentSize++
	}
}

func (stack *ArrayStack) IsFull() bool {
	return stack.currentSize == stack.capSize
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.currentSize == 0
}
