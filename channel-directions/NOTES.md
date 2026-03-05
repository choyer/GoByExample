#channel-directions NOTES

- Channel direction can be specified to restrict a channel to only sending or only receiving data, which improves type-safety and code clarity.

## Types of Channel Directions

- Bidirectional (chan T): A channel created with make(chan int) (for example) is bidirectional by default, meaning it can be used for both sending and receiving values of type T.
- Send-only (chan<- T): The arrow points towards the chan keyword, indicating that data can only be sent into the channel. Attempting to receive from a send-only channel results in a compile-time error.
- Receive-only (<-chan T): The arrow points away from the chan keyword, indicating that data can only be received from the channel. Attempting to send to a receive-only channel results in a compile-time error. 
