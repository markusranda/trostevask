package renamer

import (
	"fmt"
	"strings"
)

func rename(filename string) string {
	fmt.Print(filename)
	formattedString := strings.ToLower(filename)
	return formattedString
}
