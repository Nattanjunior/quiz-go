package main

import (
	"fmt"
	"bufio"
	"os"
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

func (g *Game) Init() {
	fmt.Println("Seja bem-vindo ao jogo de perguntas e respostas!")
	fmt.Println("Qual é o seu nome?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n') 

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Nome = name
	fmt.Printf("Vamos ao jogo %s", g.Nome)
	
}

func main() {
	game := &Game{}
	game.Init()
}

