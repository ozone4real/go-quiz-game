package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"flag"
	"log"
	// "reflect"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A CSV file in the format of questions and answers")

	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(err)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit(err)
	}

	score := 0

	for _, problem := range parseLines(lines) {
		var answer string
		fmt.Println(problem.question)
		fmt.Scanf("%v\n", &answer)
		if answer == problem.answer {
			score ++
		}
	}

	fmt.Printf("You scored %d out of %d\n", score, len(lines))

}

func parseLines(problems [][]string) []problem {
	res := make([]problem, len(problems))
	for i, y := range problems {
		question := fmt.Sprintf("Question #%d:\n %v is?\n", i+1, y[0])
		res[i] = problem{ question: question, answer: y[1] }
	}
	return res
}

type problem struct {
	question, answer string
}

func exit(err error) {
	log.Fatal(err)
	os.Exit(1)
}