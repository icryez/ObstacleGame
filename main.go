package main

import (
	"time"

	"atomicgo.dev/cursor"
	colors "github.com/MultiplayerObsGame/Colours"
	gametick "github.com/MultiplayerObsGame/GameTick"
	keyboard "github.com/MultiplayerObsGame/Keyboard"
	player "github.com/MultiplayerObsGame/PlayerModule"
	"github.com/MultiplayerObsGame/terminal"
)

func main() {
	terminal.CallClearCmd()
	cursor.Hide()
	terminal.MoveCursor(0,0)
	player.PlayerStart()
	keyboard.KeysState = *keyboard.CreateKeyBoardState()
	go gametick.ListenForPlayerMovements()
	go gametick.Tick()
	go gametick.StartGravity()
	// go connection.ConnectToServer()
	for gametick.EndGame == false{}
	terminal.CallClearCmd()
	colors.BlueText.Println("Esc Pressed - exiting game")
	cursor.Show()
	time.Sleep(10*time.Millisecond)
}
