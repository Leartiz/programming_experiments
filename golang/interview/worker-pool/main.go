package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
	Fn func() error
}

type WorkerPool struct {
	tasks     chan Task
	results   chan error
	workerNum int
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	mu        sync.Mutex
	closed    bool
}

func NewWorkerPool(workerNum int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		tasks:     make(chan Task),
		results:   make(chan error),
		workerNum: workerNum,
		ctx:       ctx,
		cancel:    cancel,
	}
}

func (p *WorkerPool) Start() {
	for i := 0; i < p.workerNum; i++ {
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *WorkerPool) worker() {
	defer p.wg.Done()
	for {
		select {
		case <-p.ctx.Done():
			return
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			// выполняем задачу с таймаутом
			ctx, cancel := context.WithTimeout(p.ctx, 2*time.Second)
			defer cancel()

			done := make(chan error, 1)
			go func() {
				done <- task.Fn()
			}()

			select {
			case err := <-done:
				p.results <- err
			case <-ctx.Done():
				p.results <- fmt.Errorf("task %d timed out", task.ID)
			}
		}
	}
}

func (p *WorkerPool) Submit(task Task) error {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return fmt.Errorf("pool is closed")
	}
	p.mu.Unlock()

	select {
	case p.tasks <- task:
		return nil
	case <-p.ctx.Done():
		return fmt.Errorf("pool is shutting down")
	}
}

func (p *WorkerPool) Shutdown(ctx context.Context) error {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return nil
	}
	p.closed = true
	p.mu.Unlock()

	close(p.tasks)
	p.cancel()

	done := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (p *WorkerPool) Results() <-chan error {
	return p.results
}

func main() {
	pool := NewWorkerPool(3)
	pool.Start()

	for i := 0; i < 5; i++ {
		id := i
		err := pool.Submit(Task{
			ID: id,
			Fn: func() error {
				time.Sleep(time.Duration(id) * time.Second)
				fmt.Println("Task", id, "done")
				return nil
			},
		})
		if err != nil {
			fmt.Println("Submit error:", err)
		}
	}

	time.Sleep(3 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Shutdown(ctx); err != nil {
		fmt.Println("Shutdown error:", err)
	}
	fmt.Println("Shutdown complete")
}
