package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func loadQuestions(filename string) ([]Question, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var questions []Question
	err = json.Unmarshal(file, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func shuffleQuestions(questions []Question, r *rand.Rand) {
	r.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}

func runQuiz(questions []Question) {
	score := 0
	totalQuestions := len(questions)

	reader := bufio.NewReader(os.Stdin)

	for i, q := range questions {
		fmt.Printf("Question %d/%d: %s\n", i+1, totalQuestions, q.Question)
		fmt.Print("Your answer: ")

		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if strings.EqualFold(answer, q.Answer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Sorry, the correct answer is: %s\n", q.Answer)
		}
		fmt.Println()
	}

	fmt.Printf("Quiz finished! Your score: %d/%d\n", score, totalQuestions)
}

func main() {
	fmt.Println("Welcome to GoQuiz!")

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	questions, err := loadQuestions("questions.json")
	if err != nil {
		log.Fatal("Error loading questions:", err)
	}

	fmt.Printf("Loaded %d questions\n", len(questions))
	fmt.Println("Shuffling questions...")
	shuffleQuestions(questions, r)
	fmt.Println("Let's begin the quiz!")
	fmt.Println()

	runQuiz(questions)
}
