package barplot

import "github.com/hidetatz/unicodeplot"

type Option func(*barPlot)

func WithTitle(title string) Option {
	return func(p *barPlot) {
		p.title = title
	}
}

func WithWidth(width int) Option {
	return func(p *barPlot) {
		p.width = width
	}
}

func WithColor(color unicodeplot.Color) Option {
	return func(p *barPlot) {
		p.color = color
	}
}

func WithSymbol(symbol string) Option {
	return func(p *barPlot) {
		p.symbol = symbol
	}
}

func WithMargin(margin int) Option {
	return func(p *barPlot) {
		p.margin = margin
	}
}

func WithShowBorder(showBorder bool) Option {
	return func(p *barPlot) {
		p.showBorder = showBorder
	}
}

func WithBorder(border string) Option {
	return func(p *barPlot) {
		p.border = border
	}
}
