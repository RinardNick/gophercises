package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// set flag to allow user to input name of problems csv file
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
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

	correct := 0 // initiate counter for correct answers
	for i, p := range problems {
		correct = testProblem(i, p, correct) // test user with problem
	}
	fmt.Printf("You scored %d/%d ", correct, len(problems))
}

func testProblem(i int, p problem, correct int) int {
	fmt.Printf("Problem #%d: %s = ", i+1, p.question)
	var response string          // initiate variable for response from user
	fmt.Scanf("%s\n", &response) // scan for response (\n captures when user hits enter)
	if response == p.answer {
		fmt.Println("Correct")
		correct++
	} else {
		fmt.Println("Incorrect")
	}
	return correct
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