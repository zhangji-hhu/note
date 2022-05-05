# golang equal

## example 1
```Go
type People struct {
	ID   int64
	Name string
}

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
```

### output
```bash
=== RUN   TestEqualPeople
0xc000049738 0xc000049720 false
false
--- PASS: TestEqualPeople (0.00s)
=== RUN   TestEqualPeople2
p1 addr=0xc0000be060
p2 addr=0xc0000be078
0xc0000be060 0xc0000be078 false
false
--- PASS: TestEqualPeople2 (0.00s)
```


## example 2
```Go
type Teacher struct{}

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
```

### output
```bash
=== RUN   TestEqualTeacher
0xc00004673f 0xc00004673f false
false
--- PASS: TestEqualTeacher (0.00s)
=== RUN   TestEqualTeacher2
t1 addr=0x121d050
t2 addr=0x121d050
0x121d050 0x121d050 true
true
--- PASS: TestEqualTeacher2 (0.00s)
```

结合example 1和example 2可以看出，初始化两个带有字段的结构体，为其分配了不同的内存地址；而初始化两个空结构体，分配的地址相同，但是使用==比较结果却不同。

查看编译过程，找出原因。运行如下命令：
```bash
go test -v --count=1 -gcflags="-m -l" ./
```
输出如下：
```bash
# github/note/golang [github/note/golang.test]
./equal.go:29:8: &Teacher{} does not escape
./equal.go:30:8: &Teacher{} does not escape
./equal.go:32:13: ... argument does not escape
./equal.go:32:17: t1 == t2 escapes to heap
./equal.go:36:8: &Teacher{} escapes to heap
./equal.go:37:8: &Teacher{} escapes to heap
./equal.go:38:12: ... argument does not escape
./equal.go:39:12: ... argument does not escape
./equal.go:41:13: ... argument does not escape
./equal.go:41:17: t1 == t2 escapes to heap
./equal_test.go:7:22: t does not escape
./equal_test.go:11:23: t does not escape
./equal_test.go:15:23: t does not escape
./equal_test.go:19:24: t does not escape
# github/note/golang.test
/var/folders/vj/mqq03d1n41ggtrlcfr_8nh7w0000gp/T/go-build2625861694/b001/_testmain.go:47:42: testdeps.TestDeps{} escapes to heap
=== RUN   TestEqualTeacher
0xc00004673f 0xc00004673f false
false
--- PASS: TestEqualTeacher (0.00s)
=== RUN   TestEqualTeacher2
t1 addr=0x121d050
t2 addr=0x121d050
0x121d050 0x121d050 true
true
--- PASS: TestEqualTeacher2 (0.00s)
PASS
ok      github/note/golang      1.097s
```

可以看到，对于example2来说，TestEqualTeacher2变量逃逸到了heap上，根据代码可知造成逃逸的原因是使用了fmt.Printf，该方法内部有大量的反射操作，从而编译器在编译阶段使该变量分配到heap上。

**为什么逃逸后的空结构体值相等**

主要是由于 Go Runtime的优化，Go Runtime在对逃逸的空（0字节）变量，在heap上都指向同一个zerobase的变量

```Go
// runtime/malloc.go
var zerobase uintptr
```

知道上述问题是由于编译器优化后导致的，因此，禁用编译器优化查看结果：
```bash
go test -v --count=1 -gcflags="-N -m -l" ./ 

// output
# github/note/golang [github/note/golang.test]
./equal.go:29:8: &Teacher{} does not escape
./equal.go:30:8: &Teacher{} does not escape
./equal.go:32:13: ... argument does not escape
./equal.go:32:17: t1 == t2 escapes to heap
./equal.go:36:8: &Teacher{} escapes to heap
./equal.go:37:8: &Teacher{} escapes to heap
./equal.go:38:12: ... argument does not escape
./equal.go:39:12: ... argument does not escape
./equal.go:41:13: ... argument does not escape
./equal.go:41:17: t1 == t2 escapes to heap
./equal_test.go:7:22: t does not escape
./equal_test.go:11:23: t does not escape
./equal_test.go:15:23: t does not escape
./equal_test.go:19:24: t does not escape
# github/note/golang.test
/var/folders/vj/mqq03d1n41ggtrlcfr_8nh7w0000gp/T/go-build402340628/b001/_testmain.go:47:42: testdeps.TestDeps{} escapes to heap
=== RUN   TestEqualTeacher
0xc0000496fe 0xc0000496fe true
true
--- PASS: TestEqualTeacher (0.00s)
=== RUN   TestEqualTeacher2
t1 addr=0x121d050
t2 addr=0x121d050
0x121d050 0x121d050 true
true
--- PASS: TestEqualTeacher2 (0.00s)
PASS
ok      github/note/golang      1.129s
```
可以看到两种情况的比较结果都为true了

**todo:** Go编译器变量结构