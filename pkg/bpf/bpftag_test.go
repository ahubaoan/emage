package bpf

import "testing"

func TestGetDefaultVal(t *testing.T) {
	type BpfObj struct {
		Have  bool   `bpf:"name" default:"false"`
		Age   uint32 `bpf:"age" default:"32"`
		Count uint64 `bpf:"count" default:"1234567"`
	}

	test := BpfObj{
		Have: true,
	}

	if GetDefaultVal(test, "Have").(bool) != false {
		t.Fatal("test Have err")
	}
	if GetDefaultVal(test, "Age").(uint32) != uint32(32) {
		t.Fatal("test Age err, val=", GetDefaultVal(test, "Age"))
	}
	if GetDefaultVal(test, "Count").(uint64) != uint64(1234567) {
		t.Fatal("test Count err, val=", GetDefaultVal(test, "Count"))
	}
}

func TestGetBpfCString(t *testing.T) {
	type BpfObj struct {
		Age   uint32 `bpf:"age" default:"32"`
		Count uint64 `bpf:"count" default:"1234567"`
	}
	test := BpfObj{
		Age: 33,
	}
	str, res := GetBpfCString(test, "Age")
	if res == false {
		t.Fatal("test GetBpfCString failed by Name")
	}
	if str != "age" {
		t.Fatal("test GetBpfCString failed by age")
	}
	str, res = GetBpfCString(test, "Count")
	if res == false {
		t.Fatal("test GetBpfCString failed by Count")
	}
	if str != "count" {
		t.Fatal("test GetBpfCString failed by Count")
	}
}
