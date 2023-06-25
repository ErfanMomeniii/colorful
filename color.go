package colorful

import (
	"fmt"
)

// Color
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

// Background color
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

type Pattern struct {
	Color           string
	BackgroundColor string
}

func WithColor(color string) *Pattern {
	return &Pattern{
		Color:           color,
		BackgroundColor: DefaultBackgroundColor,
	}
}

func WithBackgroundColor(backgroundColor string) *Pattern {
	return &Pattern{
		Color:           DefaultColor,
		BackgroundColor: backgroundColor,
	}
}

func (p *Pattern) WithColor(color string) *Pattern {
	p.Color = color

	return p
}

func (p *Pattern) WithBackgroundColor(backgroundColor string) *Pattern {
	p.BackgroundColor = backgroundColor

	return p
}

func Print(text ...any) {
	fmt.Print(text...)
}

func (p *Pattern) Print(text ...any) {
	fmt.Print(p.Color, p.BackgroundColor, fmt.Sprint(text...), ResetColor)
}

func Println(text ...any) {
	fmt.Println(text...)
}

func (p *Pattern) Println(text ...any) {
	fmt.Println(p.Color, p.BackgroundColor, fmt.Sprint(text...), ResetColor)
}

func Printf(format string, params ...any) {
	fmt.Printf(format, params...)
}

func (p *Pattern) Printf(format string, params ...any) {
	fmt.Print(p.Color, p.BackgroundColor, fmt.Sprintf(format, params...), ResetColor)
}

func Sprint(text ...any) string {
	return fmt.Sprint(text...)
}

func (p *Pattern) Sprint(text ...any) string {
	return p.Color + p.BackgroundColor + fmt.Sprint(text...) + ResetColor
}

func Sprintln(text ...any) string {
	return fmt.Sprintln(text...)
}

func (p *Pattern) Sprintln(text ...any) string {
	return p.Color + p.BackgroundColor + fmt.Sprint(text...) + ResetColor + "\n"
}

func Sprintf(format string, params ...any) string {
	return fmt.Sprintf(format, params...)
}

func (p *Pattern) Sprintf(format string, params ...any) string {
	return p.Color + p.BackgroundColor + fmt.Sprintf(format, params...) + ResetColor
}

func GetColorFromCode(code int) string {
	return fmt.Sprintf("\u001b[38;5;%dm", code)
}

func GetBackgroundColorFromCode(code int) string {
	return fmt.Sprintf("\u001b[48;5;%dm", code)
}
