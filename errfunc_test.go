package errfunc_test

import (
	"fmt"
	"testing"

	"deedles.dev/errfunc"
)

func TestNoError(t *testing.T) {
	ef := errfunc.New(func(v int) (float64, error) {
		return 2 * float64(v), nil
	})

	var n float64
	n += ef.Call(3)
	n += ef.Call(2)

	if ef.Err() != nil {
		t.Fatalf("expected nil error, but got %q", ef.Err())
	}
	if n != 10 {
		t.Fatalf("expected total result to be 10, but got %v", n)
	}
}

func TestError(t *testing.T) {
	ef := errfunc.New(func(v int) (float64, error) {
		return float64(v), fmt.Errorf("%v", v)
	})

	var n float64
	n += ef.Call(3)
	n += ef.Call(2)

	if ef.Err() == nil {
		t.Fatalf("expected error to be non-nil")
	}
	if ef.Err().Error() != "3" {
		t.Fatalf("expected error to be \"3\", but got %q", ef.Err())
	}
	if n != 3 {
		t.Fatalf("expected total result to be 3, but got %v", n)
	}
}
