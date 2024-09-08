# log
[![Go Reference](https://pkg.go.dev/badge/github.com/bartventer/log.svg)](https://pkg.go.dev/github.com/bartventer/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/bartventer/log)](https://goreportcard.com/report/github.com/bartventer/log)
[![Test](https://github.com/bartventer/log/actions/workflows/default.yml/badge.svg)](https://github.com/bartventer/log/actions/workflows/default.yml)
[![codecov](https://codecov.io/gh/bartventer/log/graph/badge.svg?token=btaj9v5KWM)](https://codecov.io/gh/bartventer/log)

`log` provides a structured logger with context support.

![Made with VHS](https://vhs.charm.sh/vhs-1Bb0tQxFBTi6YSCGGo5SII.gif)

_Refer to [examples/app](examples/app/app.go) for the source code._

## Installation

```bash
go get -u github.com/bartventer/log
```

## Usage

```go
package main

import (
    "github.com/bartventer/log"
)

func main() {
    logger := log.New(log.UseLevel(log.DebugLevel))
    
    logger.Debug("Oh, hi!")
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

This project is a refactored version of the fantastic [charmbracelet/log](https://github.com/charmbracelet/log), which is a part of [Charm](https://charm.sh/).
