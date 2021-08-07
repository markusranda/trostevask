package printer

type Color string

const (
    RED Color = "\033[31m"
    GREEN Color = "\033[32m"
    RESET Color = "\033[0m"
    YELLOW Color = "\033[33m"
	BLUE Color = "\033[34m"
	PURPLE Color = "\033[35m"
	CYAN Color = "\033[36m"
	WHITE Color = "\033[37m"
)


func PrintColoredText(text string, colorString Color) {
	println(string(colorString))
	println(text)
	println(string(RESET))
}