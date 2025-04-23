package myTheme

import (
	"com.cyy/sudoku/utils"
	"image/color"
)

func LineBorderColor() color.Color {
	return utils.HTML2FyneRGB(0, 0, 255)
}

func SimpleTextColor() color.Color {
	return utils.HTML2FyneRGB(0, 0, 85)
}

func BlueTextColor() color.Color {
	return utils.HTML2FyneRGB(0, 0, 178)
}
