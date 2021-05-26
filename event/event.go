package event

import (
	"github.com/nats-io/nats.go"
	"github.com/tradersclub/PocArquitetura/app"
	"github.com/tradersclub/PocArquitetura/event/item"
	"github.com/tradersclub/TCUtils/logger"
)

// Options struct de opções para a criação de uma instancia das rotas
type Options struct {
	Apps *app.Container
	Nats *nats.Conn
}

// Register handler instance
func Register(opts Options) {
	item.Register(opts.Apps, opts.Nats)

	logger.Info("Registered -> Event (Nats)")

}
