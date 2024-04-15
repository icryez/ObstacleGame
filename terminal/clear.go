package terminal

import (
	"os"
	"os/exec"
	"runtime"

	tm "github.com/buger/goterm"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClearCmd(){
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok { //if we defined a clear func for that platform:
	    value()  //we execute it
	} else { //unsupported platform
	    panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func CallClear() {
	tm.Clear()
}

func MoveCursor(pos1 int , pos2 int){
	tm.MoveCursor(pos1,pos2)
}
 func CallFlush(){
	tm.Flush()
}
