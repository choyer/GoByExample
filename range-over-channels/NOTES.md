#range-over-channels NOTES

- `for ... range` loop is the standard way to **receive values from a channel repeatedly until it is closed.**
- The loop automatically terminates when the sender calls `close()` on the channel AND all values have been received.
