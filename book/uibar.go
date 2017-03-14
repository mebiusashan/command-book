package main

import (
	"github.com/gizak/termui"
)

func initBar() {
	Bar = termui.NewPar(barText)

	Bar.Height = 1
	Bar.Width = termui.TermWidth()

	Bar.TextBgColor = termui.ColorWhite
	Bar.TextFgColor = termui.ColorBlack

	Bar.Border = false

	updateBar()
}

func updateBar() {
	Bar.Y = termui.TermHeight() - 1
}
