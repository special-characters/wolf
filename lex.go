package wolf

import (
	"bufio"
	"math/rand"
	"strings"
	"time"
)

type Lexicon struct {
	words [65536]string
	index map[string]int
}

func New() *Lexicon {
	lex := &Lexicon{index: make(map[string]int)}
	scanner := bufio.NewScanner(strings.NewReader(words))

	for i := 0; scanner.Scan(); i++ {
		word := scanner.Text()
		lex.words[i] = word
		lex.index[word] = i
	}

	return lex
}

func (l *Lexicon) Random(n int) string {
	var output []string

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		output = append(output, l.words[r.Intn(len(l.words))])
	}

	return strings.Join(output, "-")
}
