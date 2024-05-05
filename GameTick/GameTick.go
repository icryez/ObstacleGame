package gametick

import (
	"fmt"
	"os/exec"
	"time"

	colors "github.com/MultiplayerObsGame/Colours"
	keyboard "github.com/MultiplayerObsGame/Keyboard"
	mapmodule "github.com/MultiplayerObsGame/MapModule"
	player "github.com/MultiplayerObsGame/PlayerModule"
	structs "github.com/MultiplayerObsGame/Structs"
	"github.com/MultiplayerObsGame/terminal"
)

var inair bool
var GameStarted bool
var EndGame bool

func Tick() {
	mapmodule.GenMap()
	if !GameStarted {
		printStartScreen()
	}
	terminal.CallClear()
	if GameStarted {
		for EndGame == false {
			time.Sleep(10 * time.Millisecond)
			PrintMap()
		}
	}
}

func printStartScreen() {
	colors.BlueText.Println("Press A to move LEFT and D to move RIGHT")
	colors.BlueText.Println("Guess what the JUMP key is?")
	colors.BlueText.Println("Press ESC to END game")
	colors.BlueText.Println("Press SPACE to start")
	terminal.CallFlush()
	for !GameStarted {
		if keyboard.KeysState.GetKey("space") {
			GameStarted = true
		} else if keyboard.KeysState.GetKey("Esc"){
			EndGame = true
		}
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
	go keyboard.StartWatcher()
	for EndGame == false {
		time.Sleep(40 * time.Millisecond)
		if keyboard.KeysState.GetKey("space") && inair == false && structs.VisibleMatrix[player.PlayerPos[0]+1][player.PlayerPos[1]].IsVisible {
			inair = true
			go jump()
		}
		if keyboard.KeysState.GetKey("D") {
			moveRight()
		}
		if keyboard.KeysState.GetKey("A") {
			moveLeft()
		}
		if keyboard.KeysState.GetKey("Esc"){
			EndGame = true
		}
	}
}

func jump() {
	for i := 0; i < 6; i++ {
		time.Sleep(50 * time.Millisecond)
		if player.PlayerPos[0] >= 0 && player.PlayerPos[0] < 29 {
			player.PlayerPos[0] = player.PlayerPos[0] - 1
		}
	}
	inair = false
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
		time.Sleep(80 * time.Millisecond)
		if player.PlayerPos[0] < 28 &&
			structs.VisibleMatrix[player.PlayerPos[0]][player.PlayerPos[1]].IsVisible == false &&
			inair == false {
			player.PlayerPos[0] += 1
		}
	}
}
