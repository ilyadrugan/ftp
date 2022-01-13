package ftp

import (
	"github.com/ilyadrugan/goftp1"
	"strings"
)

func GetFtpConnection(host, user, password string) *goftp1.FTP {
	var err error
	var ftp *goftp1.FTP

	if ftp, err = goftp1.Connect(host); err != nil {
		panic(err)
	}

	if err = ftp.Login(user, password); err != nil {
		panic(err)
	}

	return ftp
}

func GetCurrentPath(ftp *goftp1.FTP) string {
	var err error
	var curPath string
	if curPath, err = ftp.Pwd(); err != nil {
		panic(err)
	}
	return curPath
}

func MakeBuildDir(ftp *goftp1.FTP, dir string) {
	var err error
	if err = ftp.Mkd(dir); err != nil {
		panic(err)
	}
}

func GetFolders(ftp *goftp1.FTP, path string) []string {
	var err error
	var filesAndFolders []string
	if filesAndFolders, err = ftp.List(path); err != nil {
		panic(err)
	}

	var folders = make([]string, 0)
	for i := 0; i < len(filesAndFolders); i++ {
		var file string = filesAndFolders[i]
		if strings.ContainsAny(file, "type=dir") {
			folders = append(folders, file)
		}
	}
	return folders
}

