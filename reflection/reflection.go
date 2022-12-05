package main

import "reflect"

func walk(x interface{}, fn func(input string)) {

	val := getValueOfInterface(x)
	
	switch val.Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			walk(val.MapIndex(k).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			walk(field.Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}

	//if val.Kind() == reflect.Slice {
	//	for i := 0; i < val.Len(); i++ {
	//		walk(val.Index(i).Interface(), fn)
	//	}
	//	return
	//}

	//for i := 0; i < val.NumField(); i++ {
	//	field := val.Field(i)

	//	switch field.Kind() {
	//	case reflect.Struct:
	//		walk(field.Interface(), fn)
	//	case reflect.String:
	//		fn(field.String())
	//	}
	//}
}

func getValueOfInterface(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
