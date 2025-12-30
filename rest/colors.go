package rest

func RGB(r, g, b uint8) [3]uint8 {
	return [3]uint8{r, g, b}
}

var (
	White       = RGB(255, 255, 255)
	Red         = RGB(255, 0, 0)
	Green       = RGB(0, 255, 0)
	Blue        = RGB(0, 0, 255)
	Orange      = RGB(255, 166, 64)
	Yellow      = RGB(255, 255, 0)
	Cyan        = RGB(0, 255, 255)
	Magenta     = RGB(255, 0, 255)
	Purple      = RGB(128, 0, 128)
	Pink        = RGB(255, 192, 203)
	Lime        = RGB(0, 255, 128)
	Turquoise   = RGB(64, 224, 208)
	Teal        = RGB(0, 128, 128)
	Olive       = RGB(128, 128, 0)
	Maroon      = RGB(128, 0, 0)
	Navy        = RGB(0, 0, 128)
	Indigo      = RGB(75, 0, 130)
	Violet      = RGB(238, 130, 238)
	Coral       = RGB(255, 127, 80)
	Gold        = RGB(255, 215, 0)
	Salmon      = RGB(250, 128, 114)
	SkyBlue     = RGB(135, 206, 235)
	Lavender    = RGB(230, 230, 250)
	Chocolate   = RGB(210, 105, 30)
	Crimson     = RGB(220, 20, 60)
	SpringGreen = RGB(0, 255, 127)
	Aqua        = RGB(0, 255, 255)
	DeepPink    = RGB(255, 20, 147)
)
