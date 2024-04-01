package main

import (
	"math"
)

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

func Sig(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func (n *NN) Train() {

}

func (n *NN) Predict() {

}

func (n *NN) Init() {
	
}