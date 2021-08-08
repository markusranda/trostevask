package renamer

import (
	"github.com/markusranda/trostevask/pkg/filemanager"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
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

func GetCleanFilename(file filemanager.FullFileInfo) (cleanFile filemanager.FullFileInfo) {
	cleanFile = file

	if IsTvShowFileName(file.Name()) {
		cleanFile.Path = handleTvShowName(file)
	} else {
		cleanFile.Path = handleMovieName(file)
	}

	return
}

func handleMovieName(file filemanager.FullFileInfo) (cleanFileName string) {
	cleanFileName = "movies/" + CleanMovieName(file.Name())
	return
}

func handleTvShowName(file filemanager.FullFileInfo) (cleanFileName string) {
	// Handle tv show
	tvShowName := GrabName(file.Name())
	if tvShowName == "" {
		// If no tv show tvShowName found return empty
		return
	}
	filemanager.CreateDirIfNotExists(filemanager.GetOutputDir()+"tv_shows/"+tvShowName, 0755)

	// Handle season
	tvShowSeason := GrabSeason(file.Name())
	filemanager.CreateDirIfNotExists(filemanager.GetOutputDir()+"tv_shows/"+tvShowName+"/season_"+tvShowSeason, 0755)

	// Handle full name
	cleanFileName = "tv_shows/" + tvShowName + "/season_" + tvShowSeason + "/" + CleanTvShowName(file.Name())
	return
}

func GrabName(filename string) (name string) {
	name = strings.ToLower(filename)

	var re = regexp.MustCompile(`(?mi)(?P<name>.*?)(.season|.s\d{2})`)

	if len(re.FindStringSubmatch(name)) < 1 {
		name = ""
	} else {
		name = re.FindStringSubmatch(name)[1]
	}

	return name
}

func GrabSeason(filename string) (seasonNum string) {
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

func IsTvShowFileName(filename string) (isTvShow bool) {
	regex := regexp.MustCompile(`(?i:complete|S0\d|Season)`)
	return regex.MatchString(filename)
}
