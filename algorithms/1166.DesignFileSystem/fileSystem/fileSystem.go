package fileSystem

import "strings"

type FileSystem struct {
	Path map[string]int
}


func Constructor() FileSystem {
	path := make(map[string]int)
	fileSystem := FileSystem{Path:path}
	return fileSystem
}


func (this *FileSystem) Create(path string, value int) bool {
	seps := strings.Split(path, "/")
	if len(seps) > 2 {
		temp := ""
		for i := 1; i < len(seps) - 1; i++ {
			if len(seps[i]) == 0 {
				continue
			}
			temp = temp + "/" + seps[i]
			if _, ok := this.Path[temp]; !ok {
				return false
			}
		}
	}
	if _, ok := this.Path[path]; ok {
		return false
	}
	this.Path[path] = value
	return true
}


func (this *FileSystem) Get(path string) int {
	v, ok := this.Path[path]
	if !ok {
		return -1
	}
	return v
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Create(path,value);
 * param_2 := obj.Get(path);
 */