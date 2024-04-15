package main

import (
	gametick "github.com/MultiplayerObsGame/GameTick"
	player "github.com/MultiplayerObsGame/PlayerModule"
	"github.com/MultiplayerObsGame/terminal"
)

func main() {
	terminal.CallClearCmd()
	terminal.MoveCursor(0,0)
	player.PlayerStart()
	go gametick.Tick()
	go gametick.StartGravity()
	go gametick.ListenForPlayerMovements()
	for {}
}
