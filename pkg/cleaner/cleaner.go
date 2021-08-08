package cleaner

import (
	uuid2 "github.com/google/uuid"
	"github.com/markusranda/trostevask/pkg/filemanager"
	"github.com/markusranda/trostevask/pkg/renamer"
	"github.com/sirupsen/logrus"
	"regexp"
)

func CleanAllFilenames() {
	var fileList = filemanager.GetFilesFromDirRecursive(filemanager.GetInputDir())

	for _, file := range fileList {
		CleanFileName(file)
	}
}

func CleanFileName(file filemanager.FullFileInfo) {
	uuid := uuid2.New().String()
	if shouldSkipFile(file) {
		if file.FileInfo != nil {
			logrus.Debug("Skipping file: " + file.Name())
		} else {
			logrus.Debug("Skipping file corrupt file")
		}
		return
	}

	logrus.Infof("[%s] Cleaning file: %s", uuid, file.Name())
	cleanFile := renamer.GetCleanFilename(file)

	if filemanager.FileExists(filemanager.GetOutputDir() + cleanFile.Path) {
		logrus.Infof("[%s] Skipping, file already exists", uuid)
		return
	}

	logrus.Infof("[%s] Copying file: %s", uuid, cleanFile.Path)
	err := filemanager.CopyFile(file.Path, filemanager.GetOutputDir()+cleanFile.Path)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("[%s] File copy complete!", uuid)
}

func shouldSkipFile(file filemanager.FullFileInfo) bool {
	if file.FileInfo == nil {
		logrus.Error("File is missing FullFileInfo field")
		return true
	}

	return file.Name() == "dirty" || IsNotValid(file) || filemanager.IsFolder(file.Path)
}

func IsNotValid(file filemanager.FullFileInfo) bool {
	re := regexp.MustCompile(`(?m).+mp4|avi|mkv`)
	validFileExtensionExists := re.MatchString(file.Name())
	if !validFileExtensionExists {
		return true
	}

	return false
}
