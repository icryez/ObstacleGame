package gametick

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	colors "github.com/MultiplayerObsGame/Colours"
	connection "github.com/MultiplayerObsGame/Connection"
	keyboard "github.com/MultiplayerObsGame/Keyboard"
	mapmodule "github.com/MultiplayerObsGame/MapModule"
	player "github.com/MultiplayerObsGame/PlayerModule"
	structs "github.com/MultiplayerObsGame/Structs"
	"github.com/MultiplayerObsGame/terminal"
)

var inair bool
var GameStarted bool
var EndGame bool
var connectedToServer bool
var sessionId string = ""
var listeningForInput bool
var printSessionScreen bool

func Tick() {
	mapmodule.GenMap()
	if !GameStarted {
		printStartScreen()
	}
	if printSessionScreen {
		printSessionInputScreen()
	}
	if GameStarted {
		terminal.CallClearCmd()
		for EndGame == false {
			time.Sleep(1 * time.Millisecond)
			PrintMap()
		}
	}
}

func printStartScreen() {
	colors.BlueText.Println("Press A to move LEFT and D to move RIGHT")
	colors.BlueText.Println("Guess what the JUMP key is?")
	colors.BlueText.Println("Press ESC to END game")
	colors.BlueText.Println("Press SPACE to start")
	if connectedToServer {
		colors.BlueText.Println("Connected to Server with SessionID:", sessionId) //TODO: prints this even if connection failed - fix this
	} else {
		colors.BlueText.Println("Not connected to server, Press C to enter sessionID")
	}
	for {
		if keyboard.KeysState.GetKey("space") {
			GameStarted = true
			break
		} else if keyboard.KeysState.GetKey("Esc") {
			EndGame = true
			break
		} else if keyboard.KeysState.GetKey("C") && printSessionScreen == false {
			printSessionScreen = true
			break
		}
	}
}

func printSessionInputScreen() {
	colors.BlueText.Print("Enter Session ID (Length of 6) : ") //TODO: implement length check
	var b []byte = make([]byte, 1)
	buf := bufio.NewReader(os.Stdin)
	var flag bool
	var buffer bytes.Buffer
	for {
		buf.Read(b)
		if string(b) != "\n" {
			if flag {
				buffer.WriteString(string(b))
				fmt.Print(string(b))
			} else if string(b) == "c" || string(b) == "C" && flag == false {
				flag = true
			}
		} else {
			sessionId = buffer.String()
			terminal.CallClearCmd()
			go connection.ConnectToServer(sessionId)
			connectedToServer = true //TODO: check if connection was successful then make this true
			printStartScreen()
			break
		}
	}
}

func PrintMap() {
	terminal.MoveCursor(0, 0)
	for r := range structs.VisibleMatrix {
		for c, val := range structs.VisibleMatrix[r] {
			if player.PlayerPos == [2]int{r, c} || player.Player2Pos == [2]int{r, c} {
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
		if keyboard.KeysState.GetKey("Esc") {
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
