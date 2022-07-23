package unicodeplot

import "fmt"

type Color int

const escape = "\x1b"

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func (c Color) sequence() int {
	return int(c)
}

func Colorize(s string, c Color) string {
	return fmt.Sprintf("%s[%dm%s%s[0m", escape, c.sequence(), s, escape)
}

func Bold(s string) string {
	return fmt.Sprintf("%s[1m%s%s[0m", escape, s, escape)
}
