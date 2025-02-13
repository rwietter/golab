package channels

import (
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				fmt.Printf("Worker %d terminando...\n", id)
				return
			}
			fmt.Printf("Worker %d processando tarefa %d\n", id, task)
		case <-done:
			fmt.Printf("Worker %d parando...\n", id)
			return
		}
	}
}

func WorkerPool() {
	wg := sync.WaitGroup{}
	tasks := make(chan int, 10)
	done := make(chan struct{})

	// Inicia 3 workers
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, done, &wg)
	}

	// Envia tarefas para o canal
	for i := 1; i <= 5; i++ {
		tasks <- i
	}
	close(tasks)

	// Espera um pouco para os workers terminarem
	wg.Wait()

	// Sinaliza para os workers pararem
	close(done)
	fmt.Println("Fim!")
}
