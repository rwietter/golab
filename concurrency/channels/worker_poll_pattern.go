package channels

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// Representa uma tarefa a ser processada
type Job struct {
	ID     int
	Data   string
	Result string
}

// Worker pool para processamento concorrente
type WorkerPool struct {
	numWorkers int
	jobs       chan Job
	results    chan Job
	done       chan struct{}
}

func NewWorkerPool(numWorkers int, bufferSize int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Job, bufferSize),
		results:    make(chan Job, bufferSize),
		done:       make(chan struct{}),
	}
}

func (wp *WorkerPool) Start(ctx context.Context) {
	var wg sync.WaitGroup

	// Inicia os workers
	for i := 0; i < wp.numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			wp.worker(ctx, workerID)
		}(i)
	}

	// Goroutine para fechar o canal de resultados quando todos os workers terminarem
	go func() {
		wg.Wait()
		close(wp.results)
		close(wp.done)
	}()
}

func (wp *WorkerPool) worker(ctx context.Context, id int) {
	for {
		select {
		case job, ok := <-wp.jobs:
			if !ok {
				return // Canal de jobs foi fechado
			}

			// Simula processamento com potencial timeout
			result, err := wp.processJob(ctx, job)
			if err != nil {
				log.Printf("Worker %d: Erro processando job %d: %v", id, job.ID, err)
				continue
			}

			// Envia resultado apenas se o contexto ainda é válido
			select {
			case <-ctx.Done():
				return
			case wp.results <- result:
			}

		case <-ctx.Done():
			log.Printf("Worker %d: Cancelado pelo contexto", id)
			return
		}
	}
}

func (wp *WorkerPool) processJob(ctx context.Context, job Job) (Job, error) {
	// Simula processamento que pode ser cancelado
	select {
	case <-ctx.Done():
		return Job{}, ctx.Err()
	case <-time.After(100 * time.Millisecond):
		job.Result = fmt.Sprintf("Processado: %s", job.Data)
		return job, nil
	}
}

func (wp *WorkerPool) Submit(job Job) error {
	select {
	case wp.jobs <- job:
		return nil
	default:
		return errors.New("buffer de jobs cheio")
	}
}

func (wp *WorkerPool) Results() <-chan Job {
	return wp.results
}

func (wp *WorkerPool) Stop() {
	close(wp.jobs)
	<-wp.done // Espera todos os workers terminarem
}

func WorkerPoolPattern() {
	// Exemplo de uso
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool := NewWorkerPool(3, 10)
	pool.Start(ctx)

	// Submete jobs
	go func() {
		for i := 0; i < 20; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("Dados %d", i),
			}
			if err := pool.Submit(job); err != nil {
				log.Printf("Erro ao submeter job %d: %v", i, err)
			}
		}
		pool.Stop()
	}()

	// Processa resultados
	for job := range pool.Results() {
		fmt.Printf("Resultado do job %d: %s\n", job.ID, job.Result)
	}
}
