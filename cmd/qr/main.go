package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/mdp/qrterminal"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "usage of %s: [arguments]\n", os.Args[0])
		os.Exit(1)
	}
	cfg := qrterminal.Config{
		Level:          qrterminal.L,
		HalfBlocks:     true,
		Writer:         os.Stdout,
		BlackChar:      qrterminal.BLACK_BLACK,
		WhiteBlackChar: qrterminal.WHITE_BLACK,
		WhiteChar:      qrterminal.WHITE_WHITE,
		BlackWhiteChar: qrterminal.BLACK_WHITE,
	}
	if runtime.GOOS == "windows" {
		cfg.HalfBlocks = false
		cfg.Writer = colorable.NewColorableStdout()
		cfg.BlackChar = qrterminal.BLACK
		cfg.WhiteChar = qrterminal.WHITE
	}

	qrterminal.GenerateWithConfig(strings.Join(os.Args[1:], "\n"), cfg)
}
