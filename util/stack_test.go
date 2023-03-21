package util

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	var want, got any
	//Push first element
	stack.Push(1)
	want = 1
	got = stack.Size()
	if got != want {
		t.Fatalf("expected size: %d, got: %v\n", want, got)
	}

	// Push second element
	stack.Push(6)
	want = 2
	got = stack.Size()
	if got != want {
		t.Fatalf("expected size: %d, got: %v\n", want, got)
	}

	// Check emptiness
	want = false
	got = stack.Empty()
	if got != want {
		t.Fatalf("expected empty: %t, got: %t\n", want, got)
	}

	var err error

	// Check peek
	want = 6
	got, err = stack.Peek()
	if err != nil {
		t.Fatalf("expected nil error, got %v\n", err)
	}
	if got != want {
		t.Fatalf("expected peek: %v, got: %v\n", want, got)
	}

	stack2 := NewStack()
	wantErrorStr := "stack is empty"
	_, gotErr := stack2.Peek()
	if gotErr.Error() != wantErrorStr {
		t.Fatalf("expected error: %v, got: %v\n", wantErrorStr, gotErr)
	}

	// Check Pop
	want = 6
	got, err = stack.Pop()
	if err != nil {
		t.Fatalf("expected nil error, got %v\n", err)
	}
	if got != want {
		t.Fatalf("expected pop: %v, got: %v\n", want, got)
	}

	_, gotErr = stack2.Pop()
	if gotErr.Error() != wantErrorStr {
		t.Fatalf("expected error: %v, got: %v\n", wantErrorStr, gotErr)
	}

	// String
	stack.Push(4)
	stack.Push(6)
	stack.Push(7)
	want = "values currently held in the stack are: 7 6 4 1 \n"
	got = stack.String()
	if got != want {
		t.Fatalf("expected: %q, got: %q\n", want, got)
	}
}
