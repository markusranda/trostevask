package sniffer

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/markusranda/trostevask/pkg/cleaner"
	"github.com/markusranda/trostevask/pkg/filemanager"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var watcher *fsnotify.Watcher

func Sniff() {
	// creates a new file watcher
	watcher, _ = fsnotify.NewWatcher()
	defer stopListener()

	// starting at the root of the project, walk each file/directory searching for
	// directories
	if err := filepath.Walk(filemanager.GetInputDir(), watchDir); err != nil {
		fmt.Println("ERROR", err)
	}

	for {
		select {
		// watch for events
		case event := <-watcher.Events:
			if event.Op == fsnotify.Create {
				processFile(event.Name)
			}
		// watch for errors
		case err := <-watcher.Errors:
			fmt.Println("ERROR", err)
		}
	}
}

func processFile(path string) {
	var file filemanager.FullFileInfo
	file.Path = path
	file.FileInfo, _ = os.Stat(path)
	go cleaner.CleanFileName(file)
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi os.FileInfo, err error) error {

	if fi == nil {
		return nil
	}

	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}

func stopListener() {
	err := watcher.Close()
	if err != nil {
		logrus.Fatal(err)
	}
}
