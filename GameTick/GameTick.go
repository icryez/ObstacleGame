package gametick

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"time"

	colors "github.com/MultiplayerObsGame/Colours"
	mapmodule "github.com/MultiplayerObsGame/MapModule"
	player "github.com/MultiplayerObsGame/PlayerModule"
	structs "github.com/MultiplayerObsGame/Structs"
	"github.com/MultiplayerObsGame/terminal"
)

func Tick() {
	mapmodule.GenMap()
	for {
		time.Sleep(10 * time.Millisecond)
		PrintMap()
	}
}

func PrintMap() {
	terminal.MoveCursor(0, 0)
	for r := range structs.VisibleMatrix {
		for c, val := range structs.VisibleMatrix[r] {
			if player.PlayerPos == [2]int{r, c} {
				colors.Red.Print(" ")
			} else if val.IsFloor {
				colors.Yellow.Print(" ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	terminal.CallFlush()
}

func ListenForPlayerMovements() {

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// var b []byte = make([]byte, 1)
	f, err := os.Open("/dev/input/event6")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 24)
	for {
		//TODO: Implement keyboard watcher layer instead of actiing on raw binary
		f.Read(b)
		var value int32
		// typ := binary.LittleEndian.Uint16(b[16:18])
		code := binary.LittleEndian.Uint16(b[18:20])
		binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &value)
		// fmt.Printf("type: %x\ncode: %d\nvalue: %d\n", typ, code, value)
		if code == 57 && value == 1 && structs.VisibleMatrix[player.PlayerPos[0]+1][player.PlayerPos[1]].IsVisible {
			go jump()
		} else if code == 32 && (value == 1 || value == 2) {
			moveRight()
		} else if code == 30 && (value == 1 || value == 2) {
			moveLeft()
		}
	}
}

func jump() {
	for i := 0; i < 4; i++ {
		time.Sleep(20 * time.Millisecond)
		if player.PlayerPos[0] >= 0 && player.PlayerPos[0] < 29 {
			player.PlayerPos[0] = player.PlayerPos[0] - 1
		}
	}
}
func moveLeft() {
	if player.PlayerPos[1] >= 0 && player.PlayerPos[1] < 100 {
		player.PlayerPos[1] = player.PlayerPos[1] - 1
	}
}

func moveRight() {
	if player.PlayerPos[1] >= 0 && player.PlayerPos[1] < 100 {
		player.PlayerPos[1] = player.PlayerPos[1] + 1
	}
}

func IsBlockUnderFloor() {

}

func StartGravity() {
	for {
		time.Sleep(200 * time.Millisecond)
		if player.PlayerPos[0] < 28 &&
			structs.VisibleMatrix[player.PlayerPos[0]][player.PlayerPos[1]].IsVisible == false {
			player.PlayerPos[0] += 1
		}
	}
}
