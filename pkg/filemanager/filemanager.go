package filemanager

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func RemoveContents(dir string) error {
    d, err := os.Open(dir)
    if err != nil {
        return err
    }
    defer d.Close()
    names, err := d.Readdirnames(-1)
    if err != nil {
        return err
    }
    for _, name := range names {
        err = os.RemoveAll(filepath.Join(dir, name))
        if err != nil {
            return err
        }
    }
    return nil
}

func CreateFile(dir string, filename string) {
	bytes := []byte(filename)
	err := ioutil.WriteFile(dir+filename, bytes, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func MoveFile(oldLocation string, newLocation string) {
    err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func MoveFilesFromTo(oldLocation string, newLocation string) {
    var fileNameList = GetFileNamesFromDir(oldLocation)
	
    for _, filename := range fileNameList {
		var oldFilename = oldLocation + filename
		var newFilename = newLocation + filename
		MoveFile(oldFilename, newFilename)	
	}
}

func GetFileNamesFromDir(dir string) (files []string) {
    var fileInfoArray = getFilesFromDir(dir)

    for i := 0; i < len(fileInfoArray); i++ {
        filename := fileInfoArray[i].Name()
        files = append(files, filename)
    }
    return
}

func ReadAndPrintFiles(dir string) {
    var files = getFilesFromDir(dir)
    for _, file := range files {
        fmt.Println(file.Name())
    }
}

func getFilesFromDir(dir string) (files []os.FileInfo) {
	files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }
    return
}