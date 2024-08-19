package assert

import (
"testing"
)

func Equal(t *testing.T, actual, expected interface{}) {
    t.Helper()
    if actual != expected {
        t.Errorf("got: %v; want: %v", actual, expected)
    }
}