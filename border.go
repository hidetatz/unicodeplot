package unicodeplot

type Border struct {
	TL, TR, BL, BR, T, L, B, R string
}

var (
	BorderSolid = &Border{
		TL: "┌",
		TR: "┐",
		BL: "└",
		BR: "┘",
		T:  "─",
		L:  "│",
		B:  "─",
		R:  "│",
	}

	BorderCorners = &Border{
		TL: "┌",
		TR: "┐",
		BL: "└",
		BR: "┘",
		T:  " ",
		L:  " ",
		B:  " ",
		R:  " ",
	}

	BorderBarplot = &Border{
		TL: "┌",
		TR: "┐",
		BL: "└",
		BR: "┘",
		T:  " ",
		L:  "┤",
		B:  " ",
		R:  " ",
	}
)

func IsValidBorder(b string) bool {
	return b == "solid" || b == "corners" || b == "barplot"
}

func GetBorder(b string) *Border {
	switch b {
	case "solid":
		return BorderSolid
	case "corners":
		return BorderCorners
	case "barplot":
		return BorderBarplot
	default:
		return BorderSolid
	}
}
