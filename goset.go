package goset

import (
	"reflect"
)

type Set map[interface{}]bool

func SliceInit(in interface{}) Set {
	ret := Set{}
	V := reflect.ValueOf(in)
	n := V.Len()
	for i := 0; i < n; i++ {
		raw_v := V.Index(i).Interface()
		v := reflect.ValueOf(raw_v)
		if v.Kind() == reflect.Slice {
			panic("cant insert an array as element of set")
		}
		ret[raw_v] = true
	}
	return ret
}

func (s *Set) Insert(it interface{}) {
	V := reflect.ValueOf(it)

	if V.Kind() == reflect.Slice {
		panic("cant insert an array as element of set")
	}
	(*s)[it] = true
}
func (s *Set) Remove(it interface{}) {
	V := reflect.ValueOf(it)

	if V.Kind() == reflect.Slice {
		return
	}
	if v, ok := (*s)[it]; ok && v {
		delete(*s, it)
	}
}

func (s Set) Len() int {
	V := reflect.ValueOf(s)
	return len(V.MapKeys())
}

func (s Set) Contain(it interface{}) bool {
	V := reflect.ValueOf(it)

	if V.Kind() == reflect.Slice {
		return false
	}
	if v, ok := s[it]; ok && v {
		return v
	} else {
		return false
	}
}
