package main

import (
	tm "github.com/buger/goterm"
	"github.com/nsf/termbox-go"
)

var running bool = true

func command()  {
	for running {
		ev := termbox.PollEvent()

		switch ev.Type {
		case termbox.EventKey:
			switch ev.Ch {
			case 'q':
				running = false
			case 'p':
				dataDraft.addNextShip()
			case 'd':
				dataDraft.moveLastShip(1, 0)
			case 's':
				dataDraft.moveLastShip(0, 1)
			case 'a':
				dataDraft.moveLastShip(-1, 0)
			case 'w':
				dataDraft.moveLastShip(0, -1)
			}
		}
	}
}

func run() {
	termbox.Init()
	go command()

	tm.Clear()
	for running {
		render()
	}
}
