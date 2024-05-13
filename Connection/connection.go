package connection

import (
	"fmt"
	"net"
	player "github.com/MultiplayerObsGame/PlayerModule"
)

func ConnectToServer() {
	conn, err := net.Dial("tcp", "localhost:3000")
	defer conn.Close()
	if err != nil {
		fmt.Println("Error while connecting to server :", err)
	} else {
		fmt.Println("Connected to :",conn.RemoteAddr())
	}
	buf := make([]byte, 2048)
	for{
		//TODO: better implementation
		str := fmt.Sprint("098765", player.PlayerPos[0], player.PlayerPos[1])
		fmt.Println(str)
		conn.Write([]byte("sdfsfs"))
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Read Loop Error - disconnected from %s : %s", conn.RemoteAddr(), err)
			break
		}
		readStr := string(buf[:n])
		fmt.Println(readStr)
	}
}

