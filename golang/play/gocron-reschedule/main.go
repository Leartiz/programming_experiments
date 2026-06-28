package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	scheduler, err := gocron.NewScheduler(
		gocron.WithLimitConcurrentJobs(1, gocron.LimitModeWait),
	)
	if err != nil {
		panic(err)
	}

	_, err = scheduler.NewJob(gocron.DurationJob(100*time.Millisecond), gocron.NewTask(func() {
		fmt.Println("Task 1: start (sleep 2s)")
		time.Sleep(2 * time.Second)
		fmt.Println("Task 1: done")
	}))
	if err != nil {
		panic(err)
	}

	scheduler.NewJob(gocron.DurationJob(500*time.Millisecond), gocron.NewTask(func() {
		fmt.Println("Task 2: start (sleep 1s)")
		time.Sleep(1 * time.Second)
		fmt.Println("Task 2: done")
	}))

	scheduler.NewJob(gocron.DurationJob(1*time.Second), gocron.NewTask(func() {
		fmt.Println("Task 3: start (sleep 1s)")
		time.Sleep(1 * time.Second)
		fmt.Println("Task 3: done")
	}))

	scheduler.Start()
	time.Sleep(15 * time.Second)
}
