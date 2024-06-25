package main

import (
	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	Run()
	return
	//testDisassembly()
	//return
	testParsePe()
	testScript()

}

func testDisassembly() {
	app.Run("asm", func(w *unison.Window) {
		LayoutDisassemblyTable("hyperdbg-cli.exe", w.Content())
	})
}

func testParsePe() {
	app.Run("pe", func(w *unison.Window) {
		LayoutPeView("HPRDBGCTRL.dll", w.Content())
	})
}

func testScript() {
	app.Run("script", func(w *unison.Window) {
		LayoutScript(w.Content())
	})
}
