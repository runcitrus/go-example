package app

import (
	"context"
	_ "embed"
	"net/http"
	"time"

	"example/internal/ws"
)

//go:embed index.html
var indexHTML []byte

func webSPA(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(indexHTML)
}

func exampleBroadcast(ctx context.Context) {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case v := <-tick.C:
			currentTime := v.Format("2006-01-02 15:04:05")
			ws.Send([]byte(currentTime))
		}
	}
}

func Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go ws.SendLoop(ctx)
	go exampleBroadcast(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/ws/", ws.Handler)
	mux.HandleFunc("/", webSPA)

	return http.ListenAndServe(":8080", mux)
}
