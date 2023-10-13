package BrainFuck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadfile(t *testing.T) {
	output := readfile("Testfiles/invalid.bf")
	expected := ""
	assert.Equal(t, expected, output, "Output does not match the expected value")
}

func TestHelloword(t *testing.T) {
	code := readfile("Testfiles/helloworld.bf")

	output, _ := interpreter(code, nil)
	expected := "Hello World!\n"

	assert.Equal(t, expected, output, "Output does not match the expected value")
}

func TestHelloword_complex(t *testing.T) {
	code := readfile("Testfiles/helloworld_complex.bf")

	output, _ := interpreter(code, nil)
	expected := "Hello World!\n"

	assert.Equal(t, expected, output, "Output does not match the expected value")
}

func TestHelloword_short(t *testing.T) {
	code := readfile("Testfiles/helloworld_short.bf")

	output, _ := interpreter(code, nil)
	expected := "Hello, World!"

	assert.Equal(t, expected, output, "Output does not match the expected value")
}

func TestHiddenmessage(t *testing.T) {
	code := readfile("Testfiles/hiddenmessage.bf")

	output, _ := interpreter(code, nil)
	expected := "If you see this, the interpreter works!!\n"

	assert.Equal(t, expected, output, "Output does not match the expected value")
}

func TestSquare(t *testing.T) {
	code := readfile("Testfiles/square.b")

	output, _ := interpreter(code, nil)
	expected := ""

	// square the numbers 0-10000
	for i := 0; i <= 10000; i++ {
		if isPerfectSquare(i) {
			expected += fmt.Sprintf("%d\n", i)
		}
		// expected += fmt.Sprintf("%d\n", i*i)
	}

	assert.Equal(t, expected, output, "Output does not match the expected value")
}

// Apparently perfect squares can only end in 0, 1, 4, 5, 6, or 9
// https://en.wikipedia.org/wiki/Square_number#Properties
func isPerfectSquare(n int) bool {
	lastDigit := n % 10
	if lastDigit == 2 || lastDigit == 3 || lastDigit == 7 || lastDigit == 8 {
		return false // Last digit is not 0, 1, 4, 5, 6, or 9
	}

	root := int64(0)
	for root*root <= int64(n) {
		if root*root == int64(n) {
			return true
		}
		root++
	}
	return false
}

// The following tests are from: http://brainfuck.org/tests.b
func TestTapesize(t *testing.T) {
	code := readfile("Testfiles/tapesize.b")

	output, _ := interpreter(code, nil)
	expected := "#\n"

	assert.Equal(t, expected, output, "Output does not match the expected value")
}

func TestObscure(t *testing.T) {
	code := readfile("Testfiles/obscure.b")

	output, _ := interpreter(code, nil)
	expected := "H\n"

	assert.Equal(t, expected, output, "Output does not match the expected value")
}

func TestBroken1(t *testing.T) {
	code := readfile("Testfiles/broken1.b")

	_, err := interpreter(code, nil)

	assert.Error(t, err, "Error: No matching ']' found.")
}

func TestBroken2(t *testing.T) {
	code := readfile("Testfiles/broken2.b")

	_, err := interpreter(code, nil)

	assert.Error(t, err, "Error: No matching '[' found.")
}

// TODO: add tests for remaining test files
