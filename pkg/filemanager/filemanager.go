package filemanager

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
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

func CreateFileSkipIfExists(dir string, filename string) {
	bytes := []byte(filename)
	err := ioutil.WriteFile(dir+filename, bytes, 0644)

	if err != nil {
		return
	}
}

func CreateDir(dir string, permissions os.FileMode) {
	err := os.Mkdir(dir, permissions)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDirSkipIfExists(dir string, permissions os.FileMode) {
	err := os.Mkdir(dir, permissions)
	if err != nil {
		return
	}
}

func CreateDirIfNotExists(dir string, permissions os.FileMode) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, permissions)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func MoveFile(oldLocation string, newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		logrus.Error(err.Error())
	}
}

func CopyFile(src string, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
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

func GetFileNamesFromDirRecursive(dir string) (files []string) {
	var fileInfoArray = GetFilesFromDirRecursive(dir)

	for i := 0; i < len(fileInfoArray); i++ {
		filename := fileInfoArray[i].Name()
		files = append(files, filename)
	}
	return
}

func ReadAndPrintAllFiles(dir string) {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func getFilesFromDir(dir string) (files []os.FileInfo) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetFilesFromDirRecursive(dir string) (files []FullFileInfo) {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			var file FullFileInfo
			file.FileInfo = info
			file.Path = path
			files = append(files, file)

			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return files
}

func IsFolder(filename string) (isFolder bool) {
	info, err := os.Stat(filename)

	if err != nil {
		return false
	}
	return info.IsDir()
}

func GetInputDir() (baseDir string) {
	baseDir = os.Getenv("input_dir")
	last := baseDir[len(baseDir)-1:]
	if last != "/" {
		baseDir += "/"
	}
	return baseDir
}

func GetOutputDir() (baseDir string) {
	baseDir = os.Getenv("output_dir")
	last := baseDir[len(baseDir)-1:]
	if last != "/" {
		baseDir += "/"
	}
	return baseDir
}

func GetRejectedtDir() (baseDir string) {
	baseDir = os.Getenv("rejected_dir")
	last := baseDir[len(baseDir)-1:]
	if last != "/" {
		baseDir += "/"
	}
	return baseDir
}

func FileExists(dest string) bool {
	if _, err := os.Stat(dest); err == nil {
		// path/to/whatever exists
		return true

	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		return true

	}
}
