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
		str := fmt.Sprint("098765", 
			integerToStringOfFixedWidth(player.PlayerPos[0],2),
			integerToStringOfFixedWidth(player.PlayerPos[1],2))
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
		i,err := strconv.Atoi(strings.TrimSpace(string(buf)[0:2]))
		j,err:= strconv.Atoi(strings.TrimSpace(string(buf)[3:]))
		player.Player2Pos = [2]int{i,j}
	}
}

func integerToStringOfFixedWidth(n, w int) string {
    s := fmt.Sprintf(fmt.Sprintf("%%0%dd", w), n)
    l := len(s)
    if l > w {
        return s[l-w:]
    }
    return s
}

