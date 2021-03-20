package main

import (
	"fmt"

	"github.com/markusranda/trostevask/pkg/filemanager"
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
	moveDirtyFiles()

	// Rename all files in processing

	disposeOldFiles()
}

func setupTestEnvironment() {
	var testFileNames = []string{"Pirates of the carribEan 2003 blurAy xHaxx0r.mp4",
		"jason_borne_identity_2002 CONTRIBUTE TO MOVIEMASTERS.mp4",
		"the great escape (1963).mkv",
		"Black Swan 2010 dddddd.mkv"}
	generateTestFiles(testFileNames)
}


func generateTestFiles(testFileNames []string) {
	for i := 0; i < len(testFileNames); i++ {
		filemanager.CreateFile("./test_files/dirty/", testFileNames[i])
	}
}

func moveDirtyFiles() {
	var fileNameList = filemanager.GetRelativeFileNamesFromDir("./test_files/dirty/")
	
	fmt.Printf("len=%d cap=%d %v\n", len(fileNameList), cap(fileNameList), fileNameList)

	// filemanager.MoveFile()
}


func disposeOldFiles() {
	filemanager.RemoveContents("./test_files/dirty/")
	filemanager.RemoveContents("./test_files/processing/")
	filemanager.RemoveContents("./test_files/clean/")
}