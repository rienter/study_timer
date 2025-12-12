package timer

import "errors"

type Timer struct {
	current  int
	starting int
	running  bool
}

func InitTimer(minutes int) Timer {
	return Timer{
		starting: minutes * 60,
		current:  minutes * 60,
		running:  true,
	}
}

func (t Timer) Current() int {
	return t.current
}

func (t Timer) Elapsed() int {
	return t.starting - t.current
}

func (t Timer) Starting() int {
	return t.starting
}

func (t Timer) Running() bool {
	return t.running && (t.current > 0)
}

func (t *Timer) TogglePause() {
	t.running = !t.running
}

func (t Timer) Finished() bool {
	return !(t.current > 0)
}

func (t *Timer) Decrease() error {
	if t.Finished() {
		return errors.New("The timer has already finished")
	}
	t.current = t.current - 1
	return nil
}
