package main

import "fmt"

func PrintHomePage() {
	fmt.Println(`

_________ .____    .___    __                   __               .___.__  __                
\_   ___ \|    |   |   | _/  |_  ____ ___  ____/  |_    ____   __| _/|__|/  |_  ___________ 
/    \  \/|    |   |   | \   __\/ __ \\  \/  /\   __\ _/ __ \ / __ | |  \   __\/  _ \_  __ \
\     \___|    |___|   |  |  | \  ___/ >    <  |  |   \  ___// /_/ | |  ||  | (  <_> )  | \/
 \______  /_______ \___|  |__|  \___  >__/\_ \ |__|    \___  >____ | |__||__|  \____/|__|   
        \/        \/                \/      \/             \/     \/                        


Welcome to the CLI Text Processor!

This tool processes a plain text file and transforms it based on inline commands.

Available commands:
  (cap)       - Capitalizes the previous word.
  (cap, N)    - Capitalizes the previous N words.
  (up)        - Converts the previous word to UPPERCASE.
  (up, N)     - Converts the previous N words to UPPERCASE.
  (low)       - Converts the previous word to lowercase.
  (low, N)    - Converts the previous N words to lowercase.
  (hex)       - Converts the previous word to hexadecimal.
  (hex, N)    - Converts the previous N words to hexadecimal.
  (rev)       - Reverses the previous word.
  (rev, N)    - Reverses the previous N words.
  (pal)       - Converts the previous word to "pali" if it's a palindrome.
  (pal, N)    - Converts the previous N words to "pali" if they're palindromes.
  (len)       - Converts the previous word into its length.
  (len, N)    - Converts the previous N words into their lengths.

Features:
  - Removes extra spaces automatically.
  - Keeps punctuation intact when transforming words.

Usage:
  go run . <input.txt> <output.txt>

Example:
  go run . input.txt output.txt

Use "--help" or "-h" to see this message again.`)
}
