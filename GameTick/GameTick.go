package gametick

import (
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
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		if string(b) == " " && structs.VisibleMatrix[player.PlayerPos[0]+1][player.PlayerPos[1]].IsVisible {
			jump()
		} else if string(b) == "d" || string(b) == "D" {
			moveRight()
		} else if string(b) == "a" || string(b) == "A" {
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
