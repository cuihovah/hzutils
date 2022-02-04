package hzutils

import "reflect"

func Uniq(list interface{}, fn func(interface{})interface{}) []interface{} {
	dict := make(map[interface{}]interface{})
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		key := fn(reflect.ValueOf(list).Index(i).Interface())
		if _, exists := dict[key]; exists == false {
			dict[key] = reflect.ValueOf(list).Index(i).Interface()
		}
	}
	ret := make([]interface{}, 0)
	result := make(map[interface{}]bool)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		key := fn(reflect.ValueOf(list).Index(i).Interface())
		if _, exists := result[key]; exists == false {
			result[key] = true
			ret = append(ret, reflect.ValueOf(list).Index(i).Interface())
		}
	}
	return ret
}

func InterfaceList(list interface{}) []interface{} {
	ret := make([]interface{}, 0)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		ret = append(ret, reflect.ValueOf(list).Index(i).Interface())
	}
	return ret
}

func contract(list interface{}, fn func(interface{})interface{}) []interface{} {
	ret := make([]interface{}, 0)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		elem := reflect.ValueOf(list).Index(i)
		if elem.Kind() == reflect.Slice {
			for v := 0; v < reflect.ValueOf(elem).Len(); v++ {
				elem2 := reflect.ValueOf(elem).Index(v).Interface()
				if fn != nil {
					elem2 = Uniq(elem2, fn)
				}
				ret = append(ret, elem2)
			}
		}
	}
	return ret
}

func Union(list interface{}, fn func(interface{})interface{}) []interface{} {
	ret := contract(list, nil)
	return Uniq(ret, fn)
}

func Intersection(inner interface{}, fn func(interface{})interface{}) []interface{} {
	list := contract(inner, fn)
	dict := make(map[interface{}]int64)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		key := fn(reflect.ValueOf(list).Index(i).Interface())
		if _, exists := dict[key]; exists == false {
			dict[key] = int64(0)
		}
		dict[key] += 1
	}
	ret := Filter(list, func(x interface{})bool{
		key := fn(x)
		return dict[key] > 1
	})
	return ret
}

func Difference(inner interface{}, fn func(interface{})interface{}) []interface{} {
	list := contract(inner, fn)
	dict := make(map[interface{}]int64)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		key := fn(reflect.ValueOf(list).Index(i).Interface())
		if _, exists := dict[key]; exists == false {
			dict[key] = int64(0)
		}
		dict[key] += 1
	}
	ret := Filter(list, func(x interface{})bool{
		key := fn(x)
		return dict[key] == 1
	})
	return ret
}

func Sum(list interface{}, fn func(interface{})int64) int64 {
	ret := int64(0)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		ret += fn(reflect.ValueOf(list).Index(i).Interface())
	}
	return ret
}

func SumFloat(list interface{}, fn func(interface{})float64) float64 {
	ret := float64(0)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		ret += fn(reflect.ValueOf(list).Index(i).Interface())
	}
	return ret
}

func Map(list interface{}, fn func(interface{})interface{}) []interface{} {
	ret := make([]interface{}, reflect.ValueOf(list).Len())
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		ret[i] = fn(reflect.ValueOf(list).Index(i).Interface())
	}
	return ret
}

func Each(list interface{}, fn func(interface{})) {
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		fn(reflect.ValueOf(list).Index(i).Interface())
	}
}

func ForEach(list interface{}, fn func(int, interface{})) {
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		fn(i, reflect.ValueOf(list).Index(i).Interface())
	}
}

func Pluck(list interface{}, key string) []interface{} {
	ret := make([]interface{}, 0)
	for i := 0; i < reflect.ValueOf(list).Len(); i++ {
		elem := reflect.ValueOf(list).Index(i)
		if elem.Kind() == reflect.Interface {
			value := elem.Elem().FieldByName(key).Interface()
			ret = append(ret, value)
		} else {
			value := elem.FieldByName(key).Interface()
			ret = append(ret, value)
		}
	}
	return ret
}

func First(list interface{}) interface{} {
	if reflect.ValueOf(list).Len() == 0 {
		return list
	}
	return reflect.ValueOf(list).Index(0).Interface()
}


