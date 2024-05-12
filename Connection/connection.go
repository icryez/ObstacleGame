package connection

import (
	"fmt"
	"net"
	"time"

	gametick "github.com/MultiplayerObsGame/GameTick"
	player "github.com/MultiplayerObsGame/PlayerModule"
)

func ConnectToServer(){
	conn, err := net.Dial("tcp", "localhost:3000")
	defer conn.Close()
	if err!=nil {
		fmt.Println("Error while connecting to server :", err)
	}
	for gametick.EndGame == false {
		time.Sleep(1000*time.Millisecond)
		//TODO: better implementation
		str := fmt.Sprint("098765",player.PlayerPos[0], player.PlayerPos[1])

		conn.Write([]byte(str))
	}
}
