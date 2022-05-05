# Go ==

## example
```Go
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
	fmt.Println(a1 == a2)
}
```

output:
```bash
go test -v --count=1 ./equalv2.go ./equalv2_test.go

=== RUN   TestEqualCar
true
--- PASS: TestEqualCar (0.00s)
```

而EqualAnimal在编译阶段就报错，因此无法通过变异。


## == 和 != 的工作原理
在Golang中，可以使用==和!=比较两个可比较类型的相等性，其中可比较类型包括：

1. bool类型
2. 数值类型

    两个值具有相同的数值类型或可转化为相同的数值类型
3. 字符串
4. 指针

    比较两个指针是否为nil或指向内存中同一个值
5. Channel

    比较两个Channel是否都为nil或者来自同一个make调用
6. struct或数组（不是slice）是由上述可比较类型组成，那么其也可以使用==比较相等性

**但是map和slice无法通过==比较相等性**


## 如何比较map、slice以及包含不可比较类型的struct的相等性
使用reflect.DeepEqual，例如：
```Go
func EqualAnimal() {
	a1 := Animal{ID: 1, Attrs: []string{"1"}}
	a2 := Animal{ID: 1, Attrs: []string{"1"}}
	// fmt.Println(a1 == a2)
	equal := reflect.DeepEqual(a1, a2)
	fmt.Println(equal)
}
```