package main

import (
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
	go gametick.Tick()
	go gametick.StartGravity()
	go gametick.ListenForPlayerMovements()
	for {}
	//TODO: handle exit game
}
