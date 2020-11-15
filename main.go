package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"flag"
	"log"
	"time"
	// "reflect"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A CSV file in the format of questions and answers")
	timeLimit := flag.Int("tlimit", 5, "Time limit to complete the quiz")
	

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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0

	for _, problem := range parseLines(lines) {
		fmt.Println(problem.question)

		answerCh := make(chan string)
		
		go func() {
			var answer string
			fmt.Scanf("%v\n", &answer)
			answerCh <-answer
		}()

		select {
		case <-timer.C:
			printScore(score, len(lines))
			return
		case answer := <-answerCh:
			fmt.Println(problem.question)
			if answer == problem.answer {
				score ++
			}
		}
	}
	printScore(score, len(lines))
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

func printScore(score, count int) {
	fmt.Printf("\nYou scored %d out of %d\n", score, count)
}