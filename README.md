# CLI Text Processor

A command-line tool written in Go that processes a plain text file, transforming it based on inline commands.

## 🚀 Features

- Capitalizes, uppercases, lowercases, reverses, or transforms words based on commands.
- Handles multiple-word transformations (e.g., `(cap, 3)`).
- Automatically removes extra spaces.
- Keeps punctuation intact.

## 🧠 Supported Commands

| Command     | Description                                                        |
|-------------|--------------------------------------------------------------------|
| `(cap)`     | Capitalizes the previous word                                      |
| `(cap, N)`  | Capitalizes the previous N words                                   |
| `(up)`      | Converts the previous word to UPPERCASE                            |
| `(up, N)`   | Converts the previous N words to UPPERCASE                         |
| `(low)`     | Converts the previous word to lowercase                            |
| `(low, N)`  | Converts the previous N words to lowercase                         |
| `(hex)`     | Converts the previous word to hexadecimal                          |
| `(hex, N)`  | Converts the previous N words to hexadecimal                       |
| `(rev)`     | Reverses the previous word                                         |
| `(rev, N)`  | Reverses the previous N words                                      |
| `(pal)`     | Converts the previous word to `"pali"` if it's a palindrome        |
| `(pal, N)`  | Converts the previous N words to `"pali"` if they're palindromes   |
| `(len)`     | Converts the previous word to its character length                 |
| `(len, N)`  | Converts the previous N words to their character lengths           |

## 📦 Usage

```bash
go run . <input.txt> <output.txt>
