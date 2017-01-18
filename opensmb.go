package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"net/url"
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

	if !isWindowsSmbPath(windowsPath) { 
		fmt.Println("pathがwindowsのものじゃないです")
		os.Exit(1)
	}	

	macPath := genMacPathFromWindowsPath(windowsPath) 
	mountTargetStr := genSvrAndPathStr(macPath)
	
	fmt.Println(mountTargetStr)
}

func isWindowsSmbPath(pathString string) (b bool) {
    if m, _ := regexp.MatchString("^\\\\\\\\.*\\.*", pathString); !m {
        return false
    }
    return true
}

func genMacPathFromWindowsPath(windowsPath string) (macPath string) {
	macPath = strings.Replace(windowsPath, "\\", "/", -1)
	return macPath
}

func genSvrAndPathStr(macPath string) (string){
    re, err := regexp.Compile("//(.*?)/")
    if err != nil {
            panic(err)
    }

    svr_index := re.FindStringIndex(macPath)
    mountSvrStr := macPath[svr_index[0]:svr_index[1]]
	mountPathStr := macPath[svr_index[1]:]
	escapedMountPathStr := url.QueryEscape(mountPathStr)
	mountTargetStr := mountSvrStr + escapedMountPathStr

	return mountTargetStr
}
