package main

type Test struct {
	Name string `json:"name"`
	Flag int64  `json:"flag"`
}

func main() {

	_ = &Test{
		Name: "hla hla",
		Flag: 2,
	}

}
