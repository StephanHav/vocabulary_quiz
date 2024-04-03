package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randInt(rng *rand.Rand, min, max int) int {
	return min + rng.Intn(max-min)
}

func openCSV(fileName string) ([][]string, error) {

	//Open the CSV file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Read the CSV data into a slice of slices
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func runQuiz(dataframe [][]string, column int) {

	// Create a new source and rng at the start of each quiz
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Define the amount of words you want to practice
	var num_questions int
	fmt.Print("How many words do you want to practice?", "\n", "Please type the amount: ")
	fmt.Scanln(&num_questions)
	fmt.Println("\n")
	fmt.Println("Please write the English translation of the following words: ")

	// Quiz
	var correct_answers, wrong_answers int
	for i := 0; i < num_questions; i++ {

		// Pick random word to practice
		var word_index int = randInt(rng, 0, len(dataframe))
		fmt.Print(dataframe[word_index][column], " = ")

		// Prompt user for English translation
		var translation string
		fmt.Scanln(&translation)

		// Check answer
		if translation == dataframe[word_index][column+1] {
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

func main() {

	dataframe, err := openCSV("words.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Skip the first row (header)
	dataframe = dataframe[1:]

	// Prompt for language selection
	fmt.Println("\n", "Select a language (1 for Russian, 2 for Spanish, 3 for Dutch):", "\n")
	var language int
	fmt.Scanln(&language)
	fmt.Println("\n")

	// Adjust column based on selected language
	// Assuming Russian=0, Spanish=1, Dutch=2 for simplicity
	column := (language - 1) * 2

	//Run quiz
	runQuiz(dataframe, column)
}
