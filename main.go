package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

// Represents one set of question and answer
type Exercise struct {
	question string
	ans      string
}

func main() {
	// read quizFile name from flag or default problems.csv
	quizFileName := *flag.String("csv",
		"problems.csv",
		"CSV file in the format question, answer")
	flag.Parse()

	// read the file and parse it into list of exercises
	exercises := openExerciseFile(quizFileName)
	score := 0

	// display welcome message
	fmt.Println("WELCOME TO OUR QUIZ")
	fmt.Println("Press Enter key to start.")
	fmt.Scanln()

	// from each exercise ask question and check for answer
	for i, exercise := range exercises {
		// show question
		fmt.Printf("%d. %s\n", i, exercise.question)

		// read answer
		var ans string
		fmt.Println("Answer: ")
		n, err := fmt.Scanln(&ans)

		// no error and no multiple answers
		if err != nil && n > 1 {
			log.Fatal(err)
		}

		// check for correct answer
		if ans == exercise.ans {
			score++
		}
		// display new line to show next question
		fmt.Println()
	}
	// Display stats at last
	fmt.Printf("You scored %d out of %d.\n", score, len(exercises))
}

// openExerciseFile opens the given file and parses the exercises.
func openExerciseFile(fileName string) []Exercise {
	exercises := make([]Exercise, 0, 100)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		exercises = append(exercises, Exercise{record[0], record[1]})
	}
	return exercises
}
