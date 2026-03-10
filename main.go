package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type Answer struct {
	Response string
	Explanation string
}

type Question struct {
	Text    string
	Options []string
	Answer  Answer
}

type Game struct {
	Nome      string
	Points    int
	Questions []Question
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

func (g *Game) ProccessCSV() {
	file, err := os.Open("quizgo.csv")

	if err != nil {
		panic("Erro ao ler arquivo")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Erro ao ler csv")
	}

	for index, record := range records {
		if index > 0 {
			anwser := Answer{
				Response: record[5],
				Explanation: record[6],
			}
			question := Question{
				Text: record[0],
				Options: record[1:5],
				Answer: anwser,
			}
			
			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *Game) Run(){
	// exibir pergunta ao usuário
	for index, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", index+1, question.Text)

		for in, option := range question.Options {
			fmt.Printf("[%s] %s \n", numberToLetter(in+1), option)
		}

	}
}

func main() {
	game := &Game{}
	go game.ProccessCSV()
	game.Init()
	game.Run()
}


func numberToLetter(n int) string {
	return string(rune('a' + n - 1))
}