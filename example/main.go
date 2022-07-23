package main

import (
	"fmt"

	"github.com/hidetatz/unicodeplot/barplot"
)

func main() {
	if err := barplot.Print(
		[]string{"tokyo", "sapporo"},
		[]float64{123.1, 34.23},
		barplot.WithTitle("temperature"),
	); err != nil {
		fmt.Println(err)
	}
}
