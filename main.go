package main

import (
	"time"

	"atomicgo.dev/cursor"
	gametick "github.com/MultiplayerObsGame/GameTick"
	player "github.com/MultiplayerObsGame/PlayerModule"
	"github.com/MultiplayerObsGame/terminal"
)

func main() {
	terminal.CallClearCmd()
	cursor.Hide()
	terminal.MoveCursor(0,0)
	player.PlayerStart()
	go gametick.ListenForPlayerMovements()
	go gametick.Tick()
	go gametick.StartGravity()
	for gametick.EndGame == false{}
	terminal.CallClearCmd()
	cursor.Show()
	time.Sleep(10*time.Millisecond)
}
