package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// This type Prefix is a Markov Chain prefix of one or two words
type Prefix []string

// The String method returns the Prefix as a string for use as a map key
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// The Shift method removes the first word from the Prefix and appends
// the given word
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// The Chain struct contains a map ("chain") of prefixes to a list of
// suffixes. A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain     map[string][]string
	prefixLen int
}

// The NewChain function returns the address of a Chain with prefixes of
// prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

// The Build method reads text from the provided Reader and parses it
// into prefixes and suffixes that are stored in Chain.
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

// The Generate method returns a string of at most n words generated
// from Chain.
func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func main() {
	// Register command-line flags.
	numwords := flag.Int("words", 100, "maximum number of words to print")
	prefixLen := flag.Int("prefix", 2, "prefix length in words")

	flag.Parse()                     // Parse the command line flags
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

	c := NewChain(*prefixLen)     // Initialize a new chain.
	c.Build(os.Stdin)             // Build chains from standard input.
	text := c.Generate(*numwords) // Generate text.
	fmt.Println(text)             // Write text to standard input.
}
