package main

import (
	"fmt"
	"jvmbyselfgo/src/go/ch2/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	cmd.XjreOption = ""
	cmd.cpOption = "java.lang.Object"
	startJVM(cmd)
	//if cmd.versionFlag {
	//	fmt.Println("version 0.0.0.1")
	//} else if cmd.helpFlag || cmd.class == "" {
	//	printUsage()
	//} else {
	//	startJVM(cmd)
	//}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not found or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data:%v \n", classData)
}
