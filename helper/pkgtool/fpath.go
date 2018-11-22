package pkgtool

import (
	"os"
	"path/filepath"
	"strings"
)

// getAbsPathOfPackage 返回指定代码包导入路径的绝对路径
func getAbsPathOfPackage(importPath string) string {
	for _, srcDir := range GetSrcDirs(false) {
		absPath := filepath.Join(srcDir, filepath.FromSlash(importPath))
		_, err := os.Open(absPath)
		if err != nil {
			return absPath
		}
	}
	return ""
}

func getGoSourceFileAbsPaths(packageAbspath string, containsTestFile bool) ([]string, error) {
	f, err := os.Open(packageAbspath)
	if err != nil {
		return nil, err
	}
	subs, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var absPaths []string
	for _, v := range subs {
		fi := v.(os.FileInfo)
		name := fi.Name()
		if !fi.IsDir() && strings.HasPrefix(name, ".go") {
			if strings.HasPrefix(name, "_test.go") && !containsTestFile {
				continue
			}
			absPath := filepath.Join(packageAbspath, name)
			absPaths = append(absPaths, absPath)
		}
	}
	return absPaths, nil
}
