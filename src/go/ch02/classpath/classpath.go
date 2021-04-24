package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClasspath Entry
	exitClasspath Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExitClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *ClassPath) parseBootAndExitClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.exitClasspath = newWildcardEntry(jreExtPath)
}

func (self *ClassPath) parseUserClasspath(option string) {
	if option == "" {
		option = "."
	}

	self.userClasspath = newEntry(option)
}

/**
查找boot目录如果传入路径不存在->当前目录不存在->环境变量是否存在->返回错误
*/
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("can not find jre folder!")
}

/**
目录是否存在
*/
func exists(jreOption string) bool {
	if _, err := os.Stat(jreOption); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err != nil {
		return data, entry, err
	}

	if data, entry, err := self.exitClasspath.readClass(className); err != nil {
		return data, entry, err
	}

	return self.userClasspath.readClass(className)
}

func (self *ClassPath) String() string {
	return self.userClasspath.String()
}
