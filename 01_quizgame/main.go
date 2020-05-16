package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// set flag to allow user to input name of problems csv file
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "time limit in seconds for the quiz")
	flag.Parse()
	// open csv file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV: %s", *csvFilename))
	}
	// read in csv data
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse CSV: %s" ,*csvFilename))
	}
	// parse the problems into problem structure
	problems := parseProblems(lines)
	// create a timer for the quiz
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0 // initiate counter for correct answers

	problemLoop:
		for i, p := range problems {
			fmt.Printf("Problem #%d: %s = ", i+1, p.question)
			// create answer channel for strings to return scan response
			responseCh := make(chan string)
			// closure (anonymous function that uses data defined outside of it) to scan for answer
			go func() {
				var response string          // initiate variable for response from user
				fmt.Scanf("%s\n", &response) // scan for response (\n captures when user hits enter)
				responseCh <- response
			}()

			select {
			case <- timer.C:
				fmt.Println("You ran out of time!!")
				break problemLoop
			case response := <- responseCh:
				if response == p.answer {
					fmt.Println("Correct")
					correct++
				} else {
					fmt.Println("Incorrect")
				}
			}
		}
		fmt.Printf("You scored %d / %d ", correct, len(problems))
}

func parseProblems(lines [][]string) []problem {
	prob := make([]problem, len(lines))
	for i, line := range lines {
		prob[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return prob
}

type problem struct {
	question string
	answer string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}