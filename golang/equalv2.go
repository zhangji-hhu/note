package golang

import (
	"fmt"
	"reflect"
)

type Car struct {
	ID   int64
	Name string
}

func EqualCar() {
	car1 := Car{ID: 1, Name: "1"}
	car2 := Car{ID: 1, Name: "1"}
	fmt.Println(car1 == car2)
}

type Animal struct {
	ID    int64
	Attrs []string
}

func EqualAnimal() {
	a1 := Animal{ID: 1, Attrs: []string{"1"}}
	a2 := Animal{ID: 1, Attrs: []string{"1"}}
	// fmt.Println(a1 == a2)
	equal := reflect.DeepEqual(a1, a2)
	fmt.Println(equal)
}
