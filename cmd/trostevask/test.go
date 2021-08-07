package main

import (
	"github.com/markusranda/trostevask/pkg/filemanager"
	log "github.com/sirupsen/logrus"
)

func setupTestEnvironment() {
	log.Info("Setting up test environment")
	movieTestFiles := []string{
		"The.Favourite.2018.1080p.BluRay.x264.DTS-WiKi.mkv",
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
		"Bleach 1-366 + Movies Batch Complete",
		"Bleach Season 13 English Dubbed HD",
		"True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD",
	}
	generateTvShowTestFiles(tvShowTestFiles)

	filemanager.CreateDirSkipIfExists("./test_files/clean/movies", 0755)
	filemanager.CreateDirSkipIfExists("./test_files/clean/tv_shows", 0755)
}

func generateTvShowTestFiles(testFiles []string) {
	for i := 0; i < len(testFiles); i++ {
		filemanager.CreateDirSkipIfExists("./test_files/dirty/"+testFiles[i], 0755)
	}

	filemanager.CreateFileSkipIfExists("./test_files/dirty/True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S01E01.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists("./test_files/dirty/True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S01E02.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists("./test_files/dirty/True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S01E03.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists("./test_files/dirty/True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S01E04.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists("./test_files/dirty/True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S1E01.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists("./test_files/dirty/",
		"Batwoman.S02E07.XviD-AFG[TGx]")
	filemanager.CreateFileSkipIfExists("./test_files/dirty/",
		"Batwoman.S02E08.XviD-AFG[TGx].mkv")
}

func generateMovieTestFiles(testFiles []string) {
	for i := 0; i < len(testFiles); i++ {
		filemanager.CreateFileSkipIfExists("./test_files/dirty/", testFiles[i])
	}
}
