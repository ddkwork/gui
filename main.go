package main

import (
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/hyperdbgui/libhyperdbg"
	"io/fs"
	"os"
	"path/filepath"
)

//go:generate go build .
//go:generate go install .

func main() {
	libhyperdbg.Run()
}

func Release() { // todo let sdk work first
	// build HyperDbg
	stream.RunCommand("git clone --recursive https://github.com/HyperDbg/HyperDbg.git")
	stream.RunCommand("HyperDbg/utils/msvc-build.bat")
	filepath.Walk("bin/debug", func(path string, info fs.FileInfo, err error) error {
		switch filepath.Ext(path) {
		case ".pdb", ".exp", ".lib":
			mylog.Check(os.Remove(path))
		}
		return err
	})
	mylog.Todo("remove outer files or dirs")

	// build gui
	stream.RunCommand("go build .")
	stream.CopyFile("hyperdbgui.exe", "bin/debug")

	// release to github
	stream.RunCommand("tar -zcvf hyperdbgui.tar.gz bin/debug")
	mylog.Todo("now post to github release use action")

	// clean
	mylog.Check(os.RemoveAll("bin"))
}
