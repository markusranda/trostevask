package renamer

import "strings"

func rename(filename string) string {
	formattedString := strings.ToLower(filename)
	return formattedString
}
