package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/rdl/core/internal/platform/config"
)

type Conn = nats.Conn

func Connect(cfg config.NATSConfig) (*nats.Conn, error) {
	nc, err := nats.Connect(cfg.URL,
		nats.MaxReconnects(10),
		nats.ReconnectWait(nats.DefaultReconnectWait),
	)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
