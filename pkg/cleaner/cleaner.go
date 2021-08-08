package cleaner

import (
	"github.com/markusranda/trostevask/pkg/filemanager"
	"github.com/markusranda/trostevask/pkg/renamer"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
)

func CleanFilenames() {
	var fileList = filemanager.GetFilesFromDirRecursive(filemanager.GetInputDir())

	var skipped int
	var transferred int

	for _, file := range fileList {

		if shouldSkipFile(file) {
			logrus.Debug("Skipping file: " + file.Name())
			skipped++
			continue
		}

		logrus.Info("Cleaning file: " + file.Name())
		cleanFile := renamer.GetCleanFilename(file)

		if filemanager.FileExists(filemanager.GetOutputDir() + cleanFile.Path) {
			logrus.Info("Skipping, file already exists")
			continue
		}

		logrus.Info("Copying file: " + cleanFile.Path)
		err := filemanager.CopyFile(file.Path, filemanager.GetOutputDir()+cleanFile.Path)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	logrus.Info("Skipped " + strconv.Itoa(skipped) + " of " + strconv.Itoa(len(fileList)) + " files.")
	logrus.Info("Transferred " + strconv.Itoa(transferred) + " of " + strconv.Itoa(len(fileList)) + " files.")
}

func shouldSkipFile(file filemanager.FullFileInfo) bool {
	return IsNotValid(file) || filemanager.IsFolder(file.Path)
}

func IsNotValid(file filemanager.FullFileInfo) bool {
	re := regexp.MustCompile(`(?m).+mp4|avi|mkv`)
	validFileExtensionExists := re.MatchString(file.Name())
	if !validFileExtensionExists {
		return true
	}

	return false
}
