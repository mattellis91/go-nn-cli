package main

import (
	"fmt"
	"math"
	"math/rand"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	data := make([]float64, 25)
	for i := 0; i < 25; i++ {
		data[i] = rand.NormFloat64()
	}
	a := mat.NewDense(5, 5, data)
	fmt.Printf("%v", mat.Formatted(a))

	rand.Seed(int64(0))

	p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

type NN struct {
	inputNodes int
	hiddenNodes int
	outputNodes int
	learningRate float64
}

func NewNN(inputNodes, hiddenNodes, outputNodes int, learningRate float64) *NN {
	return &NN{
		inputNodes: inputNodes,
		hiddenNodes: hiddenNodes,
		outputNodes: outputNodes,
		learningRate: learningRate,
	}
}

func (n *NN) Train() {

}

func (n *NN) Predict() {

}

func Sig(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}