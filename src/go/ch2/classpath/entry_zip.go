package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/*
压缩包中寻找
*/
type ZipEntry struct {
	absPath string
	zipRc   *zip.ReadCloser
}

/*
解压并从中获取对应的class
*/
func (self *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	if self.zipRc == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := self.findClass(className)

	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := ReadClass(classFile)

	return data, self, err
}

func ReadClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, nil
	}

	data, err := ioutil.ReadAll(rc)

	rc.Close()
	if err != nil {
		return nil, nil
	}

	return data, nil
}

func (self *ZipEntry) String() string {
	return self.absPath
}

func (self *ZipEntry) openJar() error {
	r, err := zip.OpenReader(self.absPath)
	if err == nil {
		self.zipRc = r
	}
	return err
}

func (self *ZipEntry) findClass(className string) *zip.File {
	for _, f := range self.zipRc.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath, nil}
}
