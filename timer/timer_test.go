package timer

import (
	"reflect"
	"testing"
)

func TestCurrent(t *testing.T) {
	timer := InitTimer(60)

	expected := 60 * 60
	got := timer.Current()

	assertEqInt(t, expected, got)
}

func TestElapsed(t *testing.T) {
	timer := InitTimer(1)

	expected := 0
	got := timer.Elapsed()

	assertEqInt(t, expected, got)
}

func TestInitTimer(t *testing.T) {
	timer := InitTimer(60)

	expected := Timer{
		starting: 60 * 60,
		current:  60 * 60,
		running:  false,
	}

	assertEqAny(t, expected, timer)
}

func TestTimePassing(t *testing.T) {
	timer := InitTimer(60)
	timer.Decrease()

	got := timer.current
	expected := 3599

	assertEqInt(t, expected, got)
}

func TestTogglePause(t *testing.T) {
	timer := Timer{
		current:  60,
		starting: 70,
		running:  false,
	}

	timer.TogglePause()
	got := timer.running
	expected := true

	assertEqBool(t, expected, got)
}

func assertEqInt(t testing.TB, expected, got int) {
	t.Helper()

	if expected != got {
		t.Errorf("\nGot: %d\nExpected: %d", got, expected)
	}
}

func assertEqString(t testing.TB, expected, got string) {
	t.Helper()

	if expected != got {
		t.Errorf("\nGot: %s\nExpected: %s", got, expected)
	}
}

func assertEqBool(t testing.TB, expected, got bool) {
	t.Helper()

	if expected != got {
		t.Errorf("\nGot: %t\nExpected: %t", got, expected)
	}
}

func assertEqAny(t testing.TB, expected, got any) {
	t.Helper()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("\nGot: %d\nExpected: %d", got, expected)
	}
}
