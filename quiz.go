package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	// Open the CSV file
	file, err := os.Open("words.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Read the CSV data into a slice of slices
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the slice of slices to a dataframe
	var dataframe [][]string = rows

	// Define the amount of words you want to practice
	var num_questions int
	fmt.Print("How many words do you want to practice?", "\n", "Please type the amount: ")
	fmt.Scanln(&num_questions)
	fmt.Println("\n")
	fmt.Println("Please write the English translation of the following words: ")

	// Quiz
	var correct_answers int = 0
	var wrong_answers int = 0
	for i := 0; i < num_questions; i++ {

		// Make sure new seed is generated every run
		rand.Seed(time.Now().UnixNano())

		// Pick random word to practice
		var word_index int = randInt(0, len(dataframe))
		fmt.Print(dataframe[word_index][0], " = ")

		// Prompt user for English translation
		var translation string
		fmt.Scanln(&translation)

		// Check answer
		if translation == dataframe[word_index][1] {
			correct_answers++
			fmt.Println("Correct!", "\n")
		} else {
			wrong_answers++
			fmt.Println("Wrong, it is actually: ", dataframe[word_index][1], "\n")
		}
	}

	fmt.Println("That's it! You finished.", "\n")
	fmt.Println("Your final score out of", num_questions, "words is:", "\n")
	fmt.Println(correct_answers, "Correct")
	fmt.Println(wrong_answers, "Wrong", "\n")

}
