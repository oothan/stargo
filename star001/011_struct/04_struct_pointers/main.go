package main

import (
	"fmt"
	"github.com/fxtlabs/date"
	"time"
)

type employee struct {
	name, job    string
	lastLoggedIn string
	DOB          date.Date
}

func main() {
	var emp employee
	emp.name = "Luke"
	emp.job = "Go Programmer"
	emp.lastLoggedIn = time.Now().Format(time.RFC850)
	emp.DOB = date.New(1998, time.July, 20)

	fmt.Println(emp)
	fmt.Printf("emp : %+v \n", emp)

	p := &emp
	p.name = "Jack"
	//emp.name = "Go Expert"
	fmt.Println(*p)
	fmt.Println(p)

	fmt.Printf("%x %x", &emp.name, &p.name)
}
