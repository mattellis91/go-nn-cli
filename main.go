package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	_ "github.com/charmbracelet/bubbletea"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: go run main.go -train <train.csv> -expected <expected.csv>")
		return
	} else {
		trainPtr := flag.String("train", "xorTrain.csv", "a string")
		expectedPtr := flag.String("expected", "xorExpected.csv", "a string")
		flag.Parse()
		fmt.Println("train:", *trainPtr)
		fmt.Println("expected:", *expectedPtr)

		trainingData := ReadDataFromCSV(*trainPtr)
		expectedData := ReadDataFromCSV(*expectedPtr)

		nn := NewNN(len(trainingData[0]), len(trainingData[0]), len(expectedData[0]), 0.3)

		fmt.Printf("%+v\n", nn)

	}
}

func ReadDataFromCSV(filePath string) [][]string {

	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}

	return records

}