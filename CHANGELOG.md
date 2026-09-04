## 4.0.0

**BREAKING:** Module path is now `github.com/mdp/qrterminal/v4`. Update the import path, `go get`, and `go install` references when upgrading from v3.

### Fixes & Improvements

- Reworked Sixel terminal-state handling so it works correctly with newer `golang.org/x/term` versions and reliably restores the terminal on exit (#40, fixed by @wackerm)
- Fixed a nil-pointer panic in `IsSixelSupported` on Windows Terminal / Git Bash caused by an unhandled `term.MakeRaw` error (#35, fixed by @cbednarski; incorporated into the rewritten handler above)
- Removed a stray `oryxBuildBinary` build artifact that was accidentally committed to the repository (#36, spotted by @zonescape)
- Added CI: build, vet, and tests now run on every pull request and on pushes to `main`, across Ubuntu, macOS, and Windows with Go 1.25 and 1.26 (#41)
- Updated dependencies: `golang.org/x/term` v0.43.0, `golang.org/x/sys` v0.47.0, `go-colorable` v0.1.15
- Raised the Go directive in `go.mod` from 1.20 to 1.25.0 (the minimum required by the updated dependencies)

Thanks to [@wackerm](https://github.com/wackerm), [@cbednarski](https://github.com/cbednarski), and [@zonescape](https://github.com/zonescape) for their contributions to this release.

## 3.2.1

- Fix #33 - Default config to standard characters if not specified.

## 3.2.0

- Update to add sixel support #29
- Update deps to latest

## 3.1.1

- Update deps to latest

## 3.1.0

- Add the ability to accept input string from stdin
- Integrate github actions for build and release
- Release support for Darwin M1/M2(aarch64)

## 3.0.0

Adjust go.mod to include required version string

## 2.0.1

Add goreleaser and release to Homebrew and Github

## 2.0.0

Add a command line tool and QuietZone around QRCode

## 1.0.1

Add go.mod

## 1.0.0

Update to add a quiet zone border to the QR Code - #5 and fixed by [WindomZ](https://github.com/WindomZ) #8

  - This can be configured with the `QuietZone int` option
  - Defaults to 4 'pixels' wide to match the QR Code spec
  - This alters the size of the barcode considerably and is therefore a breaking change, resulting in a bump to v1.0.0

## 0.2.1 

Fix direction of the qr code #6 by (https://github.com/mattn)
