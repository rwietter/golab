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

func generatePairs(next func() string) chan pair {
	pairs := make(chan pair)
	go func() {
		defer close(pairs)
		for {
			word := next()
			if word == "" {
				break
			}
			digits := countDigits(word)
			pairs <- pair{word, digits}
		}
	}()
	return pairs
}

func updateStats(ch chan pair, stats counter, mu *sync.Mutex, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done() // Sinaliza que a goroutine terminou
		for result := range ch {
			mu.Lock()
			stats[result.word] = result.digits
			mu.Unlock()
		}
	}()
}

func countDigitsInWords(next func() string) counter {
	stats := counter{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	pairs := generatePairs(next)

	updateStats(pairs, stats, &mu, &wg)

	wg.Wait()
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
