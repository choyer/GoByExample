#channel-synchronization NOTES

- The core of channel synchronization lies in their blocking behavior and the "happens-before" memory model guarantees.
- By default, sends on a channel (ch <- v) and receives from a channel (v := <-ch) block until the other side is ready. This inherent blocking ensures that a value is transferred only when both the sender and receiver goroutines are synchronized at that point in the code.
