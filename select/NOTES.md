#select NOTES

- The `select` statement is a control structure used to **wait on and handle multiple channel operations** (send or receive) concurrently.
- It acts like a switch statement. It **blocks** the current goroutine until *one* of its cases can proceed.
