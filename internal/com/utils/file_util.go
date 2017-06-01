package utils
import (
	"os"
	"strings"
	"path/filepath"
	"io/ioutil"
)

/**
 * 遍历制定路径的目录, 不递归子目录
 */
func ListFiles(dirPth string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		// 忽略目录
		if fi.IsDir() {
			continue
		}
		files = append(files, dirPth+PthSep+fi.Name())
	}
	return files, nil
}

/**
 * 遍历制定路径的目录, 不递归子目录, 只获取制定后缀的文件
 */
func FilterFiles(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		// 忽略目录
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

/**
 * 遍历制定路径下的所有文件
 */
func WalkDir(dirPth string) (files []string, err error) {
	files = make([]string, 0, 30)

	//遍历目录
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		files = append(files, filename)
		return nil
	})

	return files, err
}

/**
 * 递归遍历制定路径下的所有文件
 */
func FilterWalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)

	//遍历目录
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})

	return files, err
}

/**
 * 判断路径是否存在(文件/文件夹)
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 * 移动文件到新路径
 */
func MvFile(from string, to string) error {
	return os.Rename(from, to)
}

