package main

import (
	"github.com/markusranda/trostevask/pkg/cleaner"
	"github.com/markusranda/trostevask/pkg/filemanager"
	"github.com/markusranda/trostevask/pkg/printer"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	printer.PrintColoredText("---------------------Started trostevask---------------------", printer.YELLOW)

	argsWithProg := os.Args
	if len(argsWithProg) > 1 {
		for _, arg := range argsWithProg {

			if arg == "dispose" {
				disposeFiles()
			}

			if arg == "test" {
				setupTestEnvironment()
			}
		}
	}

	printer.PrintColoredText("Read all dirty files", printer.YELLOW)
	filemanager.ReadAndPrintAllFiles("./test_files/dirty")

	printer.PrintColoredText("Cleaning up filenames", printer.YELLOW)
	cleanFilenames()
}

func cleanFilenames() {
	baseDir := "./test_files/"
	basePathProcessing := baseDir + "dirty/"
	basePathClean := baseDir + "clean/"
	basePathRejected := baseDir + "rejected/"

	var fileList = filemanager.GetFilesFromDirRecursive(basePathProcessing)

	for _, file := range fileList {

		// Skipping dirs
		if filemanager.IsFolder(file.Path) {
			log.Debug("Skipping file: " + file.Name())
			continue
		}

		log.Info("Cleaning file: " + file.Name())

		if cleaner.IsNotValidated(file) {
			rejectFile(basePathRejected, file)
			continue
		}

		cleanFile := cleaner.GetCleanFilename(file)

		if cleanFile.Name() == "" {
			log.Error("Got empty cleaned file for: " + file.Name())
			rejectFile(basePathRejected, file)
			continue
		}

		log.Info("Moving file: " + cleanFile.Path)
		err := filemanager.CopyFile(file.Path, basePathClean+cleanFile.Path)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func rejectFile(rejectDir string, file filemanager.FullFileInfo) {
	log.Error("Moving file to rejected: " + file.Name())
	err := filemanager.CopyFile(file.Path, rejectDir+file.Name())
	if err != nil {
		log.Fatal(err)
	}
}

func disposeFiles() {
	log.Info("Disposing all files")
	err := filemanager.RemoveContents("./test_files/dirty/")
	if err != nil {
		log.Fatal(err)
	}
	err = filemanager.RemoveContents("./test_files/clean/")
	if err != nil {
		log.Fatal(err)
	}
	err = filemanager.RemoveContents("./test_files/rejected/")
	if err != nil {
		log.Fatal(err)
	}
}
