package renamer

import (
	"io/ioutil"
	"log"
	"testing"
)

var testFileNames []string

func init() {
	testFileNames = []string{"Pirates of the carribEan 2003 blurAy xHaxx0r.mp4",
		"jason_borne_identity_2002 CONTRIBUTE TO MOVIEMASTERS.mp4",
		"the great escape (1963).mkv",
		"Black Swan 2010 dddddd.mkv"}
	generateTestFiles()
}

func RenameNonLowerCaseLowercase(t *testing.T) {
	oldFileName := "Pirates of the carribEan 2003 blurAy xHaxx0r.mp4"
	correctFileName := "pirates of the carribean 2003 bluray xhaxx0r.mp4"
	newFileName := rename(oldFileName)

	if newFileName != correctFileName {
		t.Error("rename was incorrect, should be all lowercase")
	}
}

func generateTestFiles() {

	for i := 0; i < len(testFileNames); i++ {
		createFile(testFileNames[i])
	}
}

func createFile(filename string) {
	bytes := []byte(filename)
	err := ioutil.WriteFile("../../test_files/"+filename, bytes, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
