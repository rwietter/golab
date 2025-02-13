# Channel Signal

Nesta etapa usamos uma goroutine para producer e outra para consumer. No entanto, no escopo externo não ficou nenhum sinalizador do canal para dizer para nossa main thread esperar para receber dados do canal, então nosso código finalizaria antes de o producer terminal.

Assim sendo, podemos usar `sync.WaitGroup` para aguardar a goroutine finalizar, mas podemos usar um channel `done` quando temos apenas uma goroutine producer. Isso funciona, a main thread fica travada até que o producer termine.

```go
// Counting digits in words.
package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

// counter stores the number of digits in each word.
// The key is the word, and the value is the number of digits.
type counter map[string]int

// solution start
type pair struct {
	word   string
	digits int
}

// countDigitsInWords counts the number of digits in the words of a phrase.
// solution start

// countDigitsInWords counts the number of digits in the words of a phrase.
func countDigitsInWords(next func() string) counter {
	stats := counter{}
	ch := make(chan pair)
	done := make(chan struct{})

	go func() {
		defer close(ch)
		for {
			word := next()
			if word == "" {
				break
			}
			digits := countDigits(word)
			ch <- pair{word, digits}
		}
	}()

	go func() {
		for result := range ch {
			stats[result.word] = result.digits
		}
		close(done) // sinaliza que a goroutine terminou
	}()

	<-done
	return stats
}

// solution end

// countDigits returns the number of digits in a string.
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// asStats converts statistics from sync.Map to a regular map.
func asStats(m *sync.Map) counter {
	stats := counter{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

// printStats prints the number of digits in words.
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

- [Next](part9.md)