package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	//绝对路径
	absDir string
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}

/*
获取绝对路径
*/
func newDirEntry(path string) *DirEntry {

	absDir, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}
