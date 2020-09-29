package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int //数组大小
	Get(index int) (interface{}, error)
	Set(index int, newVal interface{}) error
	Insert(index int, newVal interface{}) error
	Append(newVal interface{})
	Clear()
	Delete(index int) error
	String() string
	Iterator() Iterator
}

type ArrayList struct {
	dataStore []interface{}
	TheSize   int
}

func NewArrayList() *ArrayList { // 构造一个接口
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.TheSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}

	return list.dataStore[index], nil
}

func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(list.dataStore, newVal)
	list.TheSize++
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}

	list.dataStore[index] = newVal

	return nil
}
func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.checkIsFull()
	list.dataStore = list.dataStore[:list.TheSize+1]
	for i := list.TheSize; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newVal
	list.TheSize++
	return nil
}

func (list *ArrayList) checkIsFull() {
	if list.TheSize == cap(list.dataStore) {
		newDataStore := make([]interface{}, 2*list.TheSize)
		copy(newDataStore, list.dataStore)
		list.dataStore = newDataStore
	}
}
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}
func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}

	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.TheSize--
	return nil
}
