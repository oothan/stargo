package main

import (
	"fmt"
	"sort"
)

func main() {
	employeeSalary := map[string]float64{
		"Blake":  640.00,
		"Parker": 580.45,
		"Dakota": 700.50,
	}

	fmt.Printf("Parker Salary %.2f\n", employeeSalary["Parker"])

	blakeSalary, ok := employeeSalary["Blake"]
	if ok {
		fmt.Printf("Value %.2f was found.\n", blakeSalary)
	} else {
		fmt.Println("The specified key was not found.")
	}

	if jordanSalary, ok := employeeSalary["Jordan"]; ok {
		fmt.Printf("value %.2f was found. \n", jordanSalary)
		delete(employeeSalary, "Jordan")
	} else {
		fmt.Println("The specified key was not found.")
	}

	_, ok = employeeSalary["Dakota"]
	if ok {
		fmt.Println("Key Exists.")
	} else {
		fmt.Println("Key doesn't exist.")
	}

	for key, value := range employeeSalary {
		fmt.Println(key, value)
	}

	sal := map[string]float64{
		"Blake":  640.00,
		"Parker": 580.45,
		"Dakota": 700.50,
	}

	names := make([]string, 0, len(sal))
	for n := range sal {
		names = append(names, n)
	}
	fmt.Println(names)

	sort.Strings(names)
	fmt.Println(names)
	for _, n := range names {
		fmt.Printf("%s\t%.2f\n", n, sal[n])
	}

	values := make([]float64, 0, len(sal))
	for _, val := range sal {
		values = append(values, val)
	}
	fmt.Println(values)

	sort.Float64s(values)
	fmt.Println(values)

}
