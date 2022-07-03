package internal

import (
	"context"
	"fmt"
	"time"
)

func TestFunc(ctx context.Context) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Start task at %s\n", time.Now().String())
}

func TestFunc2(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, time.Second*5)

	i := 0
	for {
		time.Sleep(time.Millisecond * 100)
		i++
		fmt.Printf("%d ", i)

		select {
		case <-ctx.Done():
			fmt.Println()
			return
		default:
		}
	}
}
