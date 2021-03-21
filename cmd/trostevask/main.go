package main

import (
	"github.com/markusranda/trostevask/pkg/filemanager"
	"github.com/markusranda/trostevask/pkg/renamer"
)


func main() {
	println("---------------------Started trostevask---------------------")
	println("")

	disposeOldFiles()

	setupTestEnvironment()

	println("Read all dirty files")
	filemanager.ReadAndPrintFiles("./test_files/dirty")
	println("")

	// Move dirty files to processing
	println("Moving dirty files to processing")
	println("")
	moveDirtyFiles()

	println("Read all processing files")
	filemanager.ReadAndPrintFiles("./test_files/processing")
	println("")

	// Rename and move processed files
	println("Cleaning up filenames")
	println("")
	cleanFilenames()

	println("Read all processed files")
	filemanager.ReadAndPrintFiles("./test_files/processing")
	println("")

	println("Read all cleaned files")
	filemanager.ReadAndPrintFiles("./test_files/clean")
	println("")

	disposeOldFiles()
}

func setupTestEnvironment() {
	var testFileNames = []string{"Pirates of the carribEan 2003 blurAy xHaxx0r.mp4",
		"jason_borne_identity_2002 CONTRIBUTE TO MOVIEMASTERS.mp4",
		"the great escape (1963).mkv",
		"Black Swan 2010 dddddd.mkv",
		"2001.A.Space.Odyssey.1968.720p.BluRay.DD5.1.x264-LiNG.mkv"}
	generateTestFiles(testFileNames)
}


func generateTestFiles(testFileNames []string) {
	for i := 0; i < len(testFileNames); i++ {
		filemanager.CreateFile("./test_files/dirty/", testFileNames[i])
	}
}

func cleanFilenames() {
	var fileNameList = filemanager.GetFileNamesFromDir("./test_files/processing/")

	basePathProcessing := "./test_files/processing/"
	basePathClean := "./test_files/clean/"

	for _, filename := range fileNameList {
		println("Cleaning file: " + filename)
		var oldFilename = basePathProcessing + filename
		var newFilename = basePathClean + renamer.GetCleanFilename(filename)
		filemanager.MoveFile(oldFilename, newFilename)
	}
}

func moveDirtyFiles() {
	var fileNameList = filemanager.GetFileNamesFromDir("./test_files/dirty/")
	
    for _, filename := range fileNameList {
		var oldFilename = "./test_files/dirty/" + filename
		var newFilename = "./test_files/processing/" + filename
		filemanager.MoveFile(oldFilename, newFilename)	
	}
}


func disposeOldFiles() {
	filemanager.RemoveContents("./test_files/dirty/")
	filemanager.RemoveContents("./test_files/processing/")
	filemanager.RemoveContents("./test_files/clean/")
}