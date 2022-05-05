package golang

import "fmt"

type People struct {
	ID   int64
	Name string
}

type Teacher struct{}

func EqualPeople() {
	p1 := &People{}
	p2 := &People{}
	println(p1, p2, p1 == p2)
	fmt.Println(p1 == p2)
}

func EqualPeople2() {
	p1 := &People{}
	p2 := &People{}
	fmt.Printf("p1 addr=%p\n", p1)
	fmt.Printf("p2 addr=%p\n", p2)
	println(p1, p2, p1 == p2)
	fmt.Println(p1 == p2)
}

func EqualTeacher() {
	t1 := &Teacher{}
	t2 := &Teacher{}
	println(t1, t2, t1 == t2)
	fmt.Println(t1 == t2)
}

func EqualTeacher2() {
	t1 := &Teacher{}
	t2 := &Teacher{}
	fmt.Printf("t1 addr=%p\n", t1)
	fmt.Printf("t2 addr=%p\n", t2)
	println(t1, t2, t1 == t2)
	fmt.Println(t1 == t2)
}

type Student struct{}

func EqualPeopleTeacher() {
	t1 := &Teacher{}
	s1 := &Student{}
	println(t1, s1)
}

func EqualPeopleTeacher2() {
	t1 := &Teacher{}
	s1 := &Student{}
	fmt.Printf("t1=%p\n", t1)
	fmt.Printf("s1=%p\n", s1)
}
