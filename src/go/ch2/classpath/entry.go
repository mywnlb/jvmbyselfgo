package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

/*
基类定义
*/
type Entry interface {
	//读取class
	readClass(className string) ([]byte, Entry, error)
	String() string
}

/*
根据传入的路径不同利用不同的实现类获取class文件
*/
func newEntry(path string) Entry {

	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, "zip") || strings.HasSuffix(path, "ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
