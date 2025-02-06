package base

import "fmt"

func Loops() {
	// 1. For tradicional
	fmt.Println("\n1. For tradicional:")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}

	// 2. For como while
	fmt.Println("\n\n2. For como while:")
	count := 0
	for count < 3 {
		fmt.Printf("%d ", count)
		count++
	}

	// 3. For infinito com break
	fmt.Println("\n\n3. For infinito com break:")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("%d ", i)
		i++
	}

	// 4. For range em slice
	fmt.Println("\n\n4. For range em slice:")
	numbers := []int{1, 2, 3}
	for index, value := range numbers {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// 5. For range ignorando index
	fmt.Println("\n5. For range ignorando index:")
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}

	// 6. For range em map
	fmt.Println("\n\n6. For range em map:")
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	for key, value := range colors {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	// 7. For range em string (itera sobre runes - caracteres Unicode)
	fmt.Println("\n7. For range em string:")
	text := "Olá"
	for index, char := range text {
		fmt.Printf("index: %d, char: %c\n", index, char)
	}

	// 8. For range em channel
	fmt.Println("\n8. For range em channel:")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // Importante fechar o canal antes do range
	for num := range ch {
		fmt.Printf("%d ", num)
	}

	// 9. For com continue
	fmt.Println("\n\n9. For com continue:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Pula a iteração quando i == 2
		}
		fmt.Printf("%d ", i)
	}

	// 10. For range em array multidimensional
	fmt.Println("\n\n10. For range em array multidimensional:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("[%d][%d]=%d ", i, j, val)
		}
		fmt.Println()
	}

	// 11. For label (útil para nested loops)
	fmt.Println("\n11. For label:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // Quebra o loop externo
			}
			fmt.Printf("[%d,%d] ", i, j)
		}
	}

	// 12. For range em struct slice
	fmt.Println("\n\n12. For range em struct slice:")
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	for _, person := range people {
		fmt.Printf("%s: %d anos\n", person.Name, person.Age)
	}

	// 13. For range com select em múltiplos channels
	fmt.Println("\n13. For range com select em múltiplos channels:")
	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)

	ch1 <- "canal 1"
	ch2 <- "canal 2"

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recebido do canal 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recebido do canal 2:", msg2)
		}
	}
}
