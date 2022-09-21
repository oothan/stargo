package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)

	intB, _ := json.Marshal(1)
	fmt.Println(intB)

	strB, _ := json.Marshal("gopher")
	fmt.Println(strB)

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(mapB)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	resD1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	resD2 := &response1{}
	resB1, _ := json.Marshal(resD1)
	fmt.Println(resB1)
	_ = json.Unmarshal(resB1, resD2)
	fmt.Println(resD2)

	byt := []byte(`{"num": 6.13, "str": ["a", "b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(reflect.TypeOf(dat["num"]))
	fmt.Println(dat["str"])

	num := dat["num"].(float64)
	fmt.Println(reflect.TypeOf(num))

}
