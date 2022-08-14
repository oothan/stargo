package main

import "fmt"

func main() {

	// array slicing
	var a [3]int
	fmt.Printf("a = %T \n\n", a)

	b := [...]int{1, 2, 3, 4, 5}
	b1 := b[:3]
	fmt.Printf("b = %T  b1 = %T\n\n", b, b1)

	c := [5]int{1, 2, 3, 4, 5}
	c1 := c[:3]
	fmt.Printf("c = %T  c1 = %T\n\n", c, c1)

	s := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	s1 := s[:3]
	fmt.Printf("s = %T s1 = %T\n\n", s, s1)

	// string slicing
	str := "ABCDEFGHIJKL"
	strSlc := str[1:5]
	fmt.Printf("str=%v type(str)=%T  strSlc=%v type(strSlc)=%T\n\n", str, str, strSlc, strSlc)

	// slice append : add end of the slice element
	var f []float32

	f = append(f, 1.2)
	f = append(f, 2.4, 4.5, 8.6)

	fmt.Println(f)

	f = append(f, f...)
	fmt.Println(f)

	// slicing string calculate
	team1 := []string{"Messi", "Pele", "Maradona", goalkeeper()}

	team := append([]string{"Maldini", "Cafu", "Carlos", "Beckenbauer"}, team1...)
	team = append(team, midfielders()...)

	fmt.Println(team)

	for i, name := range team {
		fmt.Println(i, name)
	}

}

func midfielders() []string {
	return []string{"Iniesta", "Zidane", "Platini"}
}
func goalkeeper() string {
	return "Bufon"
}
