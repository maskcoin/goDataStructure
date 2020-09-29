package hashmap

import (
	"errors"
	"hash/crc32"
	"sort"
	"sync"
)

//自定义索引容器的类型
type RIndexArr []uint32 //hash环索引

type Ring struct {
	RMap         map[uint32]string
	RIndexArr    RIndexArr // 索引
	sync.RWMutex           //线程安全
}

func (indexArr RIndexArr) Less(i, j int) bool {
	return indexArr[i] < indexArr[j]
}

func (indexArr RIndexArr) Len() int {
	return len(indexArr)
}

func (indexArr RIndexArr) Swap(i, j int) {
	indexArr[i], indexArr[j] = indexArr[j], indexArr[i]
}

func (ring Ring) AddNode(nodeName string) {
	ring.Lock()
	defer ring.Unlock()
	index := crc32.ChecksumIEEE([]byte(nodeName))
	if _, ok := ring.RMap[index]; ok {
		return
	}
	ring.RMap[index] = nodeName
	ring.RIndexArr = append(ring.RIndexArr, index)
	sort.Sort(ring.RIndexArr)
}

func (ring Ring) RemoveNode(nodeName string) {
	ring.Lock()
	defer ring.Unlock()
	index := crc32.ChecksumIEEE([]byte(nodeName))
	if _, ok := ring.RMap[index]; ok {
		return
	}
	delete(ring.RMap, index)
	ring.RIndexArr = RIndexArr{}
	for k := range ring.RMap {
		ring.RIndexArr = append(ring.RIndexArr, k)
	}
	sort.Sort(ring.RIndexArr)
}

func (ring Ring) GetNode(nodeName string) (string, error) {
	ring.RLock()
	defer ring.RUnlock()
	index := crc32.ChecksumIEEE([]byte(nodeName))
	//i := sort.Search(len([]uint32(ring.rIndexArr)), func(i int) bool {
	//	return ring.rIndexArr[i] >= index
	//})
	ret, ok := ring.RMap[index]
	if ok {
		return ret, nil
	} else {
		return "", errors.New("no this key")
	}
}
