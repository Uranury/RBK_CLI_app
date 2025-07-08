package main

import (
	"CLI_app/processes"
	"testing"
)

func TestProcessText(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// Capitalization
		{"It (cap) was the best of times", "It was the best of times"},
		{"it was the best of times (up)", "it was the best of TIMES"},
		{"it was the best of times (cap,2)", "it was the best Of Times"},
		{"one two three (low,2)", "one two three"},
		{"one TWO THREE (low,3)", "one two three"},
		{"keep SOME THINGS (cap,2), as-is.", "keep Some Things, as-is."},
		{"skip unknown COMMANDS (foo,2) safely", "skip unknown COMMANDS (foo,2) safely"},
		{"hello THERE (low,1)", "hello there"},
		{"NOW or NEVER EVER (cap,2)", "NOW or Never Ever"},
		{"hello! (cap)", "Hello!"},
		{"expedition 33... (cap, 2)", "Expedition 33..."},
		{"проверяем работает ли с русской раскладкой (cap,6)", "Проверяем Работает Ли С Русской Раскладкой"},
		{"You didn't know (up, 2)?!", "You DIDN'T KNOW?!"},

		// Binary conversion
		{"data 1010 (bin)", "data 10"},
		{"start 1101. (bin)", "start 13."},
		{"1001; 11101, 1100, (bin, 3)", "9; 29, 12,"},
		{"false case (bin)", "false case"},
		{"bin 1010 1111 (bin,2)", "bin 10 15"},
		{"10010, ((1101))! 390 (bin, 3)", "18, ((13))! 390"},
		{"result 0b1010 (bin)", "result 10"},                       // 0b prefix
		{"value 0B1111! (bin)", "value 15!"},                       // 0B prefix with punctuation
		{"x y z (bin,3)", "x y z"},                                 // Invalid binary (letters)
		{"101 110 111 (bin,3)", "5 6 7"},                           // Multiple valid binary
		{"0b101, 0b110! 0B111. (bin,3)", "5, 6! 7."},               // Multiple with prefixes and punctuation
		{"zz (bin)", "zz"},                                         // Invalid binary
		{"value: 0b1010 1100 1111 (bin,2)", "value: 0b1010 12 15"}, // Partial conversion
		{"binary 1010 1111 (bin,2)", "binary 10 15"},

		// Hex conversion
		{"convert 1f (hex)", "convert 31"},
		{"value FF! (hex)", "value 255!"},
		{"a b c (hex,3)", "10 11 12"},
		{"A B c (hex,3)", "10 11 12"},
		{"ff, a9! 2A. (hex,3)", "255, 169! 42."},
		{"zz (hex)", "zz"}, // Invalid hex
		{"value: 1a2b 3c 0f (hex,2)", "value: 1a2b 60 15"},
		{"hex 1a ff (hex,2)", "hex 26 255"},
		{"result 0x1f (hex)", "result 31"},                     // 0x prefix
		{"value 0XFF! (hex)", "value 255!"},                    // 0X prefix with punctuation
		{"10 11 12 (hex,3)", "16 17 18"},                       // Valid hex digits
		{"0x1a, 0x2b! 0X3c. (hex,3)", "26, 43! 60."},           // Multiple with prefixes and punctuation
		{"gg (hex)", "gg"},                                     // Invalid hex
		{"value: 0x1a2b 3c 0f (hex,2)", "value: 0x1a2b 60 15"}, // Partial conversion with prefix
		{"hexadecimal 1a ff (hex,2)", "hexadecimal 26 255"},

		{"number (1010) (bin)", "number (10)"},           // Single parentheses
		{"result ((1111))! (bin)", "result ((15))!"},     // Double parentheses
		{"values (((1010)))? (bin)", "values (((10)))?"}, // Triple parentheses
		{"data ((ff))! (hex)", "data ((255))!"},          // Hex with double parentheses
		{"result (((1a)))? (hex)", "result (((26)))?"},   // Hex with triple parentheses
		{"mixed (1010), ((1111))! (bin,2)", "mixed (10), ((15))!"},
		{"empty () (bin)", "empty ()"},                 // Empty parentheses
		{"just parens (()) (bin)", "just parens (())"}, // Empty nested parentheses
		{"0b (bin)", "0b"},                             // Just prefix
		{"0x (hex)", "0x"},                             // Just prefix
		{"(0b1010) (bin)", "(10)"},                     // Prefix inside parentheses
		{"(0xff) (hex)", "(255)"},

		// Len conversion
		{"banana (len)", "6"},
		{"banana and apple, (len,3)", "6 3 5,"},
		{"привет (len)", "6"},
		{"!!!! (len)", "!!!!"},
		{"((World hello, there!)) (len, 3)", "((5 5, 5!))"},

		// Reversing the word
		{"hello everyone (rev,2)", "olleh enoyreve"},
		{"привет всем (rev, 2)", "тевирп месв"},

		// Check if the word is palindrome
		{"texet (pal)", "pali"},
		{"texet, word, apple, level, radar (pal, 5)", "pali, word, apple, pali, pali"},
		{"a (pal)", "pali"},
		{"hey (pal)", "hey"},
		{"this ((level)) is so dope! (pal, 5)", "this ((pali)) is so dope!"},
		{"чекаем че будет с русской боб (pal)", "чекаем че будет с русской pali"},

		// Basic robustness tests
		{"hello world (cap, -1)", "hello World"},
		{"hello world (cap, abc)", "hello World"},
		{"hello world (cap) (rev) (up)", "hello DLROW"},
		{"test (cap, 2) (low, 1)", "test"},
		{"测试 中文 (up)", "测试 中文"},
		{"hello\nworld (cap)", "hello\nWorld"},
		{"-1110110 (bin)", "-118"},
		{"0xFF (hex)", "255"},
		{"3e8, (hex)", "1000,"},
		{"(cap)", ""},
		{"(hello) (rev)", "(olleh)"},
		{"((hello)), (rev)", "((olleh)),"},
		{"((hello)), (cap)", "((Hello)),"},
		{"((hello, everyone)), (cap, 2)", "((Hello, Everyone)),"},
		{"text with (cap", "text with (cap"},
		{"(rev (cap)", "(Rev"},
		{"hello(cap)guys", "Helloguys"},
		{"abc (rev, abc)", "cba"},
		{"hello (rev,1) (up,1)", "OLLEH"},
		{"one two three (cap,10)", "One Two Three"},
		{"  (rev)", ""},
		{"one(cap ,2)", "One"},
		{"word (cap, 1)!", "Word!"},
		{"-101010 (bin)", "-42"},
		{"one two (rev, 2) (pal, 2)", "eno owt"},
		{"hello! world? (cap,2)", "Hello! World?"},
		{"(cap)(rev)", ""},
		{"((world))! (len)", "((5))!"},
		{"((world))! (up)", "((WORLD))!"},
		{"(cap ,   2)", ""},
		{"abc123, (up)", "ABC123,"},
		{"word (rev, 1) (cap, 1)", "Drow"},
		{"  ", " "},

		// Grammar check, getting rid of extra spaces
		{"Hello    World", "Hello World"},
		{"I 'm here", "I'm here"},
		{"I'm here (rev, 2)", "m'I ereh"},
		{"i'm (up)", "I'M"},
		{"i'm (cap)", "I'm"},
		{"rock 'n' roll (up, 3)", "ROCK 'N' ROLL"},
		{"didn't (rev)", "tndi'd"},
		{"it's fine (up, 2)", "IT'S FINE"},
		{"((Hello everyone here, how are you doing?)!)} (up, 7)", "((HELLO EVERYONE HERE, HOW ARE YOU DOING?)!)}"},

		// Quotes
		{"he said \"hello\" (cap,3)", "He Said \"Hello\""},
		{"'hello world' (cap,2)", "'Hello World'"},
		{"it's 'important' here (cap,3)", "It's 'Important' Here"},

		// Mixed punctuation
		{"word?! (cap)", "Word?!"},
		{"really?!?! (up)", "REALLY?!?!"},
		{"wow... (cap)", "Wow..."},
		{"end.) (cap)", "End.)"},
		{"word@domain (cap)", "Word@domain"},
		{"user@email.com (cap)", "User@email.com"},
		{"word#tag (cap)", "Word#tag"},
		{"50% off (cap,2)", "50% Off"},

		// Extra
		{"well‑known (cap)", "Well‑known"},
		{"foo_bar (up)", "FOO_BAR"},
		{"a|b|c (rev)", "c|b|a"},
		{"path/to/file (low)", "path/to/file"},
		{"one two\nthree four (cap,2)", "one two\nThree Four"},
		{"((hello,)) (up)", "((HELLO,))"},
		{"((hello,)) (rev)", "((olleh,))"},
		{"((101011,)) (bin)", "((43,))"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := processes.ProcessText(tt.input)
			if got != tt.want {
				t.Errorf("processText(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
