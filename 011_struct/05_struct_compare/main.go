package main

import (
	"fmt"
	"unsafe"
)

type book struct {
	name   string
	author string
}

func main() {
	b1 := book{"Go Programming", "Author 1"}
	b2 := book{"Go Programming", "Author 1"}
	b3 := book{"C++ Programming", "Author 2"}

	fmt.Printf("LOC1 - %v \n", b1)

	fmt.Println(compare1(b1, b2))
	fmt.Println(compare1(b1, b3))
	fmt.Printf("LOC2 - %v \n\n", b1)

	fmt.Println(compare2(&b1, &b2))
	fmt.Printf("LOC3 - %v \n\n", b1)

	p1 := &book{"Go Web Programming", "Author 3"}

	p2 := new(book)
	*p2 = book{"Go Web Programming", "Author 3"}

	fmt.Printf("LOC4 - &p1=%v *p1=%v \n", &p1, *p1)
	fmt.Println(compare2(p1, p2))
	fmt.Printf("LOC5 - &p1=%v *p1=%v \n", &p1, *p1)

	testSizeOfStruct()
}

func compare1(p book, q book) bool {
	if p.name == q.name && p.author == q.author {
		p.name = "Java Programming"
		return true
	}
	return false
}

func compare2(p *book, q *book) bool {
	if p.name == q.name && p.author == q.author {
		p.name = "Python Programming"
		return true
	}
	return false
}

func testSizeOfStruct() {
	type values struct {
		val1 int32
		val2 int32
		val3 int64
		val4 int64
	}

	x := values{10, 20, 30, 40}
	p := &x
	fmt.Println(unsafe.Sizeof(p), unsafe.Sizeof(*p), unsafe.Sizeof(x))

	q := &values{11, 21, 31, 41}
	fmt.Println(unsafe.Sizeof(q), unsafe.Sizeof(*q))
}
