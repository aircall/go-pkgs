# Lambbar

*Lambda meets rollbar*

This module provides utilities to wrap golang lambda functions with rollbar.

## Usage:

1. Setup `ENVIRONMENT` and `ROLLBAR_API_TOKEN` in your lambda environment variables

2. Use wrapper to start and send errors

```go
// import package
import (
  "github.com/aircall/go-pkgs/aws-lambda/lambbar"
  "fmt"
)

// Use wrapper to start your lambda function
func main () {
  lambbar.WrapAndStart(handler)
}

func handler(ctx context.Context) error {

  err := fmt.Errorf("nope")

  // Log error using context
  lambbar.LogErrWithContext(ctx, "doing stuff failed", err)

  // Log error without context
  lambbar.LogErr("doing stuff failed", err)

  // send warning with context
  lambbar.WarnWithContext(ctx, "I'm warning you")

  // send warning without context
  lambbar.Warn("I'm warning you")

  // returned err will also be logged by the wrapper
  return NewErrWithContext(ctx, "run failed", err)
}

```
