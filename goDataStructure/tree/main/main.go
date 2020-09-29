package main

//1+(2-1*3)+(1+2)*(1+(1+(2-1))
//栈，二叉树
import (
	"datastructure/link/queue"
	stack2 "datastructure/link/stack"
	"encoding/hex"
	"fmt"
	"tree/avltree"
	"tree/binarytree"
	"tree/bplustree"
	"tree/merkletree"
	"tree/rbtree"
	"tree/segmenttree"
	"tree/splaytree"
	"tree/stack"
	"tree/trietree"
)



func main11()  {
	t := splaytree.NewTree()
	t.Insert(2, "hi")
	t.Insert(10, "hi")
	t.Insert(6, "hi")
	t.Insert(7, "hi")
	t.Insert(8, "hi")
	t.Insert(9, "hi")
	t.Insert(5, "hi")
	t.Insert(1, "hi")
	t.Insert(4, "hi")
	t.Insert(3, "hi")
	t.PrintTree()
	//fmt.Println(t.Root)
	//for i := 0; i < 36; i++ {
	//	_, err := t.Insert(i+1, "hello bitcoin")
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		t.PrintTree()
	//	}
	//}

	t.SearchNode(9)
	t.PrintTree()

	//for i := 0; i < 36; i++ {
	//	err := t.Delete(i)
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		t.PrintTree()
	//	}
	//}
}

func main10() {
	data1, _ := hex.DecodeString("4b901a4ae4ae99df360492166f2ba267cc7f465f01234dc4133adf9244b64fe8")
	data2, _ := hex.DecodeString("45e68b8246700358461f2aff605c09f2180e9871c76cfdbc103b68bd74905340")
	data3, _ := hex.DecodeString("a953d73a18d8d812eaca29867d08e218d571ec6cf7ea1b9ff3d15f2fc23194e0")
	data4, _ := hex.DecodeString("29d928baa4b9f2ac52da1532960274ead11d9f46a5eaee454775abcfcb80cd8d")
	data5, _ := hex.DecodeString("b7824100064ebc2f9cafecd1a11e1f74b9f261ca154ea128756e545d940b265c")
	//data6,_ := hex.DecodeString("308b46a59e0b2dcaf304a681e42bba8b60b8587efa962f440d67c0d29da8a113")

	//merkletree.ReverseBytes(data2)
	datas := [][]byte{data1, data2, data3, data4, data5}
	t := merkletree.NewTree(datas)
	//fmt.Printf("%x\n", t.Root.Hash)
	t.Root.PrintSelf()
}

func main9() {
	nums := []segmenttree.Merger{segmenttree.Int(-2), segmenttree.Int(0), segmenttree.Int(3),
		segmenttree.Int(-5), segmenttree.Int(2), segmenttree.Int(-1)}
	segTree := segmenttree.NewSegmentTree(nums)
	segTree.BuildSegmentTree(0, 0, len(nums)-1)
	//fmt.Println(segTree.Tree)
	fmt.Println(segTree.Query(0, 2))
	fmt.Println(segTree.Query(2, 5))
	fmt.Println(segTree.Query(0, 5))
}

func main8() {
	strs := []string{"java", "ps", "php", "ui", "css", "js"} //右边相当于一个匿名对象，js中把{}左边的类型也省略了
	trie := trietree.NewTrie()
	for _, str := range strs {
		trie.Insert(trie.Root, str)
	}
	fmt.Println("插入完成")
	fmt.Println(trie.Find("java"))
	fmt.Println(trie.Find("jav"))
}

func main7() {
	bt := bplustree.NewBTree(2)
	bt.Insert(1)
	bt.Insert(3)
	bt.Insert(5)
	bt.Insert(7)
	bt.Insert(2)
	bt.Insert(4)
	bt.Insert(8)
	//
	//bt.Delete(3)
	//bt.Delete(1)
	//bt.Delete(8)
	//bt.Delete(2)
	////
	//bt.Delete(5) //此处出现bug
	//bt.Delete(7)
	//bt.Delete(5)

	//nd, inx := bt.FindLeftChildMinNodeAndIndex(bt.Root, 0)
	//fmt.Println(nd.Keys, "inx =", inx)
	//bt.Delete(4)

	//bt.PrintTree(bt.Root)
	//fmt.Println(bt)
	//fmt.Println(bt.Root.Keys[0])
	if bt.Root == nil {
		fmt.Println("空树")
	} else {
		fmt.Println(bt)
		//bt.PrintTree(bt.Root)
	}
}

func main6() {
	t := rbtree.NewRBTree()
	for i := 0; i < 1000000; i++ {
		node := t.NewNode(rbtree.Int(i))
		t.Insert(node)
	}

	for i := 0; i < 900000; i++ {
		t.Delete(rbtree.Int(i))
	}

	fmt.Println(t.Depth(t.Root))
}

func _compare(a, b interface{}) int {
	var newA, newB int
	var ok bool
	if newA, ok = a.(int); !ok {
		return -2
	}
	if newB, ok = b.(int); !ok {
		return -2
	}
	if newA > newB {
		return 1
	} else if newA < newB {
		return -1
	} else {
		return 0
	}
}

func main5() {
	avl, _ := avltree.NewAVLTree(3, _compare)
	avl = avl.Insert(2)
	avl = avl.Insert(1)
	avl = avl.Insert(4)
	avl = avl.Insert(5)
	avl = avl.Insert(6)
	//avl = avl.Insert(7)
	avl = avl.Insert(15)
	avl = avl.Insert(26)
	avl = avl.Insert(17)
	avl = avl.Insert(11)
	avl = avl.Insert(13)
	//avl = avl.Delete(7)
	fmt.Println(avl.GetAll())
}

func main4() {
	stack := stack.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	for !stack.Empty() {
		fmt.Println(stack.Pop())
	}
}

func main3() {
	t := binarytree.NewBinaryTree()
	node0 := binarytree.NewNode(4)
	node1 := binarytree.NewNode(2)
	node2 := binarytree.NewNode(6)
	node3 := binarytree.NewNode(1)
	node4 := binarytree.NewNode(3)
	node5 := binarytree.NewNode(5)
	node6 := binarytree.NewNode(7)
	node7 := binarytree.NewNode(8)
	node0.Left = node1
	node0.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	node6.Right = node7

	t.Root = node0

	//fmt.Println(t.FindLowerAncestor(node1, node5))
	fmt.Println(t.NumNode())
}

func main2() {
	t := binarytree.NewBinaryTree()
	node0 := binarytree.NewNode(4)
	node1 := binarytree.NewNode(2)
	node2 := binarytree.NewNode(6)
	node3 := binarytree.NewNode(1)
	node4 := binarytree.NewNode(3)
	node5 := binarytree.NewNode(5)
	node6 := binarytree.NewNode(7)
	node0.Left = node1
	node0.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6
	t.Root = node0

	fmt.Println("先序遍历：")
	t.PreOrder()
	fmt.Println()

	fmt.Println("PreOrderWithStack")
	stack := stack2.NewStack()
	t.PreOrderWithStack(stack)
	fmt.Println()
	fmt.Println("t.LevelShow(q)")
	q := queue.NewQueue()
	t.LevelShow(q)
}
