//+buildignore
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	path  = "/opt/gopath/src/github.com/wangfmD/fstools/logs"
	Files = make([]string, 0, 20)
)

// func test() {
// pa := filepath.Dir("/opt/gopath/src/github.com/wangfmD/fstools/logs")
// fmt.Println(pa)
// osType := os.Getenv("GOOS")
// fmt.Println(osType)
// }

func Listfunc(path string) {
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, f := range fs {
			if f.IsDir() {
				// fmt.Println("dir: ", f.Name())
				Listfunc(path + "/" + f.Name())
			} else {
				fmt.Println("file:", f.Name())
			}
		}
	}
}

func ListF(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("prevent panic by handling failure accessing a path  %v\n", err)
		return err
	}
	if !info.IsDir() {
		fmt.Println("file:" + info.Name())
	}
	return nil
}

// func getFileList(path string) {
// filepath.Walk(path, ListF)
// }

func getFileList(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println("file:", path)
		Files = append(Files, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

// func main() {
// 	// test()
// 	getFileList(path)
// 	for _, v := range Files {
// 		println(v)
// 	}
// 	// Listfunc(path)
// }
