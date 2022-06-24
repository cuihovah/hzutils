package hzutils

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"time"
)


func GenerateDictionary(list interface{}, fn func(interface{})interface{}) map[interface{}]interface{} {
	ret := make(map[interface{}]interface{})
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		key := fn(reflect.ValueOf(list).Index(i).Interface())
		ret[key] = reflect.ValueOf(list).Index(i).Interface()
	}
	return ret
}

func Clone(list interface{}, fn func(interface{})interface{}) []interface{} {
	ret := make([]interface{}, reflect.ValueOf(list).Len())
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		ret[i] = fn(reflect.ValueOf(list).Index(i).Interface())
	}
	return ret
}

func HashMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}

func GroupBy(list interface{}, fn func(interface{})interface{}) map[interface{}][]interface{} {
	dict := make(map[interface{}][]interface{})
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		key := fn(reflect.ValueOf(list).Index(i).Interface())
		if _, exists := dict[key]; exists == false {
			dict[key] = make([]interface{}, 0)
		}
		dict[key] = append(dict[key], reflect.ValueOf(list).Index(i).Interface())
	}

	return dict
}

func OrderedGroupBy(list interface{}, fn func(interface{})interface{}) [][]interface{} {
	dict := make(map[interface{}][]interface{})
	ordered := make([]interface{}, 0)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		key := fn(reflect.ValueOf(list).Index(i).Interface())
		if _, exists := dict[key]; exists == false {
			dict[key] = make([]interface{}, 0)
			ordered = append(ordered, key)
		}
		dict[key] = append(dict[key], reflect.ValueOf(list).Index(i).Interface())
	}
	ret := make([][]interface{}, 0)
	for _, key := range ordered {
		ret = append(ret, dict[key])
	}
	return ret
}

func CountBy(list interface{}, fn func(interface{})interface{}) map[interface{}]int64 {
	dict := make(map[interface{}]int64)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		key := fn(reflect.ValueOf(list).Index(i).Interface())
		if _, exists := dict[key]; exists == false {
			dict[key] = int64(0)
		}
		dict[key] += 1
	}

	return dict
}

func GroupValues(list map[interface{}][]interface{}) []interface{} {
	ret := make([]interface{}, 0)
	for _, v := range list {
		ret = append(ret, v)
	}
	return ret
}

func Values(list map[interface{}]interface{}) []interface{} {
	ret := make([]interface{}, 0)
	for _, v := range list {
		ret = append(ret, v)
	}
	return ret
}

func Keys(list map[interface{}]interface{}) []interface{} {
	ret := make([]interface{}, 0)
	for k, _ := range list {
		ret = append(ret, k)
	}
	return ret
}

func Filter(list interface{}, fn func(interface{})bool) []interface{} {
	ret := make([]interface{}, 0)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		if fn(reflect.ValueOf(list).Index(i).Interface()) {
			ret = append(ret, reflect.ValueOf(list).Index(i).Interface())
		}
	}
	return ret
}

func Range(start, end int64, skeps ...int64) []int64 {
	skep := int64(1)
	if len(skeps) > 0 {
		skep = skeps[0]
	}
	ret := make([]int64, 0)
	for i := start; i <= end; i += skep {
		ret = append(ret, i)
	}
	return ret
}

func Now() int64 {
	return time.Now().Unix()
}

func Reduce(list interface{}, fn func(interface{}, interface{})interface{}, zero interface{}) interface{} {
	last := zero
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		last = fn(last, reflect.ValueOf(list).Index(i).Interface())
	}
	return last
}

func FindIndex(list interface{}, fn func(interface{})bool) int {
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		if fn(reflect.ValueOf(list).Index(i).Interface()) == true {
			return i
		}
	}
	return -1
}

func Contains(list interface{}, fn func(interface{})bool) bool {
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		if fn(reflect.ValueOf(list).Index(i).Interface()) == true {
			return true
		}
	}
	return false
}

