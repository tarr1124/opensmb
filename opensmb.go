package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"net/url"
	"path"
	"syscall"
	"crypto/sha1"
	"encoding/hex"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Input windows file path as args")
		os.Exit(1)
	}

	windowsPath := os.Args[1]

	if !isWindowsSmbPath(windowsPath) { 
		fmt.Println("pathがwindowsのものじゃないです")
		os.Exit(1)
	}	

	macPath := genMacPathFromWindowsPath(windowsPath) 
	mountTargetStrWithPlus, targetFile := genSvrAndPathStr(macPath)
	mountTargetStr := strings.Replace(mountTargetStrWithPlus, "+", "%20", -1)

	workingDirPath := "/Volumes/"

	mountedHashedDirName := genDirName(mountTargetStr)
	mountedDirPath := workingDirPath + mountedHashedDirName
	if isExist(mountedDirPath) {
		fmt.Println("対象のDirがすでにマウント済みです。")
	} else {
		fmt.Println("マウントします")
		prepareWorkingDir(mountedDirPath)
		mountNewVol(mountTargetStr, mountedDirPath)
	}

	fmt.Println(mountedDirPath + "/" + targetFile)
	out, err := exec.Command("open", mountedDirPath + "/" + targetFile).Output()
    if err != nil {
            panic(err)
    }
	fmt.Println(out)
}

func mountNewVol(mountTargetStr string, mountedDirPath string) {
	fmt.Println(mountTargetStr + "         " + mountedDirPath)
	out, err := exec.Command("mount_smbfs", mountTargetStr, mountedDirPath).Output()
    if err != nil {
            panic(err)
    }
	fmt.Println(out)
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

func genSvrAndPathStr(macPath string) (string, string){
    re, err := regexp.Compile("//(.*?)/")
    if err != nil {
            panic(err)
    }

    svr_index := re.FindStringIndex(macPath)
    mountSvrStr := macPath[svr_index[0]:svr_index[1]]
	mountPathStr := macPath[svr_index[1]:]
	mountDir, targetFile := path.Split(mountPathStr)
	escapedMountPathStr := url.QueryEscape(mountDir)
	mountTargetStr := "smb:" + mountSvrStr + escapedMountPathStr

	return mountTargetStr, targetFile
}

func prepareWorkingDir(dirPath string) {
    defaultUmask := syscall.Umask(0)
    os.Mkdir(dirPath, 0777)
    syscall.Umask(defaultUmask)
}

func genDirName(dirNameKey string) (string) {
	data := []byte(dirNameKey)
	b := sha1.Sum(data)
	fmt.Println(hex.EncodeToString(b[:]))
	return hex.EncodeToString(b[:])
}

func isExist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}
