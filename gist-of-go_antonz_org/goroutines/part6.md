# Channel Sync

Nesta etapa podemos otimizar nossa forma de sincronização com canais, na parte 4, nós usamos um `sync.Store` Map para sincronizar a contagem, com canais podemos enviar a contagem de dígitos realizados em cada thread pelo canal e receber posteriormente quando a contagem termina e inserir no nosso stats store e teremos o resultado de dígitos por palavras:

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

// countDigitsInWords counts the number of digits in the words of a phrase.
// solution start

// countDigitsInWords counts the number of digits in the words of a phrase.
func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)
	stats := make(counter)

	go func() {
		defer close(counted) // Fecha o channel quando a goroutine terminar

		// count the number of digits in each,
		// and write it to the counted channel.
		for _, word := range words {
			counted <- countDigits(word)
		}
	}()

	// Read values from the counted channel
	// and fill stats.
	for _, word := range words {
		digits := <-counted
		stats[word] = digits
	}

	// As a result, stats should contain words
	// and the number of digits in each.

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

func main() {
	phrase := "0ne 1wo thr33 4068"
	counts := countDigitsInWords(phrase)
	fmt.Printf("Counts of digits in words: %v\n", counts)
	printStats(counts)
}
```

- [Next](part7.md)