package set

type Set struct {
	list *List
}

type SetIterator struct {
	index uint64 // 索引
	set   *Set
}

func (set *Set) GetAt(index uint64) Object {
	return (*set).list.GetAt(index)
}

func (set *Set) GetSize() uint64 {
	return (*set).list.GetSize()
}

func (set *Set) Init(match ...MatchFun) {
	myList := new(List)
	(*set).list = myList // 初始化
	if len(match) == 0 {
		myList.Init()
	} else {
		myList.Init(match[0])
	}
}

func (set *Set) IsMember(data Object) bool {
	return (*set).list.IsMember(data)
}

func (set *Set) Insert(data Object) bool {
	if !set.IsMember(data) {
		return (*set).list.Append(data)
	}
	return false
}

func (set *Set) IsEmpty() bool {
	return (*set).list.IsEmpty()
}



func (set *Set) Remove(data Object) bool {
	return (*set).list.Remove(data)
}