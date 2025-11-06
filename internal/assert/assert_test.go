package assert

import (
	"errors"
	"testing"
)

func TestEqual(t *testing.T) {
	t.Run("equal values", func(t *testing.T) {
		if !equal(42, 42) {
			t.Error("expected equal values to return true")
		}
	})
	t.Run("not equal values", func(t *testing.T) {
		if equal("a", "b") {
			t.Error("expected not equal values to return false")
		}
	})
}

func TestNoError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		if !noError(nil) {
			t.Error("expected nil error to return true")
		}
	})
	t.Run("with error", func(t *testing.T) {
		if noError(errors.New("fail")) {
			t.Error("expected non-nil error to return false")
		}
	})
}

func TestError(t *testing.T) {
	t.Run("with error", func(t *testing.T) {
		if !hasError(errors.New("boom")) {
			t.Error("expected non-nil error to return true")
		}
	})
	t.Run("no error", func(t *testing.T) {
		if hasError(nil) {
			t.Error("expected nil error to return false")
		}
	})
}

func TestTrue(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		if !isTrue(true) {
			t.Error("expected true to return true")
		}
	})
	t.Run("false", func(t *testing.T) {
		if isTrue(false) {
			t.Error("expected false to return false")
		}
	})
}

func TestFalse(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		if !isFalse(false) {
			t.Error("expected false to return true")
		}
	})
	t.Run("true", func(t *testing.T) {
		if isFalse(true) {
			t.Error("expected true to return false")
		}
	})
}

func TestNil(t *testing.T) {
	t.Run("nil value", func(t *testing.T) {
		if !nilCheck(nil) {
			t.Error("expected nil value to return true")
		}
	})
	t.Run("non-nil value", func(t *testing.T) {
		if nilCheck(1) {
			t.Error("expected non-nil value to return false")
		}
	})
}

func TestNotNil(t *testing.T) {
	t.Run("non-nil value", func(t *testing.T) {
		if !notNil(1) {
			t.Error("expected non-nil value to return true")
		}
	})
	t.Run("nil value", func(t *testing.T) {
		if notNil(nil) {
			t.Error("expected nil value to return false")
		}
	})
}

func TestLen(t *testing.T) {
	t.Run("correct length", func(t *testing.T) {
		if !lenCheck([]int{1, 2, 3}, 3) {
			t.Error("expected correct length to return true")
		}
	})
	t.Run("wrong length", func(t *testing.T) {
		if lenCheck([]int{1}, 3) {
			t.Error("expected wrong length to return false")
		}
	})
	t.Run("unsupported type", func(t *testing.T) {
		if lenCheck(123, 1) {
			t.Error("expected unsupported type to return false")
		}
	})
}
