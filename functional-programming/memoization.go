package fp

import (
	"fmt"
	"sync"
)

var (
	cache     = make(map[int]int) // 1. Initialize a map 'cache' to store computed results.
	cacheLock sync.Mutex          // 2. Create a Mutex to control concurrent access to the cache.
)

/*
Memorization, or memoization, is a programming technique that involves
storing results of functions that are expensive in terms of execution
time and reusing them when the same function items are supplied again.
This technique is used to update functions that have high computational
complexity, saving design time.

It is important to note that in concurrent environments, such as Go
applications, it is necessary to ensure that a cache is accessed securely
using mutexes, as in the example above.
*/
func expensiveFunction(n int) int {
	cacheLock.Lock()          // 1. Lock the cache to ensure exclusive access to it.
	result, found := cache[n] // 2. Check if the result for 'n' is already in the cache.
	cacheLock.Unlock()        // 3. Unlock the cache after reading from it.

	// 4. If the result is found in the cache, return it.
	if found {
		return result
	}

	result = n * 2 // 5. Otherwise, if the result is not in the cache, compute it (in this case, multiply 'n' by 2).

	cacheLock.Lock()   // 6. Lock the cache again to update it.
	cache[n] = result  // 7. Store the computed result in the cache.
	cacheLock.Unlock() // 8. Unlock the cache.

	return result // 9. Return the computed result.
}

func Memoization() {
	a := expensiveFunction(5)
	b := expensiveFunction(5) // Call expansiveFunction memoized
	c := expensiveFunction(10)

	fmt.Println("Memoization:", a, b, c) // Output: 10 10 20
}
