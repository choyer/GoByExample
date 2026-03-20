#timers NOTES

- The `time` package provides **timers** to execute code **once** after a specified duration. `time.Timer` represents a single future event.
- `time.NewTimer(d Duration) *Timer` creates a new `Timer` that will send the current time on its channel `C` after the duration `d` has elapsed. The program can then block on this channel using `<- timer.C` to wait for the event.
