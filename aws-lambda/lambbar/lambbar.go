// Package lambbar provides a wrapper to run aws lambda with rollbar, along with contextualized errors that can automatically be sent to rollbar and stdout (-> cloudwatch)
package lambbar

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rollbar/rollbar-go"
	"log"
	"os"
)

// ErrWithContext defines a structured error that can be used to wrap and send formatted errors to both rollbar and aws cloudwatch
type ErrWithContext struct {
	Ctx context.Context
	Err error
	Msg string
}

var initError error = nil

func init() {
	rollbarAPIToken := os.Getenv("ROLLBAR_API_TOKEN")
	if rollbarAPIToken == "" {
		// used to panic on WrapAndStart. Otherwise we cannnot even start the test suite
		initError = fmt.Errorf("please provide a ROLLBAR_API_TOKEN env variable")
		return
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		initError = fmt.Errorf("please provide a ENVIRONMENT env variable")
		return
	}
	rollbar.SetToken(rollbarAPIToken)
	rollbar.SetEnvironment(env)
}

// WrapAndStart wraps a handler function with rollbar, and start execution with aws lambda
func WrapAndStart(handler interface{}) {
	if initError != nil {
		panic(initError)
	}
	rollbar.WrapAndWait(func() {
		lambda.Start(handler)
	})
}

// NewErr creates a ErrWithContext with empty context
func NewErr(msg string, err error) ErrWithContext {
	return ErrWithContext{
		Msg: msg,
		Ctx: context.TODO(),
		Err: err,
	}
}

// NewErrWithContext returns a new ErrWithContext struct
func NewErrWithContext(ctx context.Context, msg string, err error) ErrWithContext {
	return ErrWithContext{
		Msg: msg,
		Ctx: ctx,
		Err: err,
	}
}

// LogErrWithContext creates and send a ErrWithContext
func LogErrWithContext(ctx context.Context, msg string, err error) {
	lerr := NewErrWithContext(ctx, msg, err)
	lerr.Log()
}

// LogErr creates and send an error without any context
func LogErr(msg string, err error) {
	lerr := NewErr(msg, err)
	lerr.Log()
}

// WarnWithContext sends a warning message to rollbar with context
func WarnWithContext(ctx context.Context, msg string) {
	log.Printf("WARNING: %s\n", msg)
	rollbar.ErrorWithExtrasAndContext(ctx, "warning", fmt.Errorf(msg), nil)
}

// Warn sends a warning message to rollbar
func Warn(msg string) {
	log.Printf("WARNING: %s\n", msg)
	rollbar.Error(context.TODO(), "warning", fmt.Errorf(msg), nil)
}

// Error prints the error as as string
func (e ErrWithContext) Error() string { return e.Msg + ": " + e.Err.Error() }

// Unwrap returns the wrapped error
func (e *ErrWithContext) Unwrap() error { return e.Err }

// Log sends the error to both stdout and rollbar
func (e ErrWithContext) Log() {
	log.Println(e.Error())
	rollbar.ErrorWithExtrasAndContext(e.Ctx, "error", e, nil)
}
