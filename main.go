package main


var dataObj = DataObj{
	"", " ",
	"╔", "╗",
	"╚", "╝",
	"═", "║",
	"╬",
	"╠", "╣",
	"╦", "╩",
	"█", "Ø", "×",
}

var dataDraft DataDraft = DataDraft{
	board: generateBlankBoard(sizeBoardW, sizeBoardH),
	ships: make(map[string][][]int),
}


func main() {
	run()
}
