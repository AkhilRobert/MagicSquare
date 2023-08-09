package main

import (
	"fmt"
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
			fmt.Printf("  %d  ", list[i][j])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Print("Enter the size of the square : ")
	var input string
	fmt.Scanln(&input)

	size, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic("something went wrong when parsing the input")
	}

	if size%2 == 0 {
		panic("even order square are not supported")
	}

	list := make([][]int, size)
	for i := range list {
		list[i] = make([]int, size)
	}

	buildSquare(&list, len(list))
	printSquare(list)
}
