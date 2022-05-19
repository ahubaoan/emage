package bpf

import (
	"reflect"
	"strconv"
)

//func GetDefaultVal(obj interface{}) interface{} {
//	v := reflect.ValueOf(obj).FieldByName("")
//}

func GetBpfCString(obj interface{}, FiledName string) (string, bool) {
	filed, find := reflect.ValueOf(obj).Type().FieldByName(FiledName)
	if find {
		return filed.Tag.Get("bpf"), true
	}
	return "", false
}

func GetDefaultVal(obj interface{}, FiledName string) interface{} {
	filed, find := reflect.ValueOf(obj).Type().FieldByName(FiledName)
	val := ""
	if find {
		val = filed.Tag.Get("default")
	}
	var err error
	var retVal interface{}

	switch filed.Type.String() {
	case "bool":
		retVal, err = strconv.ParseBool(val)

	case "uint32":
		var valUint64 uint64
		valUint64, err = strconv.ParseUint(val, 10, 32)
		retVal = uint32(valUint64)

	case "uint64":
		retVal, err = strconv.ParseUint(val, 10, 64)

	default:
		panic(any("Unsupported kern config type, do you really need it?"))
	}

	if err != nil {
		panic(any(err.Error()))
	}
	if retVal == nil {
		panic(any("Unsupported type"))
	}
	return retVal
}
