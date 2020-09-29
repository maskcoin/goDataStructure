package ArrayList

import "errors"

type ArrayListStackIterator struct {
	list         *ArrayListStack
	currentIndex int //当前索引
}

func (it *ArrayListStackIterator) HasNext() bool {
	return it.currentIndex < it.list.myArray.Size() && it.currentIndex > 0
}

func (it *ArrayListStackIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个")
	}

	value, err := it.list.myArray.Get(it.currentIndex)
	it.currentIndex--
	return value, err
}

func (it *ArrayListStackIterator) Remove() {
	it.list.myArray.Delete(it.currentIndex)
	it.currentIndex--
}

func (it *ArrayListStackIterator) GetIndex() int {
	return it.currentIndex
}


