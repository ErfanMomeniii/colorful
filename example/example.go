package main

import (
	"github.com/erfanmomeniii/colorful"
)

func main() {
	// only with colorful color
	colorful.WithColor(colorful.BlackColor).Println("BlackColor")
	colorful.WithColor(colorful.RedColor).Println("RedColor")
	colorful.WithColor(colorful.GreenColor).Println("GreenColor")
	colorful.WithColor(colorful.YellowColor).Println("YellowColor")
	colorful.WithColor(colorful.BlueColor).Println("BlueColor")
	colorful.WithColor(colorful.MagentaColor).Println("MagentaColor")
	colorful.WithColor(colorful.CyanColor).Println("CyanColor")
	colorful.WithColor(colorful.WhiteColor).Println("WhiteColor")

	//only with colorful background color
	colorful.WithBackgroundColor(colorful.RedBackgroundColor).Println("RedBackgroundColor")
	colorful.WithBackgroundColor(colorful.GreenBackgroundColor).Println("GreenBackgroundColor")
	colorful.WithBackgroundColor(colorful.YellowBackgroundColor).Println("YellowBackgroundColor")
	colorful.WithBackgroundColor(colorful.BlueBackgroundColor).Println("BlueBackgroundColor")
	colorful.WithBackgroundColor(colorful.BlackBackgroundColor).Println("BlackBackgroundColor")

	// with coloful color and background color
	colorful.WithColor(colorful.RedColor).WithBackgroundColor(colorful.BlueBackgroundColor).Println("RedColor+BlueBackgroundColor")
	colorful.WithColor(colorful.YellowColor).WithBackgroundColor(colorful.GreenBackgroundColor).Println("YellowColor+GreenBackgroundColor")
	colorful.WithColor(colorful.BlackColor).WithBackgroundColor(colorful.WhiteBackgroundColor).Println("BlackColor+WhiteBackgroundColor")
	colorful.WithColor(colorful.WhiteColor).WithBackgroundColor(colorful.BlackBackgroundColor).Println("WhiteColor+BlackBackgroundColor")
}
