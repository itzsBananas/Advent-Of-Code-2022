package main

import (
	"fmt"
)

type FileSystem struct {
	rootDir *Dir
	currDir *Dir
}

func CreateFileSystem() FileSystem {
	root := &Dir{"/", make([]Item, 0), nil}
	return FileSystem{root, root}
}

func (fs *FileSystem) ChangeDirectory(arg string) {
	switch arg {
	case "/":
		fs.currDir = fs.rootDir
	case "..":
		fs.currDir = fs.currDir.parentDir
	default:
		for _, item := range fs.currDir.list {
			switch dir := item.(type) {
			case *Dir:
				if dir.GetName() == arg {
					fs.currDir = dir
					break
				}
			}
		}
	}
}

func (fs *FileSystem) AddFile(f *File) {
	fs.currDir.list = append(fs.currDir.list, f)
}

func (fs *FileSystem) AddDir(d *Dir) {
	d.parentDir = fs.currDir
	fs.currDir.list = append(fs.currDir.list, d)
}

type File struct {
	name string
	size int
}

func CreateFile(n string, s int) File {
	return File{name: n, size: s}
}

type Dir struct {
	name      string
	list      []Item
	parentDir *Dir
}

func CreateDir(n string) Dir {
	return Dir{name: n}
}

func (d Dir) TotalSize(callback func(int)) int {
	size := 0
	for _, item := range d.list {
		switch item := item.(type) {
		case *Dir:
			size += item.TotalSize(callback)
		case *File:
			size += item.size
		}
	}
	if callback != nil {
		callback(size)
	}

	return size
}

// could have removed item and made two seperate lists
type Item interface {
	// fmt.Stringer
	GetName() string
}

func (f File) GetName() string {
	return f.name
}

func (d Dir) GetName() string {
	return d.name
}

func (f File) String() string {
	return fmt.Sprintf("- %s (file, size=%d)", f.GetName(), f.size)
}

func (d Dir) String() string {
	return fmt.Sprintf("- %s (dir)", d.GetName())
}
