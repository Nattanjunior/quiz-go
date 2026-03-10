package main

import (
	"fmt"
)


type Questions struct {
	Text string
	Options []string
	Answer int
}

type Game struct {
	Nome  string
	Points int 
	Questions []Questions
}



func main() {
	fmt.Println("Hello, World!")
}

