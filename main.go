package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func next(n, d int) int {
	n = n + 1
	if n == d {
		n = n - d
	}
	return n
}

func prev(n, d int) int {
	n = n - 1
	if n == -1 {
		n = d - 1
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid arguments")
		return
	}
	d, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create d x d array
	square := make([][]int, d)
	for i := range square {
		square[i] = make([]int, d)
	}

	var n = 1
	var xx = 0
	var yy = d / 2
	square[xx][yy] = n
	for {
		var x1 = prev(xx, d)
		var y1 = next(yy, d)
		if square[x1][y1] == 0 {
			xx = x1
			yy = y1
			n = n + 1
			square[xx][yy] = n
		} else {
			y1 = yy
			x1 = next(xx, d)
			if square[x1][y1] == 0 {
				xx = x1
				yy = y1
				n = n + 1
				square[xx][yy] = n
			} else {
				break
			}
		}
	}

	// Print the magic-square
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.SetColumnSeparator("|")
	table.SetCenterSeparator("+")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, v := range square {
		var row = make([]string, d)
		for i, cell := range v {
			row[i] = strconv.Itoa(cell)
		}
		table.Append(row)
	}
	table.Render()
}
