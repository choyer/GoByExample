#strings-and-runes NOTES

- A string is an immutable, read-only slice of bytes (`[]byte`)
- Go treats strings as containers of UTF-8 text by default.
- `len()` returns the number of **bytes**, NOT necessarily the number of visible characters (code points/runes)
- Indexing a string, yields the byte at that index. For multi-byte characters, accessing a single byte will not return a valid character.

- A rune is 32-bit integer (`int32`) that represents a single Unicode code point. It is designed to hold any character in the Unicode standard, which may take up to 4 bytes in UTF-8.
