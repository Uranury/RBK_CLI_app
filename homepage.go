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

This tool processes a plain text file and transforms it based on inline commands or global commands.

== PROCESSING MODES ==

1. INLINE COMMANDS MODE (default):
   Processes embedded commands within the text.

2. GLOBAL COMMAND MODE:
   Applies a single command to the entire text.

== AVAILABLE COMMANDS ==

Text Case:
  (cap)       - Capitalizes the previous word.
  (cap, N)    - Capitalizes the previous N words.
  (up)        - Converts the previous word to UPPERCASE.
  (up, N)     - Converts the previous N words to UPPERCASE.
  (low)       - Converts the previous word to lowercase.
  (low, N)    - Converts the previous N words to lowercase.

Number Conversion:
  (hex)       - Converts the previous word from hexadecimal to decimal.
  (hex, N)    - Converts the previous N words from hexadecimal to decimal.
  (bin)       - Converts the previous word from binary to decimal.
  (bin, N)    - Converts the previous N words from binary to decimal.

Text Manipulation:
  (rev)       - Reverses the previous word.
  (rev, N)    - Reverses the previous N words.
  (len)       - Converts the previous word into its length.
  (len, N)    - Converts the previous N words into their lengths.
  (pal)       - Converts the previous word to "pali" if it's a palindrome.
  (pal, N)    - Converts the previous N words to "pali" if they're palindromes.

== FEATURES ==
  - Removes extra spaces automatically.
  - Keeps punctuation intact when transforming words.
  - Preserves line breaks and text structure.

== USAGE ==

Inline Commands Mode:
  go run . <input.txt> <output.txt>

Global Command Mode:
  go run . <input.txt> <output.txt> <command>

== EXAMPLES ==

Process file with embedded commands:
  go run . story.txt output.txt

Apply capitalization to entire text:
  go run . story.txt output.txt cap

Convert all binary numbers to decimal:
  go run . binary_data.txt output.txt bin

Convert all text to uppercase:
  go run . document.txt output.txt up

Reverse all words in the text:
  go run . poem.txt output.txt rev

Use "--help" or "-h" to see this message again.`)
}
