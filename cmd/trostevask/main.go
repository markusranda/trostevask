package main

import (
	"github.com/markusranda/trostevask/pkg/cleaner"
	"github.com/markusranda/trostevask/pkg/filemanager"
)

func main() {
	println("---------------------Started trostevask---------------------")
	println("")

	disposeFiles()

	setupTestEnvironment()

	println("Read all dirty files")
	filemanager.ReadAndPrintFiles("./test_files/dirty")
	println("")

	println("Cleaning up filenames")
	println("")
	cleanFilenames()

	println("Read all cleaned files")
	println("")
	filemanager.ReadAndPrintFiles("./test_files/clean")
	println("")

	disposeFiles()
}

func setupTestEnvironment() {
	movieTestFiles := []string{
		"Pirates of the carribEan 2003 blurAy xHaxx0r.mp4",
		"jason_borne_identity_2002 CONTRIBUTE TO MOVIEMASTERS.mp4",
		"the great escape (1963).mkv",
		"Black Swan 2010 dddddd.mkv",
		"2001.A.Space.Odyssey.1968.720p.BluRay.DD5.1.x264-LiNG.mkv",
	}
	generateMovieTestFiles(movieTestFiles)

	tvShowTestFiles := []string{
		"Bobs.Burgers.S04E12.WEB.x264-PHOENiX[TGx]",
		"Stargate SG1 Complete 1997-2007 DVD Rip x264 AC3-MEECH",
		"Stargate Atlantis S01-05 BR 10bit ace hevc-",
		"Batwoman.S02E07.XviD-AFG[TGx]",
		"Bleach 1-366 + Movies Batch Complete",
		"Bleach Season 13 English Dubbed HD",
	}
	generateTvShowTestFiles(tvShowTestFiles)

	filemanager.CreateDir("./test_files/clean/movies", 0755)
	filemanager.CreateDir("./test_files/clean/tv_shows", 0755)
}

func generateTvShowTestFiles(testFiles []string) {
	for i := 0; i < len(testFiles); i++ {
		filemanager.CreateDir("./test_files/dirty/" + testFiles[i], 0755)
	}
}

func generateMovieTestFiles(testFiles []string) {
	for i := 0; i < len(testFiles); i++ {
		filemanager.CreateFile("./test_files/dirty/", testFiles[i])
	}
}

func cleanFilenames() {
	var fileNameList = filemanager.GetFileNamesFromDir("./test_files/dirty/")

	basePathProcessing := "./test_files/dirty/"
	basePathClean := "./test_files/clean/"

	for _, filename := range fileNameList {

		println("Cleaning file: " + filename)

		cleanFilename := cleaner.GetCleanFilename(filename)

		filemanager.MoveFile(basePathProcessing + filename, basePathClean + cleanFilename)
	}
}

func disposeFiles() {
	println("Disposing all files")
	filemanager.RemoveContents("./test_files/dirty/")
	filemanager.RemoveContents("./test_files/clean/")
}