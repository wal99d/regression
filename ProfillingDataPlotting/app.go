package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot/plotter"

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
	//range over columns
	for _, colName := range df.Names() {
		//Create pltter.Values and fill it with columns values
		plotVals := make(plotter.Values, df.Nrow())
		for i, floatVal := range df.Col(colName).Float() {
			plotVals[i] = floatVal
			//fmt.Printf("plotVals[%d]=%0.2f\n", i, floatVal)
		}
		p, err := plot.New()
		must(err)
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		h, err := plotter.NewHist(plotVals, 16)
		must(err)

		h.Normalize(1)
		p.Add(h)

		must(p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"))
	}
}
