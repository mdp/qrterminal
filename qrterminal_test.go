package qrterminal

import (
	"os"
	"testing"

	"github.com/mdp/qrterminal"
)

func TestGenerate(t *testing.T) {
	Generate("https://github.com/mdp/go-qrcode/terminal", qrterminal.L, os.Stdout)
}
