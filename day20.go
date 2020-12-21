package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type PuzzlePiece struct {
	TileId int
	Shape  []string
	Type   string
}

func RotatePiece(tile PuzzlePiece) PuzzlePiece {
	shape := tile.Shape
	s := []string{}

	for i := range shape {
		l := ""
		for j := range shape {
			l += string(shape[len(shape)-j-1][i])
		}
		s = append(s, l)
	}

	return PuzzlePiece{tile.TileId, s, tile.Type}
}

func FlipPieceVertical(tile PuzzlePiece) PuzzlePiece {
	shape := tile.Shape
	s := []string{}

	for i := range shape {
		l := ""
		for j := range shape {
			l += string(shape[i][len(shape)-j-1])
		}
		s = append(s, l)
	}

	return PuzzlePiece{tile.TileId, s, tile.Type}
}

func FlipPieceHorizontal(tile PuzzlePiece) PuzzlePiece {
	shape := tile.Shape
	s := []string{}

	for i := range shape {
		l := ""
		for j := range shape {
			l += string(shape[len(shape)-i-1][j])
		}
		s = append(s, l)
	}

	return PuzzlePiece{tile.TileId, s, tile.Type}
}

func RightLeftBorderValid(left PuzzlePiece, right PuzzlePiece) bool {
	left_border := ""
	right_border := ""

	for _, line := range left.Shape {
		left_border += string(line[len(left.Shape)-1])
	}
	for _, line := range right.Shape {
		right_border += string(line[0])
	}

	return left_border == right_border
}

func BottomTopBorderValid(top PuzzlePiece, bottom PuzzlePiece) bool {
	top_border := top.Shape[len(top.Shape)-1]
	bottom_border := bottom.Shape[0]

	return top_border == bottom_border
}

func IsValidPuzzle(board map[int]PuzzlePiece) bool {
	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			current_tile, f := board[y*width+x]
			if !f {
				continue
			}

			if x < width-1 {
				right_tile, f2 := board[y*width+x+1]
				if f2 && !RightLeftBorderValid(current_tile, right_tile) {
					return false
				}
			}

			if y > 0 {
				top_tile, f2 := board[(y-1)*width+x]
				if f2 && !BottomTopBorderValid(top_tile, current_tile) {
					return false
				}
			}
		}
	}

	return true
}

func PrintPuzzle(board map[int]PuzzlePiece) {
	row := ""
	for y := 0; y < width; y++ {
		line := ""
		for k := 0; k < len(board[0].Shape); k++ {
			for x := 0; x < width; x++ {
				current_tile, f := board[y*width+x]
				if f {
					line += current_tile.Shape[k] + " "
				} else {
					line += "           "
				}
			}
			line += "\n"
		}
		row += line + "\n"
	}

	// print final result
	// fmt.Println(row)
	ioutil.WriteFile("./day20.output", []byte(row), 0644)
}

func TrimPuzzle(board map[int]PuzzlePiece) []string {
	trimmed_board := []string{}

	for y := 0; y < width; y++ {
		for k := 1; k < len(board[0].Shape)-1; k++ {
			line := ""
			for x := 0; x < width; x++ {
				current_tile, f := board[y*width+x]
				if f {
					l := current_tile.Shape[k]
					line += l[1 : len(l)-1]
				}
			}
			trimmed_board = append(trimmed_board, line)
		}
	}

	return trimmed_board
}

func FindMonsters(monster []string, completed_puzzle []string) int {
	total_found := 0

	for y := 0; y < len(completed_puzzle)-len(monster); y++ {
		for x := 0; x < len(completed_puzzle[0])-len(monster[0]); x++ {
			found := true
			for my := 0; my < len(monster); my++ {
				for mx := 0; mx < len(monster[0]); mx++ {
					m := monster[my][mx]
					cpp := completed_puzzle[y+my][x+mx]
					if string(m) == "#" && string(cpp) != "#" {
						found = false
						break
					}
				}
				if !found {
					break
				}
			}
			if found {
				total_found++
			}
		}
	}

	return total_found
}

func IdentifyPuzzlePieces(tiles map[int]PuzzlePiece) map[int]PuzzlePiece {
	m := []int{}

	for i, tile1 := range tiles {
		matches := 0
		for k, tile2 := range tiles {
			if i == k {
				continue
			}

			try_pieces1 := []PuzzlePiece{
				tile1,
				RotatePiece(RotatePiece(FlipPieceVertical(tile1))), RotatePiece(RotatePiece(RotatePiece(FlipPieceVertical(tile1)))),
				FlipPieceVertical(tile1), RotatePiece(FlipPieceVertical(tile1)),
				RotatePiece(tile1), RotatePiece(RotatePiece(tile1)), RotatePiece(RotatePiece(RotatePiece(tile1))),
			}
			try_pieces2 := []PuzzlePiece{
				tile2,
				RotatePiece(RotatePiece(FlipPieceVertical(tile2))), RotatePiece(RotatePiece(RotatePiece(FlipPieceVertical(tile2)))),
				FlipPieceVertical(tile2), RotatePiece(FlipPieceVertical(tile2)),
				RotatePiece(tile2), RotatePiece(RotatePiece(tile2)), RotatePiece(RotatePiece(RotatePiece(tile2))),
			}
			found := false
			for _, p1 := range try_pieces1 {
				for _, p2 := range try_pieces2 {
					if RightLeftBorderValid(p1, p2) || BottomTopBorderValid(p1, p2) {
						matches++
						found = true
						break
					}
				}
				if found {
					break
				}
			}
		}
		if matches == 2 {
			tile1.Type = "corner"
		} else if matches == 3 {
			tile1.Type = "edge"
		} else if matches == 4 {
			tile1.Type = "center"
		}

		tiles[tile1.TileId] = tile1

		if matches < 3 {
			m = append(m, tile1.TileId)
		}
	}

	return tiles
}

