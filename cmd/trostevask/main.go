package main

import (
	"github.com/joho/godotenv"
	"github.com/markusranda/trostevask/pkg/cleaner"
	"github.com/markusranda/trostevask/pkg/filemanager"
	"github.com/markusranda/trostevask/pkg/sniffer"
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

	log.Info("Doing initial scan of incoming dir, and running cleanup")
	log.Info("Cleaning up filenames...")
	cleaner.CleanAllFilenames()

	log.Info("Start sniffing for file changes..")
	done := make(chan bool)
	go sniffer.Sniff()

	<-done
}

func handleArguments() {
	arguments := os.Args[1:]
	for _, arg := range arguments {
		if arg == "dispose" {
			disposeFiles()
		}
	}

	for _, arg := range arguments {
		if arg == "test" {
			setupTestEnvironment()
		}
		if arg == "debug" {
			// parse string, this is built-in feature of logrus
			ll, err := log.ParseLevel("debug")
			if err != nil {
				ll = log.DebugLevel
			}
			// set global log level
			log.SetLevel(ll)
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

func disposeFiles() {
	log.Info("Disposing all files")
	err := filemanager.RemoveContents(filemanager.GetInputDir())
	if err != nil {
		log.Fatal(err)
	}
	err = filemanager.RemoveContents(filemanager.GetOutputDir() + "movies/")
	err = filemanager.RemoveContents(filemanager.GetOutputDir() + "tv_shows/")
	if err != nil {
		log.Fatal(err)
	}
	err = filemanager.RemoveContents(filemanager.GetRejectedDir())
	if err != nil {
		log.Fatal(err)
	}
}
