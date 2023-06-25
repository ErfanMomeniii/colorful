package colorful

import (
	"fmt"
)

// Colors
const (
	DefaultColor       = ""
	BlackColor         = "\u001b[30m"
	RedColor           = "\u001b[31m"
	GreenColor         = "\u001b[32m"
	YellowColor        = "\u001b[33m"
	BlueColor          = "\u001b[34m"
	MagentaColor       = "\u001b[35m"
	CyanColor          = "\u001b[36m"
	WhiteColor         = "\u001b[37m"
	BrightBlackColor   = "\u001b[30;1m"
	BrightRedColor     = "\u001b[31;1m"
	BrightGreenColor   = "\u001b[32;1m"
	BrightYellowColor  = "\u001b[33;1m"
	BrightBlueColor    = "\u001b[34;1m"
	BrightMagentaColor = "\u001b[35;1m"
	BrightCyanColor    = "\u001b[36;1m"
	BrightWhiteColor   = "\u001b[37;1m"
	ResetColor         = "\u001b[0m"
)

// Background colors
const (
	DefaultBackgroundColor       = ""
	BlackBackgroundColor         = "\u001b[40m"
	RedBackgroundColor           = "\u001b[41m"
	GreenBackgroundColor         = "\u001b[42m"
	YellowBackgroundColor        = "\u001b[43m"
	BlueBackgroundColor          = "\u001b[44m"
	MagentaBackgroundColor       = "\u001b[45m"
	CyanBackgroundColor          = "\u001b[46m"
	WhiteBackgroundColor         = "\u001b[47m"
	BrightBlackBackgroundColor   = "\u001b[40;1m"
	BrightRedBackgroundColor     = "\u001b[41;1m"
	BrightGreenBackgroundColor   = "\u001b[42;1m"
	BrightYellowBackgroundColor  = "\u001b[43;1m"
	BrightBlueBackgroundColor    = "\u001b[44;1m"
	BrightMagentaBackgroundColor = "\u001b[45;1m"
	BrightCyanBackgroundColor    = "\u001b[46;1m"
	BrightWhiteBackgroundColor   = "\u001b[47;1m"
)

// Pattern is an instantiation that defines color and background color.
type Pattern struct {
	Color           string
	BackgroundColor string
}

// WithColor returns pattern with inputted color
func WithColor(color string) *Pattern {
	return &Pattern{
		Color:           color,
		BackgroundColor: DefaultBackgroundColor,
	}
}

// WithBackgroundColor returns pattern with inputted background color.
func WithBackgroundColor(backgroundColor string) *Pattern {
	return &Pattern{
		Color:           DefaultColor,
		BackgroundColor: backgroundColor,
	}
}

// WithColor returns patterns that changed color with inputted color.
func (p *Pattern) WithColor(color string) *Pattern {
	p.Color = color

	return p
}

// WithBackgroundColor returns patterns that changed color with inputted background color.
func (p *Pattern) WithBackgroundColor(backgroundColor string) *Pattern {
	p.BackgroundColor = backgroundColor

	return p
}

// Print prints inputted text.
func Print(text ...any) *Pattern {
	fmt.Print(text...)

	return &Pattern{
		Color:           DefaultColor,
		BackgroundColor: DefaultBackgroundColor,
	}
}

// Print prints inputted text formats with the color and background that set in the pattern.
func (p *Pattern) Print(text ...any) *Pattern {
	fmt.Print(p.Color, p.BackgroundColor, fmt.Sprint(text...), ResetColor)

	return p
}

// Println prints text and ends with newline.
func Println(text ...any) *Pattern {
	fmt.Println(text...)

	return &Pattern{
		Color:           DefaultColor,
		BackgroundColor: DefaultBackgroundColor,
	}
}

// Println prints text formats with the color and background that set in the pattern and ends with newline.
func (p *Pattern) Println(text ...any) *Pattern {
	fmt.Println(p.Color, p.BackgroundColor, fmt.Sprint(text...), ResetColor)

	return p
}

// Printf prints text formats with the inputted format.
func Printf(format string, params ...any) *Pattern {
	fmt.Printf(format, params...)

	return &Pattern{
		Color:           DefaultColor,
		BackgroundColor: DefaultBackgroundColor,
	}
}

// Printf prints text formats with the inputted format ,color and background that set in the pattern.
func (p *Pattern) Printf(format string, params ...any) *Pattern {
	fmt.Print(p.Color, p.BackgroundColor, fmt.Sprintf(format, params...), ResetColor)

	return p
}

// Sprint returns string .
func Sprint(text ...any) string {
	return fmt.Sprint(text...)
}

// Sprint returns string that formats with color and background color that set in the pattern.
func (p *Pattern) Sprint(text ...any) string {
	return p.Color + p.BackgroundColor + fmt.Sprint(text...) + ResetColor
}

// Sprintln returns string that formats with the inputted format and ends with new line.
func Sprintln(text ...any) string {
	return fmt.Sprintln(text...)
}

// Sprintln returns string that formats with the inputted format and color and background color set in the pattern and ends with new line.
func (p *Pattern) Sprintln(text ...any) string {
	return p.Color + p.BackgroundColor + fmt.Sprint(text...) + ResetColor + "\n"
}

// Sprintf returns string that formats with the inputted format.
func Sprintf(format string, params ...any) string {
	return fmt.Sprintf(format, params...)
}

// Sprintf returns string that formats with the inputted format and color and background that set in the pattern.
func (p *Pattern) Sprintf(format string, params ...any) string {
	return p.Color + p.BackgroundColor + fmt.Sprintf(format, params...) + ResetColor
}

// Color returns color with the inputted code.
func Color(colorCode int) string {
	return fmt.Sprintf("\u001b[38;5;%dm", colorCode)
}

// BackgroundColor returns background color with the inputted code.
func BackgroundColor(backgroundcolorCode int) string {
	return fmt.Sprintf("\u001b[48;5;%dm", backgroundcolorCode)
}
