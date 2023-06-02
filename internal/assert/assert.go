package assert

import (
	"strings"
	"testing"
)

// Equal uses generics to test that the arguments are equal.
func Equal[T comparable](t *testing.T, got, expected T) {
	t.Helper()

	if got != expected {
		t.Errorf("got: %v; want: %v", got, expected)
	}
}

// StringContains checks response body content for some text
func StringContains(t *testing.T, actual, expectedSubstring string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("got: %q; expected to contain: %q", actual, expectedSubstring)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got: %v; expected: nil", actual)
	}
}
