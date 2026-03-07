#closing-channels NOTES

- **Only the sender should close a channel. NEVER the receiver!**
- Sending on a closed channel will cause a runtime panic
- Closing a channel twice (or an already closed channel) WILL cause a runtime panic
- Receiving from a closed channel is SAFE.
- Channels do not have to be closed in all cases, as they will be garbage collected if no goroutines reference them anymore. Closing is primarily a mechanism to signal to the receiving goroutines that no more data is coming, such as when using a for range loop to read all values or to signal a shutdown/completion event.

## Common Idiom for Receiving Values (until a channel is closed)

```
for i := range c {
    fmt.Println(i)
}
// The loop exits automatically after the channel 'c' is closed.
```
```
```

## Detecting Closed channel

```
v, ok := <-ch
if !ok {
    fmt.Println("Channel is closed and drained")
}
```
```
```
