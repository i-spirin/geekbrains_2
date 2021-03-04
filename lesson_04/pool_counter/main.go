package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	jobs            = make(chan struct{})
	resource        = make(chan struct{}, 10)
	counter         int
	done            = make(chan struct{})
	ctx, cancelFunc = context.WithCancel(context.Background())
	sigs            = make(chan os.Signal, 1)
)

func main() {
	log.Println("Created context")

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go manager(ctx)
	go waiter(cancelFunc)

	for i := 0; i < cap(resource); i++ {
		log.Println("Starting worker")
		go incrementor(&counter)
	}
	select {
	case <-ctx.Done():
		for i := 0; i < cap(resource); i++ {
			resource <- struct{}{}
		}
	case <-done:
		log.Println("count done")
	case signal := <-sigs:
		log.Println("Got signal", signal)
	}
	close(resource)
	log.Println("Last value is:", counter)
}

func waiter(cancelFunc context.CancelFunc) {
	log.Println("Waiter started")
	time.Sleep(10 * time.Minute)
	cancelFunc()
	log.Println("Ending execution")
}

func manager(ctx context.Context) {
	log.Println("Manager started")
	for job := 0; ; job++ {
		select {
		case <-ctx.Done():
			close(jobs)
			return
		default:
			// log.Printf("create job %d\n", job)
			jobs <- struct{}{}
		}
	}
}

func incrementor(value *int) {
	log.Println("Started new worker")
	defer func() {
		<-resource
	}()
	for range jobs {
		time.Sleep(time.Second)
		if *value >= 1000 {
			log.Println("Worker end")
			done <- struct{}{}
			return
		}
		*value++
		log.Println("VALUE IS:", *value)
	}
}
