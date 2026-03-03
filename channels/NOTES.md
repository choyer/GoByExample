#channels notes

- Channels act as a typed conduit for communication and synchronization between concurrently running goroutines.
- The direction of the arrow operator indicates the direction of data flow
  - `ch <- v` :Send value `v` to channel `ch`
  - `v := <-ch` :Receive a value from channel `ch` and assign it to `v`
- By default, send & receive operations on channels are **blocking**. Thus send operations will wait until a receiver is ready, and a receive operation will wait until the sender is ready. This prevents race conditions.
