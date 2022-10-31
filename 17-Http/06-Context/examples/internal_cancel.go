package examples

import (
	"context"
	"time"
)

func InternalCancellationExample() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	//Timeout
	//ctx, cancel = context.WithTimeout(ctx, time.Second * 3)
	
	//Deadline
	//ctx, cancel = context.WithDeadline(ctx, time.Date(2022, time.February, 13, 23, 30, 0, 0, time.UTC))

	defer cancel()

	go func() {
		for {
			if time.Now().Second()%12 == 0 {
				cancel()
			}
		}
	}()

	Processor(ctx, 0)
}

func Processor(ctx context.Context, num int) {
	for {
		num++
		select {
		case <-ctx.Done():
			println("Canceled...")
			return
		default:
			println("processing", num)
		}
		time.Sleep(time.Millisecond * 500)
	}
}
