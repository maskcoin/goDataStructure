package hashtable

import (
	"errors"
	"math"
)

type HashTable struct {
	Table map[int]*List
	size  int
	cap   int
}

type Item struct {
	key   string
	value interface{}
}

func NewHashTable(cap int) *HashTable {
	table := make(map[int]*List, cap) // 初始化

	return &HashTable{
		Table: table,
		size:  0,
		cap:   cap,
	}
}

func (ht *HashTable) Get(key string) (interface{}, error) {
	index := ht.Pos(key)
	item, err := ht.Find(index, key)
	if err != nil {
		return "", err
	} else {
		return item.value, nil
	}
}

func (ht *HashTable) Put(key, value string) {
	index := ht.Pos(key)
	if ht.Table[index] == nil {
		ht.Table[index] = NewList()
	}

	data, err := ht.Find(index, key)
	if err != nil {
		item := &Item{
			key:   key,
			value: value,
		}

		ht.Table[index].Append(item)
		ht.size++
	} else {
		data.value = value
	}
}

func (ht *HashTable) Del(key string) error {
	index := ht.Pos(key)
	data, err := ht.Find(index, key)
	if err != nil {
		return nil
	} else {
		myList := ht.Table[index]
		myList.Remove(data)
		return nil
	}
}

//循环哈希表的多个链表，循环每个链表的元素
func (ht *HashTable) ForEach(f func(item *Item)) {
	for k := range ht.Table {
		if ht.Table[k] != nil {
			ht.Table[k].Each(func(node Node) {
				f(node.Value.(*Item))
			})
		}
	}
}

func (ht *HashTable) Pos(s string) int {
	return HashCode(s) % ht.cap
}

func (ht *HashTable) Find(i int, key string) (*Item, error) {
	myList := ht.Table[i]
	var val *Item
	myList.Each(func(node Node) {
		if node.Value.(*Item).key == key {
			val = node.Value.(*Item) //取出数据
		}
	})
	if val == nil {
		return nil, errors.New("not find")
	}
	return val, nil
}

//根据字符串计算hash
func HashCode(str string) int {
	hash := int32(0)
	for i := 0; i < len(str); i++ {
		hash = int32(hash<<5-hash) + int32(str[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}
