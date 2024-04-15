package mapmodule

import (
	structs "github.com/MultiplayerObsGame/Structs"
)


func GenMap(){
	GenFloor()
	// for i:=0;i<30;i++{
	// 	for j:=0;j<100;j++{
	// 		structs.VisibleMatrix[i][j].IsVisible = true
	// 	}
	// }
}


func GenFloor(){
	for j:=0;j<100;j++{
		structs.VisibleMatrix[29][j].IsFloor = true
		structs.VisibleMatrix[29][j].IsVisible = true
	}
}
