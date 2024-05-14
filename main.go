package main

import (
	"time"

	"atomicgo.dev/cursor"
	connection "github.com/MultiplayerObsGame/Connection"
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
	go connection.ConnectToServer()
	for gametick.EndGame == false{}
	terminal.CallClearCmd()
	cursor.Show()
	time.Sleep(10*time.Millisecond)
}
