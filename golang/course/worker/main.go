package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// ErrFailed и ErrManual - причины остановки цикла.
var ErrFailed = errors.New("failed")
var ErrManual = errors.New("manual")

// Worker выполняет заданную функцию в цикле, пока не будет остановлен.
// Гарантируется, что Worker используется только в одной горутине.
type Worker struct {
	fn func() error

	ctx        context.Context
	cancelFunc context.CancelCauseFunc

	afterStopFuncs []func()
}

// NewWorker создает новый экземпляр Worker с заданной функцией.
// Но пока не запускает цикл с функцией.
func NewWorker(fn func() error) *Worker {
	return &Worker{
		fn:  fn,
		ctx: nil, // !

		afterStopFuncs: make([]func(), 0),
	}
}

// Start запускает отдельную горутину, в которой циклически
// выполняет заданную функцию, пока не будет вызван метод Stop,
// либо пока функция не вернет ошибку.
// Повторные вызовы Start игнорируются.
func (w *Worker) Start() {
	if w.ctx != nil {
		return
	}

	w.ctx, w.cancelFunc =
		context.WithCancelCause(context.Background())

	go func() {
		var running = true
		for running {
			select {
			case <-w.ctx.Done():
				running = false
			default:
				err := w.fn()
				if err != nil {
					w.cancelFunc(ErrFailed)
					running = false
				}
			}
		}

		// ***

		if w.afterStopFuncs != nil {
			for i := range w.afterStopFuncs {
				w.afterStopFuncs[i]()
			}
		}
	}()
}

// ----------------------------------------------

// Stop останавливает выполнение цикла.
// Вызов Stop до Start игнорируется.
// Повторные вызовы Stop игнорируются.
func (w *Worker) Stop() {
	if w.ctx == nil {
		return
	}

	w.cancelFunc(ErrManual)
}

// AfterStop регистрирует функцию, которая
// будет вызвана после остановки цикла.
// Можно зарегистрировать несколько функций.
// Вызовы AfterStop после Start игнорируются.
func (w *Worker) AfterStop(fn func()) {
	if w.ctx != nil {
		return
	}

	w.afterStopFuncs = append(w.afterStopFuncs, fn)
}

// Err возвращает причину остановки цикла:
// - ErrManual - вручную через метод Stop;
// - ErrFailed - из-за ошибки, которую вернула функция.
func (w *Worker) Err() error {
	if w.ctx == nil {
		return nil
	}

	return context.Cause(w.ctx)
}

// ----------------------------------------------

func main() {
	count := 9
	fn := func() error {
		fmt.Print(count, " ")
		count--
		time.Sleep(10 * time.Millisecond)
		return nil
	}

	worker := NewWorker(fn)
	worker.Start()
	time.Sleep(105 * time.Millisecond)
	worker.Stop()

	fmt.Printf("worker err: %v\n",
		worker.Err())

	// other tests...
}
