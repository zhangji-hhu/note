package golang

import "testing"

// 逃逸分析
// go test -v --count=1 -gcflags="-m -l" ./
// go test -v --count=1 -gcflags="-N -m -l" ./

func TestEqualPeople(t *testing.T) {
	EqualPeople()
}

func TestEqualPeople2(t *testing.T) {
	EqualPeople2()
}

func TestEqualTeacher(t *testing.T) {
	EqualTeacher()
}

func TestEqualTeacher2(t *testing.T) {
	EqualTeacher2()
}

func TestPeopleTeacher(t *testing.T) {
	EqualPeopleTeacher()
}

func TestPeopleTeacher2(t *testing.T) {
	EqualPeopleTeacher2()
}
