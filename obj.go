package main


const SPACE_ROW int = 3
const ALPHABET string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const sizeBoardW, sizeBoardH int = 10, 10

type DataObj struct {
	Blank string
	Space string
	TopLeftAngle string
	TopRightAngle string
	BottomLeftAngle string
	BottomRightAngle string
	HorizontalLine string
	VerticalLine string
	Crossroads string
	LeftCrossroads string
	RightCrossroads string
	TopCrossroads string
	BottomCrossroads string
	ShipPoint string
	Miss string
	Destroy string
}

type DataDraft struct {
	board [][]string
	drawPoints [sizeBoardW][sizeBoardH]string
	ships map[string][][]int
}
