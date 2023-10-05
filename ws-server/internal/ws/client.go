package ws

import (
	"log"
	"net"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type wsClient struct {
	conn net.Conn
}

var (
	lock    sync.RWMutex
	clients = make(map[*wsClient]struct{})
)

func newClient(conn net.Conn) *wsClient {
	c := &wsClient{
		conn: conn,
	}

	lock.Lock()
	clients[c] = struct{}{}
	lock.Unlock()

	return c
}

func (c *wsClient) close() {
	c.conn.Close()

	lock.Lock()
	delete(clients, c)
	lock.Unlock()
}

func (c *wsClient) send(data []byte) error {
	w := wsutil.NewWriter(c.conn, ws.StateServerSide, ws.OpText)

	if _, err := w.Write(data); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}

func broadcast(data []byte) {
	lock.RLock()
	defer lock.RUnlock()

	for client := range clients {
		if err := client.send(data); err != nil {
			log.Println("send to client failed", err)
			continue
		}

	}
}
