# QRCode Terminal

Pretty simple, I stole this from https://github.com/gtanner/qrcode-terminal

## Install

`go get github.com/mdp/qrterminal`

## Usage

```go
import (
    "github.com/mdp/qrterminal"
    "os"
    )

func main() {
  qrterminal.Generate("https://github.com/mdp/qrterminal", os.Stdout)
}
```

### More complicated

Inverted barcode with medium redundancy
```go
import (
    "github.com/mdp/qrterminal"
    "os"
    )

func main() {
  config := qrterminal.Config{
      Level: qrterminal.L,
      Writer: os.Stdout,
      BlackChar: qrterminal.WHITE,
      WhiteChar: qrterminal.BLACK,
  }
  qrterminal.GenerateWithConfig("https://github.com/mdp/qrterminal", config)
}
```

