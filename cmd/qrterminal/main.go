package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/mdp/qrterminal/v4"
	"rsc.io/qr"
)

var verboseFlag bool
var levelFlag string
var quietZoneFlag int
var sixelFlag bool

func getLevel(s string) qr.Level {
	switch l := strings.ToLower(s); l {
	case "l":
		return qr.L
	case "m":
		return qr.M
	case "h":
		return qr.H
	default:
		return -1
	}
}

func main() {
	flag.BoolVar(&verboseFlag, "v", false, "Output debugging information")
	flag.StringVar(&levelFlag, "l", "L", "Error correction level")
	flag.IntVar(&quietZoneFlag, "q", 2, "Size of quietzone border")
	flag.BoolVar(&sixelFlag, "sixel", false, "enable sixel format for output (opt-in)")

	// Allow flags to appear before, after, or between operands, e.g.
	// `qrterminal "hello" -q 5`. Go's flag package normally stops parsing at
	// the first non-flag argument, so reorder: flags first, operands last.
	// Flags that take a value are consumed with their value.
	valueFlags := map[string]bool{"l": true, "q": true}
	var flags, operands []string
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") && arg != "-" && arg != "--" {
			name := strings.TrimLeft(arg, "-")
			flags = append(flags, arg)
			// `--flag=value` or boolean flags carry their value already;
			// otherwise the next argument is this flag's value.
			if valueFlags[name] && !strings.Contains(arg, "=") && i+1 < len(args) {
				i++
				flags = append(flags, args[i])
			}
		} else {
			operands = append(operands, arg)
		}
	}
	os.Args = append(os.Args[:1], append(flags, operands...)...)

	flag.Parse()
	level := getLevel(levelFlag)
	content := strings.Join(flag.Args(), " ")

	if len(content) < 1 {
		// Get input from stdin until EOF
		stdin, err := io.ReadAll(os.Stdin)

		if err != nil {
			panic(err)
		}
		content = string(stdin)
	} else if level < 0 {
		fmt.Fprintf(os.Stderr, "Invalid error correction level: %s\n", levelFlag)
		fmt.Fprintf(os.Stderr, "Valid options are [L, M, H]\n")
		os.Exit(1)
	}

	cfg := qrterminal.Config{
		Level:     level,
		Writer:    os.Stdout,
		QuietZone: quietZoneFlag,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
	}
	if sixelFlag {
		cfg.WithSixel = qrterminal.IsSixelSupported(os.Stdout)
	}
	if verboseFlag {
		fmt.Fprintf(os.Stdout, "Level: %s \n", levelFlag)
		fmt.Fprintf(os.Stdout, "Quietzone Border Size: %d \n", quietZoneFlag)
		fmt.Fprintf(os.Stdout, "Encoded data: %s \n", strings.Join(flag.Args(), "\n"))
		fmt.Println("")
	}

	if runtime.GOOS == "windows" {
		cfg.Writer = colorable.NewColorableStdout()
		cfg.BlackChar = qrterminal.BLACK
		cfg.WhiteChar = qrterminal.WHITE
	}

	fmt.Fprint(os.Stdout, "\n")
	qrterminal.GenerateWithConfig(content, cfg)
}
