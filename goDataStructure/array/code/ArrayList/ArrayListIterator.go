package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
	Remove()                    // 删除
	GetIndex() int              // 得到索引
}

type Iterable interface {
	Iterator() Iterator // 构造一个接口
}

type ArrayListIterator struct {
	list         *ArrayList
	currentIndex int //当前索引
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.TheSize
}

func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("没有下一个")
	}

	value, err := it.list.Get(it.currentIndex)
	it.currentIndex++
	return value, err
}

func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex)
}

func (it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.currentIndex = 0
	it.list = list
	return it
}
