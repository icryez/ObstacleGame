package connection

import (
	"fmt"
	"net"
	"time"

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
	go readLoop(conn)
	for{
		time.Sleep(100 * time.Millisecond)
		//TODO: better implementation
		str := fmt.Sprint("098765", player.PlayerPos[0], player.PlayerPos[1])
		conn.Write([]byte(str))
	}
}

func readLoop(conn net.Conn){
	buf := make([]byte,2048)
	for {
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Print("Read loop error")
		}
		fmt.Println(string(buf[:n]))
	}
}
