package qrterminal

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	Generate("https://github.com/mdp/qrterminal", L, os.Stdout)
}

func TestGenerateWithConfig(t *testing.T) {
	config := Config{
		Level:          M,
		Writer:         os.Stdout,
		HalfBlocks:     true,
		BlackChar:      BLACK_BLACK,
		WhiteBlackChar: WHITE_BLACK,
		WhiteChar:      WHITE_WHITE,
		BlackWhiteChar: BLACK_WHITE,
	}
	GenerateWithConfig("https://github.com/mdp/qrterminal", config)
}
