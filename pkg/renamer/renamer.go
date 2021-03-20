package renamer

import (
	"regexp"
	"strings"
)

func GetCleanFilename(filename string) (cleanFilename string) {
	// Remove everything after year
	cleanFilename = removeAllAfterYear(filename)
	
	// Make string all lower case
	cleanFilename = strings.ToLower(cleanFilename)

	// Replace spaces with periods
	cleanFilename = strings.ReplaceAll(cleanFilename, " ", ".")
	cleanFilename = strings.ReplaceAll(cleanFilename, "_", ".")
	cleanFilename = strings.ReplaceAll(cleanFilename, "-", ".")
	cleanFilename = strings.ReplaceAll(cleanFilename, "..", ".")

	return
}

func removeAllAfterYear(filename string) (cleanFilename string) {
	var re = regexp.MustCompile(`^(.+).(\d{4}).+.(mp4|avi|mkv)`)
	cleanFilename = re.ReplaceAllString(filename, `$1.($2).$3`)
	return
}