package main

import (
	"github.com/markusranda/trostevask/pkg/filemanager"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
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

	tvShowTestDirs := []string{
		"Bobs.Burgers.S04E12.WEB.x264-PHOENiX[TGx]",
		"Stargate SG1 Complete 1997-2007 DVD Rip x264 AC3-MEECH",
		"Stargate Atlantis S01-05 BR 10bit ace hevc-",
		"Bleach 1-366 + Movies Batch Complete",
		"Bleach Season 13 English Dubbed HD",
		"True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD",
	}

	testFiles := []string{
		"flame-man.on.fire.2004.proper.1080p.bluray.x264.rar",
		"man.on.fire.1080p-japhson.png",
		"flame-man.on.fire.2004.proper.1080p.bluray.x264.nfo",
		"man.on.fire.1080p-japhson.png",
	}

	generateTvShowTestFiles(tvShowTestDirs, testFiles)
}

func setupDelayedExtraFiles() {
	movieTestFiles := []string{
		"The.Favourite.2.2020.1080p.BluRay.x264.DTS-WiKi.mkv",
		"Pirates of the carribEan the world is kill 2006 blurAy xHaxx0r.mp4",
		"jason_borne_identity_crisis_2069 CONTRIBUTE TO MOVIEMASTERS.mp4",
		"the greatest escapism (1969).mkv",
		"Grey Swan 2011 dddddd.mkv",
		"2002.A.Space.Adventure.1969.720p.BluRay.DD5.1.x264-LiNG.mkv",
	}

	for i := 0; i < len(movieTestFiles); i++ {
		go createDelayedFile(movieTestFiles[i])
	}
}

func createDelayedFile(file string) {
	time.AfterFunc(getRandomDuration(), func() {
		filemanager.CreateFileSkipIfExists(filemanager.GetInputDir(), file)
	})
}

func getRandomDuration() time.Duration {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(15) // n will be between 0 and 10

	return time.Duration(n) * time.Second
}

func generateTvShowTestFiles(tvShowTestDirs []string, testFiles []string) {
	for i := 0; i < len(tvShowTestDirs); i++ {
		filemanager.CreateDirSkipIfExists(filemanager.GetInputDir()+tvShowTestDirs[i], 0755)
	}

	for i := 0; i < len(testFiles); i++ {
		filemanager.CreateFileSkipIfExists(filemanager.GetInputDir(), testFiles[i])
	}

	filemanager.CreateFileSkipIfExists(filemanager.GetInputDir()+"True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S01E01.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists(filemanager.GetInputDir()+"True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S01E02.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists(filemanager.GetInputDir()+"True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S01E03.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists(filemanager.GetInputDir()+"True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S01E04.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists(filemanager.GetInputDir()+"True.Detective.Season.1.S01.1080p.BluRay.10bit.x265-POIASD/",
		"True.Detective.S1E01.1080p.BluRay.10bit.x265-POIASD.mkv")
	filemanager.CreateFileSkipIfExists(filemanager.GetInputDir(),
		"Batwoman.S02E07.XviD-AFG[TGx]")
	filemanager.CreateFileSkipIfExists(filemanager.GetInputDir(),
		"Batwoman.S02E08.XviD-AFG[TGx].mkv")
}

func generateMovieTestFiles(testFiles []string) {
	for i := 0; i < len(testFiles); i++ {
		filemanager.CreateFileSkipIfExists(filemanager.GetInputDir(), testFiles[i])
	}
}
