package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
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
		fmt.Println(record)
		if index > 0 {
			question := Question{
				Text: record[0],
				Options: record[1:5],
				Answer: toInt(record[6]),
			}

			g.Questions = append(g.Questions, question)

		}
	}
}

func main() {
	game := &Game{}
	go game.ProccessCSV()
	game.Init()

	fmt.Println(game.Questions)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i

}