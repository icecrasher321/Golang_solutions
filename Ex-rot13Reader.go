package main
// URL -> https://tour.golang.org/methods/23
import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r1 rot13Reader) Read(b []byte) (i int, err error) {
	i, err = r1.r.Read(b)
	for i := 0; i < len(b); i++ {
		if (b[i] > 78 && b[i] <= 91) || (b[i] >= 110 && b[i] <= 122) {
			b[i] = b[i] - 13
		} else if (b[i] >= 65 && b[i] < 78) || (b[i] >= 97 && b[i] < 110) {
			b[i] = b[i] + 13
		}
	}

	return i, err
}

/* 78 -> N,
   91 -> Z,
	 65 -> A,
	 110 -> n,
	 122 -> z,
	  97 -> a */

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
