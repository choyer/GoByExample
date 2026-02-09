#variadic-functions NOTES

- General info on [Variadic Functions](https://en.wikipedia.org/wiki/Variadic_function)
- A function can have only one variadic parameter. It must be the LAST one in the function signature.
- Mixed fixed and variadic parameters are possible.
- If no arguments are passed, the variadic paramter inside the function will be nil (or empty slice), which is SAFE to range over.
