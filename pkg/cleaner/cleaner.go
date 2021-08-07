package cleaner

import (
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"

	"github.com/markusranda/trostevask/pkg/filemanager"
	"github.com/markusranda/trostevask/pkg/renamer"
)

func GetCleanFilename(file filemanager.FullFileInfo) (cleanFile filemanager.FullFileInfo) {
	cleanFile = file

	if isTvShowFileName(file.Name()) {
		// Handle tv show
		tvShowName := grabName(file.Name())
		if tvShowName == "" {
			// If no tv show tvShowName found return empty
			return
		}
		filemanager.CreateDirIfNotExists("./test_files/clean/tv_shows/"+tvShowName, 0755)

		// Handle season
		tvShowSeason := grabSeason(file.Name())
		filemanager.CreateDirIfNotExists("./test_files/clean/tv_shows/"+tvShowName+"/season_"+tvShowSeason, 0755)

		// Handle full name
		cleanFile.Path = "tv_shows/" + tvShowName + "/season_" + tvShowSeason + "/" + renamer.CleanTvShowName(file.Name())

	} else {
		cleanFile.Path = "movies/" + renamer.CleanMovieName(file.Name())
	}

	return
}

func IsNotValidated(file filemanager.FullFileInfo) bool {
	re := regexp.MustCompile(`(?m).+mp4|avi|mkv`)
	fileExtensionExists := re.MatchString(file.Name())
	if !fileExtensionExists {
		logrus.Error("No file extension found for: " + file.Name())
		return true
	}

	return false
}

func isTvShowFileName(filename string) (isTvShow bool) {
	regex := regexp.MustCompile(`(?i:complete|S0\d|Season)`)
	return regex.MatchString(filename)
}

func grabName(filename string) (name string) {
	name = strings.ToLower(filename)

	var re = regexp.MustCompile(`(?mi)(?P<name>.*?)(.season|.s\d{2})`)

	if len(re.FindStringSubmatch(name)) < 1 {
		name = ""
	} else {
		name = re.FindStringSubmatch(name)[1]
	}

	return name
}

func grabSeason(filename string) (seasonNum string) {
	seasonNum = strings.ToLower(filename)
	var re = regexp.MustCompile(`[sS](\d{1,2})`)
	seasonNum = re.FindStringSubmatch(seasonNum)[1]

	// Handle situations where num is 1, and we want to display 01
	seasonNumInt, err := strconv.Atoi(seasonNum)
	if err != nil {
		logrus.Fatal(err)
	}

	if seasonNumInt < 10 && len([]rune(seasonNum)) < 2 {
		seasonNum = "0" + seasonNum
	}

	return seasonNum
}
