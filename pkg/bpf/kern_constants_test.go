package bpf

import (
	"testing"
)

func TestKernConstantsGen(t *testing.T) {

	type BpfObj struct {
		Have  bool   `bpf:"have" default:"false"`
		Age   uint32 `bpf:"age" default:"32"`
		Count uint64 `bpf:"count" default:"999888"`
	}

	test := BpfObj{
		Have:  true,
		Age:   32,
		Count: 19998881111,
	}
	err, constantsMap := KernConstantsGen(test)
	if err != nil {
		t.Fatal("KernConstantsGen err")
	}
	if val, ok := constantsMap["have"]; !ok {
		t.Fatal("have err ")
	} else {
		if val.(bool) != true {
			t.Fatal("have err val")
		}
	}
	if val, ok := constantsMap["count"]; !ok {
		t.Fatal("count err ")
	} else {
		if val.(uint64) != 19998881111 {
			t.Fatal("count err val")
		}
	}
}
