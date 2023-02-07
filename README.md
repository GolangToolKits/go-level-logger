# go-level-logger
A three level logger for Golang. It has the following levels: debug, info, and error.

### Examples

```go
import(
    logger "github.com/GolangToolKits/go-level-logger"
)

var lg logger.Logger

log:= lg.New()
log.SetLogLevel(logger.AllLevel)

log.Debug("this is debug logging", "this too")

```
