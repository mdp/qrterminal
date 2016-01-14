package terminal

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	Generate("https://github.com/mdp/go-qrcode/terminal", os.Stdout)
}
