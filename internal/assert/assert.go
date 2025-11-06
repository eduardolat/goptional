package assert

import (
	"reflect"
	"testing"
)

// Equal fails the test if expected and actual are not deeply equal.
func Equal[T any](t *testing.T, expected, actual T) {
	t.Helper()
	if !equal(expected, actual) {
		t.Fatalf("expected: %#v, got: %#v", expected, actual)
	}
}

// NoError fails the test if err is not nil.
func NoError(t *testing.T, err error) {
	t.Helper()
	if !noError(err) {
		t.Fatalf("expected no error, got: %v", err)
	}
}

// Error fails the test if err is nil.
func Error(t *testing.T, err error) {
	t.Helper()
	if !hasError(err) {
		t.Fatalf("expected an error but got none")
	}
}

// True fails the test if condition is false.
func True(t *testing.T, condition bool) {
	t.Helper()
	if !isTrue(condition) {
		t.Fatalf("expected condition to be true")
	}
}

// False fails the test if condition is true.
func False(t *testing.T, condition bool) {
	t.Helper()
	if !isFalse(condition) {
		t.Fatalf("expected condition to be false")
	}
}

// Nil fails the test if value is not nil.
func Nil(t *testing.T, value any) {
	t.Helper()
	if !nilCheck(value) {
		t.Fatalf("expected nil, got: %#v", value)
	}
}

// NotNil fails the test if value is nil.
func NotNil(t *testing.T, value any) {
	t.Helper()
	if !notNil(value) {
		t.Fatalf("expected non-nil value")
	}
}

// Len fails the test if the length of v does not equal expected.
func Len[T any](t *testing.T, v T, expected int) {
	t.Helper()
	if !lenCheck(v, expected) {
		val := reflect.ValueOf(v)
		if val.Kind() != reflect.Array && val.Kind() != reflect.Slice && val.Kind() != reflect.Map && val.Kind() != reflect.String {
			t.Fatalf("Len() only supports arrays, slices, maps, and strings; got %T", v)
		}
		t.Fatalf("expected length %d, got %d", expected, val.Len())
	}
}

// Internal testable functions

func equal[T any](expected, actual T) bool {
	return reflect.DeepEqual(expected, actual)
}

func noError(err error) bool {
	return err == nil
}

func hasError(err error) bool {
	return err != nil
}

func isTrue(condition bool) bool {
	return condition
}

func isFalse(condition bool) bool {
	return !condition
}

func nilCheck(value any) bool {
	return isNil(value)
}

func notNil(value any) bool {
	return !isNil(value)
}

func lenCheck[T any](v T, expected int) bool {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice && val.Kind() != reflect.Map && val.Kind() != reflect.String {
		return false
	}
	return val.Len() == expected
}

// helper: safe nil check that works with interfaces and reflect values
func isNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}