package main

import (
	"fmt"
	"os"
	"strconv"
)

func emod(value, modulo int) int {
	return ((value % modulo) + modulo) % modulo
}

func insert(value, row, col int, list *[][]int) (int, int) {
	dimension := len(*list)

	if row < 0 || col >= dimension || row >= dimension {
		row = emod(row, dimension)
		col = emod(col, dimension)
	}

	if (*list)[row][col] != 0 {
		return insert(value, row+2, col-1, list)
	}

	(*list)[row][col] = value
	return row, col
}

func buildSquare(list *[][]int, dimension int) {
	values := make([]int, dimension*dimension)
	for i := range values {
		values[i] = i + 1
	}

	middleIdx := dimension / 2
	(*list)[0][middleIdx] = values[0]

	row := -1
	col := middleIdx + 1

	for i := 1; i < dimension*dimension; i++ {
		insertedRow, insertedCol := insert(values[i], row, col, list)
		row = insertedRow - 1
		col = insertedCol + 1
	}
}

func printSquare(list [][]int) {
	for i, row := range list {
		for j := range row {
			fmt.Printf("%-6d", list[i][j])
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, os.Args[0]+": square size not specified")
		os.Exit(1)
	}

	size, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic("something went wrong when parsing the square size")
	}

	if size <= 2 {
		fmt.Fprintln(os.Stderr, "ERROR: square size should be greater than 2")
		os.Exit(1)
	}

	if size%2 == 0 {
		fmt.Fprintln(os.Stderr, "ERROR: only odd size squares are supported")
		os.Exit(1)
	}

	list := make([][]int, size)
	for i := range list {
		list[i] = make([]int, size)
	}

	buildSquare(&list, len(list))
	printSquare(list)
}
