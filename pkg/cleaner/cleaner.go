package cleaner

import (
	"regexp"

	"github.com/markusranda/trostevask/pkg/renamer"
)

func GetCleanFilename(filename string) (cleanFilename string) {

	if (isTvShowFileName(filename)) {
		cleanFilename = "tv_shows/" + renamer.CleanTvShowName(filename)
	} else {
		cleanFilename = "movies/" + renamer.CleanMovieName(filename)
	}

	return
}

func isTvShowFileName(filename string) (isTvShow bool) {
	regex := regexp.MustCompile(`(?i:complete|S0\d|Season)`)
	return regex.MatchString(filename)
}
