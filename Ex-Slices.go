package main

// URL -> https://tour.golang.org/moretypes/18


import "golang.org/x/tour/pic"
func Pic(dy, dx int) [][]uint8 {
	rows := make([][]uint8, dy)
	row := make([]uint8, dx)
	for i := 0; i < dx; i++ {
		row[i] = uint8(255-i) //gradient (high to low)
	}
	for i := 0; i < dy; i++ {
		rows[i] = row
	}
	return rows
}

func main() {
	pic.Show(Pic)
}
