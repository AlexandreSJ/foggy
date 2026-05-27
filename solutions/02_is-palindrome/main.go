package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// Learned a little bit about regex in go, using "regexp" package.
// I also had to try a little bit before I simply used char := string[i]
// which is way easier than I remebered.

// I believe it can be simpler than that, but this is already very readable
// For next solutions i will be trying to use less comments tho,
// except for high complexity functions

func main() {
	scanner := bufio.NewScanner(os.Stdin) // create a new scanner
	scanner.Scan()                        // user input
	input := scanner.Text()               // converts to string
	println(isPalindrome(input))          // prints output
}

func isPalindrome(input string) bool {
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")            // simple regex for alphanumerics only
	s := strings.ToLower(regex.ReplaceAllString(input, "")) // repalce all and lowercase

	// Uncomment for debug regex:
	// fmt.Printf("debug regex> \"%s\"\n", s)

	for i := 0; i < len(s); i++ {
		char := s[i]               // get the character at index i
		if char != s[len(s)-1-i] { // compare with oposite character
			return false
		}
	}
	return true
}
