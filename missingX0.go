/*
math problems
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Create a data list. This example shows an i+i*i calculated list.
	data := createDataList()

	// Display the range of the data list and generate the missing value. Set timer.
	missingValue, timerStart := generateMissingValue(data)

	// Display time and compare the answer.
	compareAnswer(missingValue, timerStart)
}


func generateMissingValue(data []int) (int, time.Time) {
	length := len(data)
	fmt.Printf("\nIn the following ordered list, you will find characters or numbers ranging from %d to %d. \n\n", data[0], data[length-1])
	missingValue := data[rand.Intn(length)]

	for _, v := range data {
		if v != missingValue {
			fmt.Print(v, " ")
		}
	}
	fmt.Println("\n\nWhat charactor or number is missing?")
	timerStart := time.Now()  // Timer set

	return missingValue, timerStart
}

func createDataList() []int {
	var data []int

	for i := 1; i < 10; i++ {
		data = append(data, i+i*i)
	}
	return data
}

// compareAnswer function checks user input against the missing value.
func compareAnswer(missingValue int, timerStart time.Time) {
	var answer int
	fmt.Print("Your answer: ")
	_, err := fmt.Scan(&answer)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	if answer == missingValue {
		timerEnd := time.Now()

		fmt.Println("Correct!")
		fmt.Println("Time taken: ", timerEnd.Sub(timerStart))
	} else {
		fmt.Printf("Oops! The answer is %d. Next!", missingValue)
	}
}
