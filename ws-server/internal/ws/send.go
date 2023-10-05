package ws

import (
	"context"
	"log"
)

var queue = make(chan []byte, 1000)

func Send(data []byte) {
	select {
	case queue <- data:
	default:
		log.Println("send failed: queue is full")
	}
}

func SendLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-queue:
			broadcast(msg)
		}
	}
}
