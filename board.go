package main

import (
	"strings"
	"strconv"
)


func IfThenElse(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}

func genLine(length int, firstItem, lastItem, startSym, endSym, lineSym, separateSym string) []string {
	line := []string{}

	for i := 1; i <= length; i++ {
		sym := dataObj.Space

		if i == 1 {
			sym = firstItem
		} else if i == SPACE_ROW {
			sym = lastItem
		} else if i > SPACE_ROW {
			tmp := SPACE_ROW + 1

			if i % tmp == 0 {
				switch i {
				case tmp:
					sym = startSym
				case length:
					sym = endSym
				default:
					sym = separateSym
				}
			} else {
				sym = lineSym
			}
		}

		line = append(line, sym)
	}

	return line
}

func genSymLine(length int) []string {
	line := []string{}

	sym_i := 0
	for i := 0; i < length; i++ {
		if i > 2 && i % (SPACE_ROW + 1) == 1 {
			line = append(line, string(ALPHABET[sym_i]))
			sym_i++
		} else {
			line = append(line, dataObj.Space)
		}
	}

	return line
}

func generateBlankBoard(w, h int) [][]string {
	line_length := (w * SPACE_ROW) + w + (SPACE_ROW + 1)
	line := genSymLine(line_length)

	var board = [][]string {}
	board = append(board, line)

	for y := 0; y < h; y++ {
		line1 := []string{}

		if y == 0 {
			line1 = genLine(
				line_length,
				dataObj.Space, dataObj.Space,
				dataObj.TopLeftAngle, dataObj.TopRightAngle,
				dataObj.HorizontalLine, dataObj.TopCrossroads)
		} else {
			line1 = genLine(
				line_length,
				dataObj.Space, dataObj.Space,
				dataObj.LeftCrossroads, dataObj.RightCrossroads,
				dataObj.HorizontalLine, dataObj.Crossroads)
		}

		item := strings.Join(
			[]string{IfThenElse(y < 9, dataObj.Space, dataObj.Blank),
				strconv.Itoa(y + 1)}, "")

		line2 := genLine(
			line_length,
			item, dataObj.Blank,
			dataObj.VerticalLine, dataObj.VerticalLine,
			dataObj.Space, dataObj.VerticalLine)

		board = append(board, line1, line2)

		if y == (h-1) {
			line3 := genLine(
				line_length,
				dataObj.Space, dataObj.Space,
				dataObj.BottomLeftAngle, dataObj.BottomRightAngle,
				dataObj.HorizontalLine, dataObj.BottomCrossroads)

			board = append(board, line3)
		}
	}

	return board
}
