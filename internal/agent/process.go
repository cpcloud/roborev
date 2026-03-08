package agent

import (
	"context"
	"io"
	"os/exec"
	"sync"
	"sync/atomic"
	"time"
)

var subprocessWaitDelay = 5 * time.Second

func configureSubprocess(cmd *exec.Cmd) {
	cmd.WaitDelay = subprocessWaitDelay
}

func closeOnContextDone(ctx context.Context, c io.Closer) func() {
	if c == nil || ctx.Done() == nil {
		return func() {}
	}
	done := make(chan struct{})
	var once sync.Once
	var stopped atomic.Bool
	go func() {
		select {
		case <-ctx.Done():
			if stopped.Load() {
				return
			}
			_ = c.Close()
		case <-done:
		}
	}()
	return func() {
		stopped.Store(true)
		once.Do(func() {
			close(done)
		})
	}
}
