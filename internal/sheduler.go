package internal

import (
	"context"
	"sync"
	"time"
)

type Task func(ctx context.Context)

type Scheduler struct {
	wg     *sync.WaitGroup
	cancel []context.CancelFunc
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		wg:     new(sync.WaitGroup),
		cancel: make([]context.CancelFunc, 0),
	}
}

func (s *Scheduler) Add(ctx context.Context, t Task, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancel = append(s.cancel, cancel)

	s.wg.Add(1)
	go s.execution(ctx, t, interval)
}

func (s *Scheduler) execution(ctx context.Context, t Task, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			t(ctx)
		case <-ctx.Done():
			s.wg.Done()
			ticker.Stop()
			return
		}
	}
}

func (s *Scheduler) Stop() {
	for _, cancel := range s.cancel {
		cancel()
	}
	s.wg.Wait()
}
