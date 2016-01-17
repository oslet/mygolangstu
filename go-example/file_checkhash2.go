package main

import (
	"crypto/md5"
	"io"
	//"bytes"

	"fmt"

	"os"
	//"os/exec"
	"path/filepath"
)

func walkFunc(path string, filename os.FileInfo, err error) error {
	if filename.IsDir() && filename.Name() == "test" { // 忽略目录
		return filepath.SkipDir
	}

	fmt.Printf("%s\n", path)
	return nil

}

func main() {
	//遍历打印所有的文件名
	//filepath.Walk(".", walkFunc)

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() && info.Name() != "test*" {

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			h := md5.New()
			io.Copy(h, file)
			file.Close()
			//dir := filepath.Dir(path)
			//fmt.Printf("%v\n", filepath.Join(dir, fmt.Sprintf("%x", h.Sum(nil))))
			fmt.Printf("%x %s\n", h.Sum(nil), path)
			//fmt.Printf("%v\n ", path)
		}
		return nil
	})
}
