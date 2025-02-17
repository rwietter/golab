package pool

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// Job representa uma tarefa que será processada pelo worker pool
// Cada job tem um ID único, dados de entrada (Data) e um resultado após processamento
type Job struct {
	ID     int    // Identificador único do job
	Data   string // Dados de entrada para processamento
	Result string // Resultado após o processamento
}

// WorkerPool gerencia um conjunto de workers que processam jobs concorrentemente
// Implementa um padrão de pool de workers para controlar a concorrência
type WorkerPool struct {
	numWorkers int           // Número de workers concorrentes
	jobs       chan Job      // Canal para distribuir jobs aos workers
	results    chan Job      // Canal para coletar resultados dos workers
	done       chan struct{} // Canal para sinalizar quando todos os workers terminaram
}

// NewWorkerPool cria uma nova instância de WorkerPool
// numWorkers: número de workers concorrentes
// bufferSize: tamanho do buffer dos canais de jobs e resultados
func NewWorkerPool(numWorkers int, bufferSize int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Job, bufferSize), // Canal com buffer para jobs
		results:    make(chan Job, bufferSize), // Canal com buffer para resultados
		done:       make(chan struct{}),        // Canal para sinalização de término
	}
}

// Start inicia o pool de workers
// - Cria numWorkers goroutines para processar jobs
// - Usa WaitGroup para controlar o término de todos os workers
// - Monitora cancelamento via context
func (wp *WorkerPool) Start(ctx context.Context) {
	var wg sync.WaitGroup

	// Inicia os workers em goroutines separadas
	for i := 0; i < wp.numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			wp.worker(ctx, workerID)
		}(i)
	}

	// Goroutine para monitorar término dos workers
	// Fecha canais quando todos workers terminarem
	go func() {
		wg.Wait()
		close(wp.results) // Fecha canal de resultados
		close(wp.done)    // Sinaliza término completo
	}()
}

// worker é a função executada por cada worker do pool
// - Escuta continuamente por novos jobs
// - Processa jobs recebidos
// - Monitora cancelamento via context
func (wp *WorkerPool) worker(ctx context.Context, id int) {
	for {
		select {
		// Tenta receber um novo job
		case job, ok := <-wp.jobs:
			if !ok {
				return // Canal de jobs foi fechado, worker deve terminar
			}

			// Processa o job recebido
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
				// Resultado enviado com sucesso
			}

		// Monitora cancelamento do contexto
		case <-ctx.Done():
			log.Printf("Worker %d: Cancelado pelo contexto", id)
			return
		}
	}
}

// processJob processa um único job
// - Simula processamento com delay
// - Pode ser cancelado via context
func (wp *WorkerPool) processJob(ctx context.Context, job Job) (Job, error) {
	// Simula processamento que pode ser cancelado
	select {
	case <-ctx.Done():
		return Job{}, ctx.Err() // Retorna erro se contexto foi cancelado
	case <-time.After(100 * time.Millisecond):
		// Simula processamento do job
		job.Result = fmt.Sprintf("Processado: %s", job.Data)
		return job, nil
	}
}

// Submit submete um novo job para processamento
// Implementa controle de backpressure:
// - Se o canal de jobs estiver cheio, retorna erro
// - Não bloqueia, permitindo que o caller decida como lidar com sobrecarga
func (wp *WorkerPool) Submit(job Job) error {
	// Tenta enviar job de forma não-bloqueante
	select {
	case wp.jobs <- job:
		return nil
	default:
		// Canal cheio - necessário lidar com sobrecarga
		return errors.New("sistema sobrecarregado")
	}
}

// Results retorna um canal somente-leitura para consumir resultados
// O caller pode iterar sobre este canal para receber resultados processados
func (wp *WorkerPool) Results() <-chan Job {
	return wp.results
}

// Stop inicia o desligamento gracioso do pool
// - Fecha o canal de jobs para sinalizar aos workers que devem terminar
// - Espera todos os workers terminarem
func (wp *WorkerPool) Stop() {
	close(wp.jobs) // Sinaliza workers para terminarem
	<-wp.done      // Espera todos workers terminarem
}

// WorkerPoolPattern demonstra o uso do padrão de worker pool
func WorkerPoolPattern() {
	// Cria contexto com timeout de 5 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cria pool com 3 workers e buffer para 10 jobs. Isso significa que até 10 jobs podem ser submetidos para processamento antes de bloquear em Submit (backpressure)
	pool := NewWorkerPool(3, 10)
	pool.Start(ctx)

	// Goroutine para submeter jobs
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
		pool.Stop() // Inicia desligamento após submeter todos jobs
	}()

	// Processa resultados conforme ficam disponíveis
	for job := range pool.Results() {
		fmt.Printf("Resultado do job %d: %s\n", job.ID, job.Result)
	}
}
