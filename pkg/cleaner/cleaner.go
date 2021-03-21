package cleaner

import (
	"errors"

	"github.com/markusranda/trostevask/pkg/consts"
	"github.com/markusranda/trostevask/pkg/renamer"
)

func GetCleanFilename(filename string, fileType string) (cleanFilename string, err error) {

	switch fileType {

	case consts.TvShow:
		cleanFilename = renamer.CleanTvShowName(filename)

	case consts.Movie:
		cleanFilename = renamer.CleanMovieName(filename)

	default:
		err = errors.New(`None of the applicable fileTypes were specified, ` + fileType + " where used instead..")
	}

	return
}
