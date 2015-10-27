package main

import (
	"fmt"
	"goset"
	"testing"
)

func TestInit(t *testing.T) {
	type int_slice []int
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("catch error:", err)
		} else {
			t.Fatal("not detect error")
		}
	}()
	s1 := int_slice{1, 2, 3}
	s2 := int_slice{1, 2, 3}
	goset.SliceInit([]int_slice{s1, s2})
}

func TestContain(t *testing.T) {
	content := []string{"abc", "xyz"}
	s := goset.SliceInit(content)

	if !s.Contain("abc") {
		t.Fatal("abc not in set, expect exists")
	} else {
		fmt.Println("abc in set")
	}
	if !s.Contain("xyz") {
		t.Fatal("xyz not in set, expect exists")
	} else {
		fmt.Println("xyz in set")
	}
}
func TestInsertRemove(t *testing.T) {
	s := goset.Set{}

	s.Insert(1)
	s.Insert("hello")

	if !s.Contain(1) {
		t.Fatal("number", 1, "not in set, expect exists")
	} else {
		fmt.Println("1 in set")
	}
	if !s.Contain("hello") {
		t.Fatal("hello not in set, expect exists")
	} else {
		fmt.Println("hello in set")
	}

	s.Remove(1)
	s.Remove("hello")
	if s.Contain(1) {
		t.Fatal("number", 1, "exists, expect remove")
	}
	if s.Contain("hello") {
		t.Fatal("hello exists, expect remove")
	}
	if s.Len() != 0 {
		t.Fatal("length of s should be zero!")
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("no paic!")
			t.Fatal("expect panic!")
		}
	}()
	s := goset.Set{}
	//s.Insert(3)
	s.Insert([]int{1, 2, 3})
}
