# QRCode Terminal

[![Build Status](https://secure.travis-ci.org/mdp/qrterminal.png)](https://travis-ci.org/mdp/qrterminal)

A golang library for generating QR codes in the terminal.

Originally this was a port of the [NodeJS version](https://github.com/gtanner/qrcode-terminal). Recently it's been updated to allow for smaller code generation using ASCII 'half blocks'

## Example
Full size ASCII block QR Code:  
<img src="https://user-images.githubusercontent.com/2868/35941992-33974eec-0c22-11e8-867c-234d3d06f016.png" alt="alt text" width="200" height="200">

Smaller 'half blocks' in the terminal:  
<img src="https://user-images.githubusercontent.com/2868/35942180-d11b565e-0c22-11e8-8df9-481cd1b7e7b3.png" alt="alt text" width="200" height="200">

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
  qrterminal.Generate("https://github.com/mdp/qrterminal", qrterminal.L, os.Stdout)
}
```

### More complicated

Large Inverted barcode with medium redundancy and a 1 pixel border
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
      QuietZone: 1,
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
[WindomZ](https://github.com/WindomZ)  
[mattn](https://github.com/mattn)  
