package qrterminal

import (
	"io"

	"github.com/mdp/rsc/qr"
)

const BLACK_WHITE = "▄"
const BLACK_BLACK = " "
const WHITE_BLACK = "▀"
const WHITE_WHITE = "█"

// Level - the QR Code's redundancy level
const H = qr.H
const M = qr.M
const L = qr.L

//Config for generating a barcode
type Config struct {
	Level          qr.Level
	Writer         io.Writer
	BlackBlackChar string
	BlackWhiteChar string
	WhiteWhiteChar string
	WhiteBlackChar string
}

// GenerateWithConfig expects a string to encode and a config
func GenerateWithConfig(text string, config Config) {
	w := config.Writer
	ww := config.WhiteWhiteChar
	bb := config.BlackBlackChar
	wb := config.WhiteBlackChar
	bw := config.BlackWhiteChar
	code, _ := qr.Encode(text, config.Level)
	// Frame the barcode in a 1 pixel border
	w.Write([]byte(bw))
	for i := 0; i <= code.Size; i++ {
		w.Write([]byte(bw))
	}
	w.Write([]byte("\n"))

	for i := 0; i <= code.Size; i += 2 {
		w.Write([]byte(ww))
		for j := 0; j <= code.Size; j++ {
			next_black := false
			if i+1 < code.Size {
				next_black = code.Black(i+1, j)
			}
			curr_black := code.Black(i, j)
			if curr_black && next_black {
				w.Write([]byte(bb))
			} else if curr_black && !next_black {
				w.Write([]byte(bw))
			} else if !curr_black && !next_black {
				w.Write([]byte(ww))
			} else {
				w.Write([]byte(wb))
			}
		}
		w.Write([]byte("\n"))
	}
}

// Generate a QR Code and write it out to io.Writer
func Generate(text string, l qr.Level, w io.Writer) {
	config := Config{
		Level:          l,
		Writer:         w,
		BlackBlackChar: BLACK_BLACK,
		BlackWhiteChar: BLACK_WHITE,
		WhiteWhiteChar: WHITE_WHITE,
		WhiteBlackChar: WHITE_BLACK,
	}
	GenerateWithConfig(text, config)
}
