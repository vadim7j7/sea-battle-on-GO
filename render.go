package main

import (
	tm "github.com/buger/goterm"
	"strings"
)


func renderField()  {
	size := len(dataDraft.board)
	for i := 0; i < size; i++ {
		tm.Println(strings.Join(dataDraft.board[i], ""))
	}
}

func drawObjectsOnBoard()  {
	iSize := len(dataDraft.board)

	var y int = 0
	for i := 1; i < iSize; i++ {
		if i % 2 == 0 {
			jSize := len(dataDraft.board[i])
			var x int = -1
			for j := SPACE_ROW; j < jSize; j++ {
				sp := j % 4

				if sp == 0 {
					x++
				}

				if sp != 3 && dataDraft.drawPoints[x][y] != "" {
					point := dataDraft.drawPoints[x][y]

					switch point {
						case dataObj.Miss, dataObj.Destroy:
							if sp == 1 {
								dataDraft.board[i][j] = point
							}
						default:
							dataDraft.board[i][j] = point
					}
				}
			}

			y++
		}
	}
}

func render()  {
	tm.MoveCursor(1,1)

	drawObjectsOnBoard()
	renderField()

	tm.Println("To exit to press key 'Q'")
	tm.Println("To add a ship to press key 'P'")
	tm.Println("To move the ship to press keys 'A,D,W,S'")

	tm.Flush()
}
