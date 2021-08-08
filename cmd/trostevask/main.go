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
	handleArguments()

	log.Info("Cleaning up filenames")
	cleaner.CleanFilenames()

	log.Info("All done!")
}

func handleArguments() {
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
}

func initLogger() {
	log.SetOutput(os.Stdout)
}

func initBaseDirs() {
	log.Info("Making sure all base directories are in order")
	inputDir := filemanager.GetInputDir()
	outputDir := filemanager.GetOutputDir()
	rejectedDir := filemanager.GetRejectedDir()

	filemanager.CreateDirIfNotExists(inputDir, 0777)
	filemanager.CreateDirIfNotExists(outputDir, 0777)
	filemanager.CreateDirIfNotExists(rejectedDir, 0777)

	filemanager.CreateDirSkipIfExists(outputDir+"movies", 0755)
	filemanager.CreateDirSkipIfExists(outputDir+"tv_shows", 0755)
}

func rejectFile(file filemanager.FullFileInfo) {
	log.Error("Moving file to rejected: " + file.Name())
	err := filemanager.CopyFile(file.Path, filemanager.GetRejectedDir()+file.Name())
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
