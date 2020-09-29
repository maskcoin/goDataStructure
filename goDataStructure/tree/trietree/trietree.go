package trietree

const MaxSize = 26

type Node struct {
	Char     rune
	Leaf     bool
	Children []*Node
}

func NewNode(char rune) *Node {
	node := &Node{
		Char:     char,
		Children: make([]*Node, MaxSize),
	}

	return node
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{Root: NewNode(0)}
}

//向这颗树的某个节点上插入char对应的节点
func (t *Trie) Insert(node *Node, str string) { //单词全部转成小写
	for _, v := range str {
		loc := v - 'a' //转化成26个字母对应的数组下标
		if node.Children[loc] == nil {
			node.Children[loc] = NewNode(v)
		}
		node = node.Children[loc]
	}
	node.Leaf = true
}

func (t *Trie) Find(str string) (ret bool) {
	curNode := t.Root
	for _, v := range str {
		loc := v - 'a' //转化成26个字母对应的数组下标
		if curNode.Children[loc] == nil {
			return
		}
		curNode = curNode.Children[loc]
	}

	if curNode.Leaf {
		ret = true
	}

	return
}
