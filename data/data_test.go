/*
@File: data_test.go.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2019-06-18 14:51  Lucien     1.0         Init
*/
package data

import (
	"encoding/json"
	"testing"
)

var key string

func TestInsert(t *testing.T) {
	var err error
	key, err = Insert(Paste{"", "plain", "<h1>Hello World!</h1>", ""})
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuery(t *testing.T) {
	t.Log(key)
	object, err := Query(key)
	if err != nil {
		t.Fatal(err)
	}
	content, err := json.Marshal(object)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(content))
}

func TestErase(t *testing.T) {
	err := Delete(key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExist(t *testing.T) {
	err := Delete("hello")
	key, err := Insert(Paste{Key: "hello"})
	if err != nil {
		t.Fatal(err)
	}
	if Exist(key) {
		t.Logf("key: %s exist.", key)
	} else {
		t.Fatalf("key: %s not exist.", key)
	}
}
