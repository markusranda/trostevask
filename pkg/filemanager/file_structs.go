package filemanager

import "os"

type FullFileInfo struct {
	os.FileInfo
	Path string
}
