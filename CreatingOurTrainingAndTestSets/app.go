package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func must(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

func main() {
	f, err := os.Open("../data/Advertising.csv")
	must(err)
	defer f.Close()
	df := dataframe.ReadCSV(f)
	trainingNum := (4 * df.Nrow()) / 5
	testNum := df.Nrow() / 5

	if trainingNum+testNum < df.Nrow() {
		trainingNum++
	}

	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)
	//filling training idices
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}
	//filling test indices
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}
	trainingDf := df.Subset(trainingIdx)
	testDf := df.Subset(testIdx)

	//Map used for writing to training and test csv files
	setMap := map[int]dataframe.DataFrame{
		0: trainingDf,
		1: testDf,
	}

	for idx, setName := range []string{"training.csv", "testing.csv"} {
		tmpF, err := os.Create(setName)
		must(err)
		w := bufio.NewWriter(tmpF)
		must(setMap[idx].WriteCSV(w))
	}
}
