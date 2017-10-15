package main


var lastChip string = "single"

func allowAdd(ship string, ships [][]int) bool {
	switch ship {
	case "single":
		return len(ships) < 4
	case "double":
		return len(ships) < 3
	case "triple":
		return len(ships) < 2
	case "quadruple":
		return len(ships) < 1
	}

	return false
}

func (dd* DataDraft) updatePoints (ship string, point []int, draw string) {
	switch ship {
	case "single":
		dd.drawPoints[point[0]][point[1]] = draw
	case "double":
		dd.drawPoints[point[0]][point[1]] = draw
		dd.drawPoints[point[2]][point[3]] = draw
	case "triple":
		dd.drawPoints[point[0]][point[1]] = draw
		dd.drawPoints[point[2]][point[3]] = draw
		dd.drawPoints[point[4]][point[5]] = draw
	case "quadruple":
		dd.drawPoints[point[0]][point[1]] = draw
		dd.drawPoints[point[2]][point[3]] = draw
		dd.drawPoints[point[4]][point[5]] = draw
		dd.drawPoints[point[6]][point[7]] = draw
	}
}

func (dd* DataDraft) addShip(ship string) {
	if !allowAdd(ship, dd.ships[ship]) {
		return
	}

	point := []int{}
	switch ship {
	case "single":
		point = []int{0,0}
	case "double":
		point = []int{0,0,0,1}
	case "triple":
		point = []int{0,0,0,1,0,2}
	case "quadruple":
		point = []int{0,0,0,1,0,2,0,3}
	}

	dd.updatePoints(ship, point, dataObj.ShipPoint)

	dd.ships[ship] = append(dd.ships[ship], point)

	lastChip = ship
}

func (dd* DataDraft) moveShip(ship string, nShip int, xStep, yStep int) {
	if (nShip + 1) > len(dd.ships[ship]) {
		return
	}

	point := []int{}
	size  := len(dd.ships[ship][nShip])

	for i := 0; i < size; i ++ {
		pointShip := dd.ships[ship][nShip][i]

		if i % 2 == 0 {
			pointShip = pointShip + xStep
		} else {
			pointShip = pointShip + yStep
		}

		point = append(point, pointShip)
	}

	dd.updatePoints(ship, dd.ships[ship][nShip], dataObj.Space)

	dd.ships[ship][nShip] = point

	dd.updatePoints(ship, point, dataObj.ShipPoint)
}

func (dd* DataDraft) addNextShip() {
	if allowAdd("single", dd.ships["single"]) {
		dd.addShip("single")
	} else if allowAdd("double", dd.ships["double"]) {
		dd.addShip("double")
	} else if allowAdd("triple", dd.ships["triple"]) {
		dd.addShip("triple")
	} else if allowAdd("quadruple", dd.ships["quadruple"]) {
		dd.addShip("quadruple")
	}
}

func (dd* DataDraft) moveLastShip(xStep, yStep int)  {
	n := len(dd.ships[lastChip]) - 1

	if n >= 0 {
		dd.moveShip(lastChip, n, xStep, yStep)
	}
}
