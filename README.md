# QRCode Terminal

[![Build Status](https://secure.travis-ci.org/mdp/qrterminal.png)](https://travis-ci.org/mdp/qrterminal)

A golang library for generating QR codes in the terminal.

Originally this was a port of the [NodeJS version](https://github.com/gtanner/qrcode-terminal). Recently it's be updated to allow for smaller code generation using ASCII 'half blocks'

## Install

`go get github.com/mdp/qrterminal`

## Usage

```go
import (
    "github.com/mdp/qrterminal"
    "os"
    )

func main() {
  // Generate a 'dense' qrcode with the 'Low' level error correction and write it to Stdout
  qrterminal.GenerateHalfBlock("https://github.com/mdp/qrterminal", qrterminal.L, os.Stdout)
}
```

### More complicated

Large Inverted barcode with medium redundancy
```go
import (
    "github.com/mdp/qrterminal"
    "os"
    )

func main() {
  config := qrterminal.Config{
      Level: qrterminal.M,
      Writer: os.Stdout,
      BlackChar: qrterminal.WHITE,
      WhiteChar: qrterminal.BLACK,
  }
  qrterminal.GenerateWithConfig("https://github.com/mdp/qrterminal", config)
}
```

HalfBlock barcode with medium redundancy
```go
import (
    "github.com/mdp/qrterminal"
    "os"
    )

func main() {
  config := qrterminal.Config{
      HalfBlocks: true,
      Level: qrterminal.M,
      Writer: os.Stdout,
  }
  qrterminal.Generate("https://github.com/mdp/qrterminal", config)
}
```

Credits:

Mark Percival m@mdp.im
[Matthew Kennerly](https://github.com/mtkennerly)
[Viric](https://github.com/viric)

