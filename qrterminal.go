package terminal

import (
	"io"

	"code.google.com/p/rsc/qr"
)

const BLACK = "\033[40m  \033[0m"
const WHITE = "\033[47m  \033[0m"

func Generate(text string, w io.Writer) {
	code, _ := qr.Encode(text, qr.L)
	for i := 0; i <= code.Stride; i++ {
		for j := 0; j <= code.Stride; j++ {
			if code.Black(i, j) {
				w.Write([]byte(WHITE))
			} else {
				w.Write([]byte(BLACK))
			}
		}
		w.Write([]byte("\n"))
	}
}
