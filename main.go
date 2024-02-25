package main 

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main() {
values := []float64{1, 2, 3, 4, 5, 6}
   weights := []float64{1, 1, 1, 1, 1, 1} //has the same effects as nil
   fmt.Println(stat.Mean(values, weights))
}