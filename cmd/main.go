package main

import (
	"context"
	"fmt"
	"github.com/basterrus/sheduler/internal"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Printf("Worker start at %s\n", time.Now().String())

	worker := internal.NewScheduler()

	worker.Add(ctx, internal.TestFunc, 5*time.Second)
	worker.Add(ctx, internal.TestFunc2, 3*time.Second)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	worker.Stop()
	defer fmt.Printf("Worker stoped at %s\n", time.Now().String())
}
