package main

import "os"
import "io"
import "crypto/sha1"
import "io/ioutil"
import "fmt"
import "strconv"
import s "strings"
import "path/filepath"

var pr = fmt.Println

func main() {
	ergodic(".")
}

func ergodic(dirPath string) {
	dir, _ := ioutil.ReadDir(dirPath)
	for _, fi := range dir {
		if fi.IsDir() {
			p, _ := filepath.Abs(filepath.Dir(fi.Name()))
			pname := s.Join([]string{p, fi.Name()}, "\\")
			ergodic(pname)
		} else {
			pname := s.Join([]string{dirPath, fi.Name()}, "\\")
			fil, _ := SHA1File(pname)
			pnamestring := s.Join([]string{fi.Name(), fil, strconv.FormatInt(fi.Size(), 10)}, " ")
			pnamestrings := s.Join([]string{"\n", pnamestring, ","}, " ")
			appendToFile("./test.txt", pnamestrings)
		}
	}
}

func SHA1File(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	h := sha1.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// fileName:文件名字(带全路径)
// content: 写入的内容
func appendToFile(fileName string, content string) error {
	// 以创建的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}
