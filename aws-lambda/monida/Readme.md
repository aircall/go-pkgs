# Monida

*MONItor your lambDA*

This module provides utilities to wrap golang lambda functions with rollbar.

## Usage:

1. Setup `ENVIRONMENT` and `ROLLBAR_API_TOKEN` in your lambda environment variables

2. Use wrapper to start and send errors

```go
// import package
import (
  "github.com/aircall/go-pkgs/aws-lambda/monida"
  "errors"
)

// Use wrapper to start your lambda function
func main () {
  monida.WrapAndStart(handler)
}

func handler(ctx context.Context) error {

  err := errors.New("nope")

  // Log error using context
  monida.LogErrWithContext(ctx, err, "doing stuff failed")

  // string interpolation is supported
  monida.LogErrWithContext(ctx, err, "doing stuff %v time failed", 1)

  // Log error without context
  monida.LogErr(err, "doing stuff failed")

  // send warning with context
  yourName := "you"
  monida.WarnWithContext(ctx, "I'm warning %s", yourName)

  // send warning without context
  monida.Warn("I'm warning you")

  // returned err will also be logged by the wrapper
  return NewErrWithContext(ctx, err, "run failed")
}

```
