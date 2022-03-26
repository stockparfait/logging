# Logging convenience methods

This package implements logging by levels: `debug`, `info`, `warning` and
`error`.

This package is inspired by, and is designed to be loosely compatible to the
[`luci/common/logging`](https://pkg.go.dev/go.chromium.org/luci/common/logging)
package.  The reason for reimplementation is that LUCI is a huge mono-repo,
which may be an overkill for smaller projects, or projects unrelated to Cloud
apps.

This repo is dedicated to a very light-weght implementation of a simple API that
gets most of the job done.

## Installation

```
go get github.com/stockparfait/logging
```

### Example usage

```go
package main

import (
	"context"

	"github.com/stockparfait/logging"
)

func main() {
	ctx := logging.Use(context.Background(), logging.DefaultGoLogger(logging.Info))
	logging.Debugf(ctx, "this shouldn't show")
	logging.Infof(ctx, "try %d", 42)
	logging.Warningf(ctx, "you've been warned")
	logging.Errorf(ctx, "that's wrong!")
}
```

This should print something like this:

```
2022/03/26 13:56:16 INFO: try 42
2022/03/26 13:56:16 WARNING: you've been warned
2022/03/26 13:56:16 ERROR: that's wrong!
```

## Development

Clone and initialize the repository, run tests:

```sh
git clone git@github.com:stockparfait/logging.git
cd logging
make init
make test
```
