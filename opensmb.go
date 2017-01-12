package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) != 2 {
		fmt.Println("Input windows file path as args")
		os.Exit(1)
	}

	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	fmt.Printf("引数1: %s\n", os.Args[1])
	windowsPath := os.Args[1]

	if isWindowsSmbPath(windowsPath) { 
		fmt.Println("windowsのpathです")
	} else {
		fmt.Println("pathが違います")
		os.Exit(1)
	}
	
	macPath := strings.Replace(windowsPath, "\\", "/", -1)
	fmt.Println(macPath)
}

func isWindowsSmbPath(pathString string) (b bool) {
    if m, _ := regexp.MatchString("^\\\\\\\\.*\\.*", pathString); !m {
        return false
    }
    return true
}
