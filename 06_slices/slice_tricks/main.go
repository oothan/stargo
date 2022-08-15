package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}

	a = append(a, b...)
	fmt.Printf(" (1) a%v b%v \n", a, b)

	c := make([]int, len(a))
	copy(c, a)
	fmt.Printf(" (2) a%v c%v \n", a, c)

	d := append([]int(nil), a...)
	fmt.Printf(" (3) a%v d%v \n", a, d)

	c = append(c[:2], c[4:]...)
	fmt.Printf(" (4) c%v \n", c)

	c = append(c[:1], c[2:]...)
	fmt.Printf(" (5) c%v \n", c)

	c = c[:1+copy(c[1:], c[2:])]
	fmt.Printf(" (6) c%v \n", c)

	fmt.Printf(" (7) d%v \n", d)
	d[2] = d[len(d)-1]
	d = d[:len(d)-1]
	fmt.Printf(" (8) d%v \n", d)

	d = append(d[:2], append(make([]int, 4), d[2:]...)...)
	fmt.Printf(" (9) d%v \n", d)

	d = append(d, make([]int, 3)...)
	fmt.Printf(" (10) d%v \n", d)

	d = append(d[:3], append([]int{9, 10}, d[3:]...)...)
	fmt.Printf(" (11) d%v \n", d)

}
