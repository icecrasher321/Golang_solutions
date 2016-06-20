package main
// URL -> https://tour.golang.org/moretypes/18
import "golang.org/x/tour/pic"
import "math/rand"

func Pic(dy, dx int) [][]uint8 {
	rows := make([][]uint8, dy)
	row := make([]uint8, dx)
	for i := 0; i < dx; i++ {
		row[i] = uint8(rand.Intn(256-i))
	}
	for i := 0; i < dy; i++ {
		rows[i] = row
	}
	return rows
}

func main() {
	pic.Show(Pic)
}
