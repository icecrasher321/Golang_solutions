package main
// URL -> https://tour.golang.org/methods/22
import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(s []byte) (int, error) { // Fills up the array
	i := 0
	for {
		s[i] = 'A'
		i += 1
		if i == (len(s)) {
			break
		}
	}
	   return 1, nil
}


func main() {
  reader.Validate(MyReader{})
}
