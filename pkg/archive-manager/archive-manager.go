package archive_manager

import (
	"github.com/markusranda/trostevask/pkg/filemanager"
	"github.com/mholt/archiver/v3"
	"github.com/sirupsen/logrus"
	"os"
	"regexp"
)

const ArchiveRegex = `(?m)(.+?)\.(rar|tar\.gz|zip|tar\.|tar$)`

func GetMediaFileFromCompressed(file filemanager.FullFileInfo) filemanager.FullFileInfo {
	fileNameWithoutSuffix := getFileNameWithoutExtension(file)

	var mediaFile filemanager.FullFileInfo

	err := archiver.Unarchive(file.Path, filemanager.GetInputDir()+fileNameWithoutSuffix)
	if err != nil {
		logrus.Info(err)
	}

	fileNames := filemanager.GetFileNamesFromDirRecursive(filemanager.GetInputDir() + fileNameWithoutSuffix)

	// Find the media file
	for _, fileName := range fileNames {
		logrus.Debugf("File: %s was found", fileName)

		var tempFile filemanager.FullFileInfo
		tempFile.FileInfo, err = os.Stat(filemanager.GetInputDir() + fileNameWithoutSuffix + "/" + fileName)
		tempFile.Path = filemanager.GetInputDir() + fileNameWithoutSuffix + "/" + fileName
		if err != nil {
			logrus.Debugf("Can't Stat file: %s, this file is probably a dir", fileName)
			continue
		}

		if IsMediaFile(tempFile) {
			mediaFile.FileInfo = tempFile.FileInfo
			mediaFile.Path = filemanager.GetInputDir() + fileNameWithoutSuffix + "/" + fileName
			continue
		} else {
			logrus.Debugf("File is not media file: %s", tempFile.Name())
		}
	}

	logrus.Infof("Returning: %s, to cleaning", mediaFile.Name())

	return mediaFile
}

func IsMediaFile(file filemanager.FullFileInfo) bool {
	re := regexp.MustCompile(`(?m).+mp4|avi|mkv`)
	validFileExtensionExists := re.MatchString(file.Name())
	if !validFileExtensionExists {
		return false
	}

	return true
}

func IsArchive(file filemanager.FullFileInfo) bool {
	var re = regexp.MustCompile(ArchiveRegex)
	return re.MatchString(file.Name())
}

func GetFileExtension(file filemanager.FullFileInfo) (extension string) {
	var re = regexp.MustCompile(ArchiveRegex)
	extension = re.ReplaceAllString(file.Name(), "$2")
	return
}

func getFileNameWithoutExtension(file filemanager.FullFileInfo) (extension string) {
	var re = regexp.MustCompile(ArchiveRegex)
	extension = re.ReplaceAllString(file.Name(), "$1")
	return
}
