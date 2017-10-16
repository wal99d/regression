package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
)
import "gonum.org/v1/plot/plotter"

import "gonum.org/v1/plot/vg"

func must(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

func main() {
	f, err := os.Open("../data/Advertising.csv")
	must(err)

	df := dataframe.ReadCSV(f)

	//yVal holds sales float values
	yVal := df.Col("Sales").Float()

	//Loop through each feature/Column in our Advertising dataset
	for _, colName := range df.Names() {
		pts := make(plotter.XYs, df.Nrow())
		for i, floatVal := range df.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVal[i]
		}
		p, err := plot.New()
		must(err)
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"

		p.Add(plotter.NewGrid())
		s, err := plotter.NewScatter(pts)
		must(err)
		s.GlyphStyle.Radius = vg.Points(3)
		p.Add(s)
		must(p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"))
	}

}
