package safemap

import (
	"sync"
)

type SyncMap struct {
	MyMap map[string]string
	Lock  *sync.RWMutex
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		MyMap: make(map[string]string),
		Lock:  &sync.RWMutex{},
	}
}


