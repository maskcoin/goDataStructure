package ArrayList

type Stack interface {
	Clear()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
	Iterator() Iterator
}

type ListStack struct {
	list  List
	It     Iterator
}

//func NewArrayListStack() Stack {
//	stack := new(ArrayListStack)
//	stack.myArray = NewArrayList()
//	stack.myIt = stack.myArray.Iterator()
//	return stack
//}

//func (stack *ArrayListStack) Clear() {
//	stack.myArray.Clear()
//	stack.myIt = stack.myArray.Iterator()
//}
//
//func (stack *ArrayListStack) Size() int {
//	return stack.myArray.Size()
//}
//
//func (stack *ArrayListStack) Pop() interface{} {
//	if !stack.IsEmpty() {
//		last := stack.myArray.dataStore[stack.myArray.Size() - 1]
//		stack.myArray.Delete(stack.myArray.Size() - 1)
//		return last
//	} else {
//		return nil
//	}
//}
//
//func (stack *ArrayListStack) Push(data interface{}) {
//	if !stack.IsFull() {
//		stack.myArray.Append(data)
//	}
//}
//
//func (stack *ArrayListStack) IsFull() bool {
//	return stack.myArray.Size() == 10
//}
//
//func (stack *ArrayListStack) IsEmpty() bool {
//	return stack.myArray.Size() == 0
//}
//
//func (stack *ArrayListStack) Iterator() Iterator {
//	it := new(ArrayListStackIterator)
//	it.currentIndex = stack.myArray.TheSize -1
//	it.list = stack
//	return it
//}
