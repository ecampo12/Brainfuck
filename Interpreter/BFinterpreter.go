package BFInterpreter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Readfile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// were just going to get rid of any characters that are not BF code
	pattern := "[><+\\-\\[\\]\\.\\,]+"
	regex := regexp.MustCompile(pattern)

	code := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		// could just do the operations as we read,
		// but loops might be an issue
		matches := regex.FindAllString(line, -1)
		code += strings.Join(matches, "")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	return code
}

func Interpreter(code string, inputs []string) (string, error) {
	res := ""
	var err error
	// original implementation used a 30,000 byte array
	tape := make([]byte, 30000)

	// used to keep track of loops
	// stack := []int{}

	// pointer to the current cell
	ptr := 0

	// loop through the code
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			ptr = (ptr + 1) % len(tape)
		case '<':
			ptr = (ptr - 1 + len(tape)) % len(tape)
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			// fmt.Printf("%c", tape[ptr])
			res += string(tape[ptr])
		case ',':
			// use the next input value or get a new one from stdin
			if len(inputs) > 0 {
				tape[ptr] = inputs[0][0]
				fmt.Println(inputs[0])
				inputs = inputs[1:]
			} else {
				fmt.Scanf("%c", &tape[ptr])
			}
		case '[':
			// If the current cell is zero, jump forward to the matching ']'.
			if tape[ptr] == 0 {
				count := 1 // Count of unmatched '[' characters
				for count > 0 {
					i++
					if i >= len(code) {
						err = errors.New("error: No matching ']' found")
						// fmt.Println("Error: No matching ']' found.")
						return "", err
					}
					if code[i] == '[' {
						count++
					} else if code[i] == ']' {
						count--
					}
				}
			}
		case ']':
			// If the current cell is not zero, jump back to the matching '['.
			if tape[ptr] != 0 {
				count := 1 // Count of unmatched ']' characters
				for count > 0 {
					i--
					if i < 0 {
						err = errors.New("error: No matching '[' found")
						// fmt.Println("Error: No matching '[' found.")
						return "", err
					}
					if code[i] == ']' {
						count++
					} else if code[i] == '[' {
						count--
					}
				}
			}

		default:
			// fmt.Println("Error: Invalid character in code.")
			err = errors.New("error: Invalid character in code")
			return "", err
		}
		// fmt.Printf("ptr: %d, code: %c, codeptr: %d, data: %c, res: %s\n", ptr, code[i], i, tape[ptr], res)
	}
	return res, err
}
