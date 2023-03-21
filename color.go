package colorful

import "fmt"

// color
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

// background color
const (
	DefaultBackground       = ""
	BlackBackground         = "\u001b[40m"
	RedBackground           = "\u001b[41m"
	GreenBackground         = "\u001b[42m"
	YellowBackground        = "\u001b[43m"
	BlueBackground          = "\u001b[44m"
	MagentaBackground       = "\u001b[45m"
	CyanBackground          = "\u001b[46m"
	WhiteBackground         = "\u001b[47m"
	BrightBlackBackground   = "\u001b[40;1m"
	BrightRedBackground     = "\u001b[41;1m"
	BrightGreenBackground   = "\u001b[42;1m"
	BrightYellowBackground  = "\u001b[43;1m"
	BrightBlueBackground    = "\u001b[44;1m"
	BrightMagentaBackground = "\u001b[45;1m"
	BrightCyanBackground    = "\u001b[46;1m"
	BrightWhiteBackground   = "\u001b[47;1m"
)

func Print(color string, background string, text ...any) {
	fmt.Print(color, background, fmt.Sprint(text...), ResetColor)
}

func Println(color string, background string, text ...any) {
	fmt.Println(color, background, fmt.Sprint(text...), ResetColor)
}

func Printf(color string, background string, format string, params ...any) {
	fmt.Print(color, background, fmt.Sprintf(format, params...), ResetColor)
}

func Sprintln(color string, background string, text ...any) string {
	return color + background + fmt.Sprint(text...) + ResetColor + "\n"
}

func Sprint(color string, background string, text ...any) string {
	return color + background + fmt.Sprint(text...) + ResetColor
}

func Sprintf(color string, background string, format string, params ...any) string {
	return color + background + fmt.Sprintf(format, params...) + ResetColor
}

func GetColorFromCode(code int) string {
	return fmt.Sprintf("\u001b[38;5;%dm", code)
}

func GetBackgroundColorFromCode(code int) string {
	return fmt.Sprintf("\u001b[48;5;%dm", code)
}
