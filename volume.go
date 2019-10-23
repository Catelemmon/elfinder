package elfinder

import (
	"io"
	"os"
)

type Volume interface {
	Stat(path string) (os.FileInfo, error)

	Readdir(path string) ([]os.FileInfo, error)

	GetFile(path string) (reader io.ReadCloser, err error)

	UploadFile(dir, uploadPath, filename string, reader io.Reader) (os.FileInfo, error)

	UploadChunk(cid int, dirPath, uploadPath, filename string, rangeData ChunkRange, reader io.Reader) error

	MergeChunk(cid, total int, dirPath, uploadPath, filename string) (os.FileInfo, error)

	MakeDir(dir, newDirname string) (os.FileInfo, error)

	MakeFile(dir, newFilename string) (os.FileInfo, error)

	Rename(oldNamePath, newname string) (os.FileInfo, error)

	Remove(path string) error

	Paste(dir, filename, suffix string, reader io.ReadCloser) (os.FileInfo, error)

	RootDir() os.FileInfo
}
