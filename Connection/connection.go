package connection

import (
	"fmt"
	"net"
	"strconv"
	"strings"
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
		_,err := conn.Read(buf)
		if err != nil {
			fmt.Print("Read loop error")
			break
		}
		
		i,err := strconv.Atoi(strings.TrimSpace(string(buf)[0:1]))
		j,err:= strconv.Atoi(strings.TrimSpace(string(buf)[3:4]))
		fmt.Println(i,j)
		player.Player2Pos = [2]int{i,j}
	}
}
