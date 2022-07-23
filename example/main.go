package main

import (
	"fmt"

	"github.com/hidetatz/unicodeplot"
	"github.com/hidetatz/unicodeplot/barplot"
)

func main() {
	if err := barplot.Print(
		[]string{"tokyo", "sapporo", "okinawa"},
		[]float64{0.1, 3.23, 4.965321},
		barplot.WithColor(unicodeplot.Magenta),
		barplot.WithTitle("temp"),
		barplot.WithBorder("barplot"),
		barplot.WithMargin(5),
		barplot.WithWidth(80),
		barplot.WithShowBorder(false),
	); err != nil {
		fmt.Println(err)
	}
}
