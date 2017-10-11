package syncmap_sample

import (
	"testing"
)

func TestStoreAndLoad(t *testing.T) {
	testMap := NewMap()
	testMap.Store("test_key", "test_value")
	// keyが登録済みの場合のテスト
	v, ok := testMap.Load("test_key")
	if v != "test_value" || ok != true {
		t.Error("TestStoreAndLoad faild.  v:", v, " ok:", ok)
	}
	// keyがない場合のテスト
	v, ok = testMap.Load("test_key1")
	if v != nil || ok != false {
		t.Error("TestStoreAndLoad faild.  v:", v, " ok:", ok)
	}
}

func TestLoadOrStore(t *testing.T) {
	testMap := NewMap()
	// 登録がない場合のテスト
	actual, loaded := testMap.LoadOrStore("test_key", "test_value")
	if actual != "test_value" || loaded != false {
		t.Error("TestLoadOrStore actual:", actual, " loaded:", loaded)
	}
	// 登録がある場合のテスト（上書きされない)
	actual, loaded = testMap.LoadOrStore("test_key", "test_value2")
	if actual != "test_value" || loaded != true {
		t.Error("TestLoadOrStore actual:", actual, " loaded:", loaded)
	}
}

func TestRange(t *testing.T) {
	testMap := NewMap()
	// testとtest1だけ出力される(go test -vで実行)
	testMap.Store("test_key", "test_value")
	testMap.Store("test", "test")
	testMap.Store("test1", "test1")
	testMap.Range()
}

func TestDelete(t *testing.T) {
	testMap := NewMap()
	testMap.Store("test_key", "test_value")
	_, ok := testMap.Load("test_key")
	if ok != true {
		t.Error("TestDelete Load fail")
	}
	testMap.Delete("test_key")
	_, ok = testMap.Load("test_key")
	if ok != false {
		t.Error("TestDelete Delete fail")
	}
}
