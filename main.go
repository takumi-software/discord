package main

import (
	"fmt"
	"github.com/takumi-software/discord/stats"
)

func main(){
	population := []float64{0, 1, 2, 58}
	mean := stats.Mean(population)
	fmt.Println(mean)
}