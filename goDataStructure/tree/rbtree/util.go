package rbtree

type Int int

func (a Int) Less(b Item) bool {
	return a < b.(Int)
}

type UInt32 uint32

func (a UInt32) Less(b Item) bool {
	return a < b.(UInt32)
}

type String string

func (a String) Less(b Item) bool {
	return a < b.(String)
}
