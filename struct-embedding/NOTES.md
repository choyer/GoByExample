#struct-embedding NOTES

- Struct embedding is a powerful form of **composition** that allows a struct to include the fields and methods of another struct, promoting code reuse without traditional class inheritance. Think of it as a "has-a" relationship, where the outer struct has an instance of the embedded struct as a field, which is implicity named by its type.
- Access to embedded structs can using shorthand or full path names.
- Be mindful of name collisions. If two embedded structs (at the same depth) have the same field or method name, direct access using the shorthand will result in a compiler error.
