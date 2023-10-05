package ws

import (
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func wsLoop(conn net.Conn) {
	client := newClient(conn)
	defer client.close()

	r := wsutil.NewReader(conn, ws.StateServerSide)

	for {
		// Wait for next frame from client
		hdr, err := r.NextFrame()
		if err != nil {
			if err != io.EOF {
				log.Panicln("failed to read frame", err)
			}
			return
		}

		if hdr.OpCode == ws.OpClose {
			return
		}

		// Drop any data
		_, _ = io.Copy(io.Discard, r)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go wsLoop(conn)
}
