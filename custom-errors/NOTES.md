#custom-errors NOTES

- For simple unformatted error messages, use `errors.New`
- For formatted messages, use `fmt.Errorf`
- For adding additional data or context, define a struct and implement the `Error()` method
- `errors.Is` checks if an error in the chain matches a specific sentinel error value
- `errors.As` extracts a specific custom error type from an error chain, allowing access to its internal fields
