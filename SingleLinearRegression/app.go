package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func must(err error) {
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

func predict(tv float64) float64 {
	return 7.07 + tv*0.05
}

func main() {
	f, err := os.Open("../data/Advertising.csv")
	must(err)
	defer f.Close()

	df := dataframe.ReadCSV(f)
	//Extract target column
	yVal := df.Col("Sales").Float()
	pts := make(plotter.XYs, df.Nrow())
	ptsPred := make(plotter.XYs, df.Nrow())

	for i, floatVal := range df.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVal[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
	}
	p, err := plot.New()
	must(err)
	p.X.Label.Text = " TV"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())
	s, err := plotter.NewScatter(pts)
	must(err)
	s.GlyphStyle.Radius = vg.Points(3)
	l, _, err := plotter.NewLinePoints(ptsPred)
	must(err)
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	p.Add(s, l)
	must(p.Save(4*vg.Inch, 4*vg.Inch, "regression_line.png"))
}
