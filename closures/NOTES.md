#closures NOTES

A **closure** in Go is *a function value that references variables from outside its own function body*.
It is "bound" to these external (lexical) environment variables and can access and assign to them even after the outer function has finished executing, effectively allowing functions to maintain a persistent state.

## Closer Look at How Closures Work

When an *inner*, anonymous function is defined within an *outer* function and then returned or passed elsewhere, it retains access to the variables in its surrounding scope. Each closure created from the same outer function will have its own independent set of captured variables, ensuring data isolation.

## Key Characteristics and Use Cases

- **State Persistence and Encapsulation:** Closures can maintain private state without needing to define a struct or separate type, effectively allowing them to act as lightweight objects bundling behaviour and data.

- **Data Isolation:** Each *closure instance* has an independent state, meaning changes in one won't affect others originating from the same outer function.

- **Callbacks and Middleware:** Closures are often used when passing functions as arguments (e.g. callbacks), allowing them to carry necessary context or configuration.

## Considerations

- **Memory Usage:** Captured variables may be allocated on the heap, so long-lived closures capturing large objects can retain more memory than expected.

- **Loop Gotchas:** Care must be taken when using closures within loops to ensure the loop variable is captured correctly for each iteration.

See [Understanding Closures in Go: A Beginner’s Guide](https://medium.com/@sanhdoan/understanding-closures-in-go-a-beginners-guide-2795f6dae640) 
