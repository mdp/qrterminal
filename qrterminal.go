package qrterminal

import (
	"io"

	"github.com/mdp/rsc/qr"
)

const BLACK = "\033[40m  \033[0m"
const WHITE = "\033[47m  \033[0m"

// Level - the QR Code's redundancy level
type Level int

const H = qr.H
const M = qr.M
const L = qr.L

// Generate a QR Code and write it out to io.Writer
func Generate(text string, l qr.Level, w io.Writer) {
	code, _ := qr.Encode(text, l)
	// Frame the barcode in a 1 pixel border
	w.Write([]byte(BLACK))
	for i := 0; i <= code.Size; i++ {
		w.Write([]byte(BLACK))
	}
	w.Write([]byte("\n"))
	for i := 0; i <= code.Size; i++ {
		w.Write([]byte(BLACK))
		for j := 0; j <= code.Size; j++ {
			if code.Black(i, j) {
				w.Write([]byte(WHITE))
			} else {
				w.Write([]byte(BLACK))
			}
		}
		w.Write([]byte("\n"))
	}
}
