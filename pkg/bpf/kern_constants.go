package bpf

import (
	"fmt"
	"reflect"
	"strconv"
)

func KernConstantsGen(obj interface{}) map[string]interface{} {
	constantsMap := make(map[string]interface{})

	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		field := v.Type().Field(i)
		tag := field.Tag

		defaultValue := tag.Get("default")
		bpfName := tag.Get("bpf")

		value := fmt.Sprintf("%v", v.Field(i))
		if value != defaultValue && bpfName != "" {
			switch field.Type.String() {
			case "bool":
				storeVal, err := strconv.ParseBool(value)
				if err != nil {
					panic(any(err))
				}
				constantsMap[bpfName] = storeVal

			case "uint32":
				storeVal, err := strconv.ParseUint(value, 10, 32)
				if err != nil {
					panic(any(err))
				}
				constantsMap[bpfName] = storeVal

			case "uint64":
				storeVal, err := strconv.ParseUint(value, 10, 64)
				if err != nil {
					panic(any(err))
				}
				constantsMap[bpfName] = storeVal
			}

		}
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}
	}
	return constantsMap
}
