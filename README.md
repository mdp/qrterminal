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
  qrterminal.Generate("https://github.com/mdp/go-qrcode/terminal", os.Stdout)
}
```