func PutPuzzleTogether(tiles map[int]PuzzlePiece, board map[int]PuzzlePiece, next_spot int, unused_pieces []int) (bool, map[int]PuzzlePiece) {
	if next_spot == len(tiles) && IsValidPuzzle(board) {
		return true, board
	}

	x := next_spot % width
	y := next_spot / width
	should_be_corner := next_spot == 0 || next_spot == width-1 || next_spot == width*width-width || next_spot == width*width-1
	should_be_edge := (x == 0 || x == width-1 || y == 0 || y == width-1) && !should_be_corner

	for i, tile_id := range unused_pieces {
		if tile_id == -1 {
			continue
		}

		tile := tiles[tile_id]
		if should_be_corner && tile.Type != "corner" {
			continue
		}
		if should_be_edge && tile.Type != "edge" {
			continue
		}
		if (!should_be_corner && !should_be_edge) && tile.Type != "center" {
			continue
		}

		try_pieces := []PuzzlePiece{
			tile,
			RotatePiece(tile), RotatePiece(RotatePiece(tile)), RotatePiece(RotatePiece(RotatePiece(tile))),
			FlipPieceVertical(tile), RotatePiece(FlipPieceVertical(tile)),
			RotatePiece(RotatePiece(FlipPieceVertical(tile))), RotatePiece(RotatePiece(RotatePiece(FlipPieceVertical(tile)))),
		}

		for _, piece := range try_pieces {
			board[next_spot] = piece
			unused_pieces[i] = -1

			if IsValidPuzzle(board) {
				found, board := PutPuzzleTogether(tiles, board, next_spot+1, unused_pieces)
				if found {
					return true, board
				}
			}
		}

		// set the tile id back if we dont end up using it
		delete(board, next_spot)
		unused_pieces[i] = tile_id
	}

	return false, nil
}

var width int

func day20() {
	inp, _ := ioutil.ReadFile("./inputs/day20.input")

	data := GetStringInput(inp)

	tiles := make(map[int]PuzzlePiece)
	tile_ids := []int{}
	tile_num := 0
	shape := []string{}
	for _, line := range data {
		if strings.Contains(line, "Tile") {
			tile_num, _ = strconv.Atoi(strings.Trim(strings.Split(line, " ")[1], ":"))
			tile_ids = append(tile_ids, tile_num)
			shape = []string{}
		} else if strings.Contains(line, ".") || strings.Contains(line, "#") {
			shape = append(shape, line)
		} else {
			tiles[tile_num] = PuzzlePiece{tile_num, shape, ""}
		}

	}

	// input vs sample
	if len(tiles) > 50 {
		width = 12
	} else {
		width = 3
	}

	// marks every tile with "corner", "edge", or "center"
	tiles = IdentifyPuzzlePieces(tiles)

	p := 1

	for tid, t := range tiles {
		if t.Type == "corner" {
			p *= tid
		}
	}
	fmt.Println("corner product:", p)

	board := make(map[int]PuzzlePiece)
	found, return_board := PutPuzzleTogether(tiles, board, 0, tile_ids)

	if !found {
		return
	}

	cb := TrimPuzzle(return_board)

	// turn monster into a string mask
	// turn the completed board back into a tile (which can be rotated/flipped)
	// check all flipped/rotated boards for monsters
	monster := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	cb_tile := PuzzlePiece{-1, cb, ""}
	cbs := []PuzzlePiece{
		cb_tile,
		RotatePiece(cb_tile), RotatePiece(RotatePiece(cb_tile)), RotatePiece(RotatePiece(RotatePiece(cb_tile))),
		FlipPieceVertical(cb_tile), RotatePiece(FlipPieceVertical(cb_tile)),
		RotatePiece(RotatePiece(FlipPieceVertical(cb_tile))), RotatePiece(RotatePiece(RotatePiece(FlipPieceVertical(cb_tile)))),
	}

	for _, cb := range cbs {
		monsters_found := FindMonsters(monster, cb.Shape)
		water_roughness := 0
		for _, l := range cb.Shape {
			water_roughness += strings.Count(l, "#")
		}
		if monsters_found > 0 {
			fmt.Println("found monsters:", monsters_found, "waters roughness is", water_roughness-15*monsters_found)
			break
		}
	}
}
