package barplot

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/hidetatz/unicodeplot"
)

type barPlot struct {
	title      string
	width      int // graph width without margin/key/border
	color      unicodeplot.Color
	symbol     string
	margin     int // margin at the left to the key
	showBorder bool
	border     string
}

// Print prints the bar plot to the stdout.
func Print(keys []string, data []float64, opts ...Option) error {
	return PrintTo(os.Stdout, keys, data, opts...)
}

// PrintTo prints the bar plot to the given writer.
func PrintTo(w io.Writer, keys []string, data []float64, opts ...Option) error {
	opt := &barPlot{
		// default configurations
		title:      "",
		width:      40,
		color:      unicodeplot.Green,
		symbol:     "■",
		margin:     3,
		showBorder: true,
		border:     "barplot",
	}
	for _, o := range opts {
		o(opt)
	}

	// validation
	if utf8.RuneCountInString(opt.symbol) != 1 {
		return fmt.Errorf("symbol must be a single character")
	}

	if opt.width < 10 {
		return fmt.Errorf("width must be greater than 10")
	}

	if !unicodeplot.IsValidBorder(opt.border) {
		return fmt.Errorf("border is invalid")
	}

	if len(keys) != len(data) {
		return fmt.Errorf("keys and data must be the same size")
	}

	if len(keys) == 0 {
		return fmt.Errorf("keys and data must not be empty")
	}

	for _, d := range data {
		if d < 0 {
			return fmt.Errorf("negative data is not supported")
		}
	}

	render(w, keys, data, opt)

	return nil
}

// render renders the bar plot to the given w.
func render(w io.Writer, keys []string, data []float64, opt *barPlot) {
	// find the longest bar key
	maxKeyLen := 0
	for _, key := range keys {
		if l := utf8.RuneCountInString(key); maxKeyLen < l {
			maxKeyLen = l
		}
	}

	// Compute the biggest data and the length of the data in string representation.
	// For example, if the data is [1.2, 3.5, 7.23], the maxData = 7.23, dLen = 4("7.23" consists of 4 characters)
	maxData, dLen := 0.0, 0
	sData := make([]string, len(data))
	for i, d := range data {
		sData[i] = strconv.FormatFloat(d, 'f', -1, 64)
		if maxData < d {
			maxData = d
			dLen = utf8.RuneCountInString(sData[i])
		}
	}
	maxDataSymbols := opt.width - 2 - dLen // 2: spaces for the padding of the data string

	mar := ss(opt.margin)
	border := unicodeplot.GetBorder(opt.border)

	// render title
	if opt.title != "" {
		titlePad := (opt.width - len(opt.title)) / 2 // compute title padding to display the title at the center
		if titlePad < 0 {
			// Just in case the title is too long
			titlePad = 0
		}
		out(w, "%s%s  %s%s\n", mar, ss(maxKeyLen), ss(titlePad), unicodeplot.Bold(opt.title))
	}

	// render top border
	if opt.showBorder {
		out(w, "%s%s %s%s%s\n", mar, ss(maxKeyLen), border.TL, rep(border.T, opt.width), border.TR)
	}

	// render rows
	for i := range keys {
		key, val := keys[i], data[i]
		keyPad := maxKeyLen - utf8.RuneCountInString(key)
		symbolsLen := int(math.Floor(float64(maxDataSymbols) / maxData * val))
		out(w, "%s%s%s %s%s %s%s\n", mar, ss(keyPad), key, border.L, unicodeplot.Colorize(rep(opt.symbol, symbolsLen), opt.color), sData[i], border.R)
	}

	// render bottom border
	if opt.showBorder {
		out(w, "%s%s %s%s%s\n", mar, ss(maxKeyLen), border.BL, rep(border.B, opt.width), border.BR)
	}
}

func out(w io.Writer, format string, a ...any) (n int, err error) {
	return io.WriteString(w, fmt.Sprintf(format, a...))
}

func ss(count int) string {
	return rep(" ", count)
}

func rep(s string, count int) string {
	return strings.Repeat(s, count)
}
