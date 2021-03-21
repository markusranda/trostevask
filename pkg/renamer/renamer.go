package renamer

import (
	"errors"
	"regexp"
	"strings"

	"github.com/markusranda/trostevask/pkg/consts"
)

func GetCleanFilename(filename string, fileType string) (cleanFilename string, err error) {

	switch fileType {

	case consts.TvShow:
		cleanFilename = cleanTvShowName(filename)

	case consts.Movie:
		cleanFilename = cleanMovieName(filename)

	default:
		err = errors.New(`None of the applicable fileTypes were specified, ` + fileType + " where used instead..")
	}

	return
}

func cleanTvShowName(filename string) (cleanFilename string) {

	return filename
}

func cleanMovieName(filename string) (cleanFilename string) {
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

func IsTvShowFileName(filename string) (isTvShow bool) {
	regex := regexp.MustCompile(`(?i:complete|S0\d|Season)`)
	return regex.MatchString(filename)
}

func removeAllAfterYear(filename string) (cleanFilename string) {
	var re = regexp.MustCompile(`^(.+).(\d{4}).+.(mp4|avi|mkv)`)
	cleanFilename = re.ReplaceAllString(filename, `$1.($2).$3`)
	return
}