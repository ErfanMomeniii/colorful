package main

import (
	"fmt"
	"github.com/ErfanMomeniii/colorful"
)

func main() {

	fmt.Print(colorful.Sprint(colorful.BlackColor, colorful.BlueBackground, " success "))
	colorful.Print(colorful.BlackColor, colorful.GreenBackground, " success ")
}
