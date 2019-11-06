package lambbar

import (
	"context"
	"fmt"
	"testing"
)

// test if ErrWithContext properly implements the error interface
func TestErrorInterface(t *testing.T) {
	var _ error = ErrWithContext{}
	var _ error = NewErr("lol", fmt.Errorf("nope"))
	var _ error = NewErrWithContext(context.TODO(), "lol", fmt.Errorf("nope"))
}
