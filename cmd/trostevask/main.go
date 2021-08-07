package main

import (
	"github.com/joho/godotenv"
	"github.com/markusranda/trostevask/pkg/cleaner"
	"github.com/markusranda/trostevask/pkg/filemanager"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	initLogger()
	log.Info("---------------------Started trostevask---------------------")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	initBaseDirs()

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

	log.Info("Read all dirty files")
	filemanager.ReadAndPrintAllFiles(filemanager.GetInputDir())

	log.Info("Cleaning up filenames")
	cleanFilenames()
}

func initLogger() {
	log.SetOutput(os.Stdout)
}

func initBaseDirs() {
	log.Info("Making sure all base directories are in order")
	inputDir := filemanager.GetInputDir()
	outputDir := filemanager.GetOutputDir()
	rejectedDir := filemanager.GetRejectedtDir()

	filemanager.CreateDirIfNotExists(inputDir, 0777)
	filemanager.CreateDirIfNotExists(outputDir, 0777)
	filemanager.CreateDirIfNotExists(rejectedDir, 0777)

	filemanager.CreateDirSkipIfExists(outputDir+"movies", 0755)
	filemanager.CreateDirSkipIfExists(outputDir+"tv_shows", 0755)
}

func cleanFilenames() {
	var fileList = filemanager.GetFilesFromDirRecursive(filemanager.GetInputDir())

	for _, file := range fileList {

		// Skipping dirs
		if filemanager.IsFolder(file.Path) {
			log.Debug("Skipping file: " + file.Name())
			continue
		}

		if cleaner.IsNotValidated(file) {
			log.Info("Skipping file: " + file.Name())
			continue
		}

		log.Info("Cleaning file: " + file.Name())

		if cleaner.IsNotValidated(file) {
			rejectFile(file)
			continue
		}

		cleanFile := cleaner.GetCleanFilename(file)

		if cleanFile.Name() == "" {
			log.Error("Got empty cleaned file for: " + file.Name())
			rejectFile(file)
			continue
		}

		if filemanager.FileExists(filemanager.GetOutputDir() + cleanFile.Path) {
			log.Info("Skipping, file already exists")
			continue
		}

		log.Info("Copying file: " + cleanFile.Path)
		err := filemanager.CopyFile(file.Path, filemanager.GetOutputDir()+cleanFile.Path)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func rejectFile(file filemanager.FullFileInfo) {
	log.Error("Moving file to rejected: " + file.Name())
	err := filemanager.CopyFile(file.Path, filemanager.GetRejectedtDir()+file.Name())
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
