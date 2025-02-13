# Namned Goroutines

Nesta etapa apenas refatoramos as goroutines anonimas para funções.

```go
package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type counter map[string]int

type pair struct {
	word   string
	digits int
}

func countDigitsWorker(next func() string, ch chan pair) {
	defer close(ch)
	for {
		word := next()
		if word == "" {
			break
		}
		digits := countDigits(word)
		ch <- pair{word, digits}
	}
}

func fillStatsWorker(ch chan pair, stats counter, done chan struct{}) {
	for result := range ch {
		stats[result.word] = result.digits
	}
	close(done) // sinaliza que a goroutine terminou
}

func countDigitsInWords(next func() string) counter {
	stats := counter{}
	ch := make(chan pair)
	done := make(chan struct{})

	go countDigitsWorker(next, ch)
	go fillStatsWorker(ch, stats, done)

	<-done
	return stats
}

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func asStats(m *sync.Map) counter {
	stats := counter{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func fields() func() string {
	phrase := "0ne 1wo thr33 4068"
	words := strings.Fields(phrase)
	counter := 0
	return func() string {
		if counter > len(words)-1 {
			return ""
		}
		next_word := words[counter]
		counter = counter + 1
		return next_word
	}
}

func main() {
	next := fields()
	counts := countDigitsInWords(next)
	printStats(counts)
}
```