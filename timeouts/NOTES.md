#timeouts NOTES

- For channel timeouts use `select` with a case in the format of `case <-time.After(1 * time.Second):`
