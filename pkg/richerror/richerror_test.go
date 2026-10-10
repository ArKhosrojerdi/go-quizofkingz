package richerror

import (
	"errors"
	"testing"
)

func TestKindAndMessageWalkTheChain(t *testing.T) {
	inner := New("mysql.GetUserByID").
		WithMessage("record not found").
		WithKind(KindNotFound)
	outer := New("userservice.Profile").
		WithErr(inner).
		WithKind(KindOf(inner)).
		WithMessage(MessageOf(inner))

	if KindOf(outer) != KindNotFound {
		t.Fatalf("kind = %d, want not found", KindOf(outer))
	}
	if MessageOf(outer) != "record not found" {
		t.Fatalf("message = %q", MessageOf(outer))
	}
	if !errors.Is(outer, inner) {
		t.Fatal("outer does not unwrap to inner")
	}
}
