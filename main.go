package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	location := "problems.csv"

	if len(os.Args) > 1 {
		location = os.Args[1]
	}

	f := openFile(location)
	defer f.Close()

	data := readFile(f)

	var i int = 0
	var input string

	fmt.Println("Answer the questions")

	for i < len(data) {
		fmt.Println(data[i][0])
		if readUserInput(&input); input == "Q" {
			break
		}

		_, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Only integers")
		} else {
			if input == data[i][1] {
				fmt.Println("Correct!")
			} else {
				fmt.Println("Wrong! the correct answer is:", data[i][1])
			}
			i++
		}

	}
}

func readFile(file *os.File) [][]string {
	r := csv.NewReader(file)

	data, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func openFile(location string) *os.File {
	file, err := os.Open(location)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func readUserInput(input *string) *string {
	_, err := fmt.Scanln(input)
	if err != nil {
		log.Fatal(err)
	}

	return input
}
