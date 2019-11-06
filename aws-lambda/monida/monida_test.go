package monida

import (
	"context"
	"errors"
	"testing"
)

func TestErrorInterface(t *testing.T) {
	var _ error = ErrWithContext{} // implement error interface

	// produce proper message
	expectedStr := "I got 99 problems: but a silent exception ain't one"

	// without context
	err := NewErr(errors.New("but a silent exception ain't one"), "I got %v problems", 99)
	if err.Error() != expectedStr {
		t.Errorf("expected %s to equal %s", err.Error(), expectedStr)
	}

	// with context
	ctxErr := NewErrWithContext(context.TODO(), errors.New("but a silent exception ain't one"), "I got %v problems", 99)
	if ctxErr.Error() != expectedStr {
		t.Errorf("expected %s to equal %s", ctxErr.Error(), expectedStr)
	}
}
