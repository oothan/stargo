package main

func main() {
	c := make(chan string)
	c <- "No ones likes my chan"
}
