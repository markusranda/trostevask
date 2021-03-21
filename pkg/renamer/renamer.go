package renamer

import (
	"regexp"
	"strings"
)


func CleanTvShowName(filename string) (cleanFilename string) {

	return filename
}

func CleanMovieName(filename string) (cleanFilename string) {
	// Remove everything after year
	var re = regexp.MustCompile(`^(.+).(\d{4}).+.(mp4|avi|mkv)`)
	cleanFilename = re.ReplaceAllString(filename, `$1.($2).$3`)

	// Make string all lower case
	cleanFilename = strings.ToLower(cleanFilename)

	// Replace spaces with periods
	cleanFilename = strings.ReplaceAll(cleanFilename, " ", ".")
	cleanFilename = strings.ReplaceAll(cleanFilename, "_", ".")
	cleanFilename = strings.ReplaceAll(cleanFilename, "-", ".")
	cleanFilename = strings.ReplaceAll(cleanFilename, "..", ".")

	return
}
