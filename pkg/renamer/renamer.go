package renamer

import (
	"regexp"
	"strings"
)

func CleanTvShowName(filename string) (cleanFilename string) {
	var re = regexp.MustCompile(`^(.+).([sS]\d{1,2}[eE]\d{1,2}).+(mp4|avi|mkv)`)
	cleanFilename = re.ReplaceAllString(filename, `$1.$2.$3`)
	return
}

func CleanMovieName(filename string) (cleanFilename string) {
	// Remove everything after year
	var re = regexp.MustCompile(`^(.+).(19\d{2}|20\d{2}).+.(mp4|avi|mkv)`)
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
