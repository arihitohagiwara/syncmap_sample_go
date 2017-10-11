package syncmap_sample

import (
	"fmt"
	"sync"
)

type syncMap struct {
	mapData *sync.Map
}

func NewMap() *syncMap {
	return &syncMap{mapData: new(sync.Map)}
}

// StoreMap mapにデータをストアする
func (thisMap *syncMap) Store(key, value interface{}) {
	thisMap.mapData.Store(key, value)
}

// LoadMap mapからデータを取得する
func (thisMap *syncMap) Load(key interface{}) (value interface{}, ok bool) {
	return thisMap.mapData.Load(key)
}

// LoadOrStore mapにデータがあった場合は取得、ない場合はストアする。
func (thisMap *syncMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	return thisMap.mapData.LoadOrStore(key, value)
}

// Range keyとvalueが同じものだけdumpしてみる
func (thisMap *syncMap) Range() {
	thisMap.mapData.Range(func(key, value interface{}) bool {
		if key == value {
			fmt.Println("key:", key, " value:", value)
		}
		return true
	})
}

// Delete mapにあるデータを削除
func (thisMap *syncMap) Delete(key interface{}) {
	thisMap.mapData.Delete(key)
}
