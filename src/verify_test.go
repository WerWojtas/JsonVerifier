package main

import (
	"testing"
)

func TestVerify(t *testing.T) {
	path0 := "../input/0.json"
	if !verify(path0) {
		t.Fail()
	}
	path1 := "../input/1.json"
	if verify(path1) {
		t.Fail()
	}
	path2 := "../input/2.json"
	if !verify(path2) {
		t.Fail()
	}
	path3 := "../input/3.json"
	if !verify(path3) {
		t.Fail()
	}
	path4 := "../input/4.json"
	if !verify(path4) {
		t.Fail()
	}
	path5 := "../input/5.json"
	if verify(path5) {
		t.Fail()
	}
	path6 := "../input/6.json"
	if !verify(path6) {
		t.Fail()
	}
	
}
