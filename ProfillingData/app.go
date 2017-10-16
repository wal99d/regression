package main

import (
	"fmt"
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
	fmt.Printf("%v\n", df.Describe())
}
