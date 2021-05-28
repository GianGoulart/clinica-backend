package app

import (
	"time"

	"github.com/GianGoulart/Clinica_backend/app/health"
	"github.com/GianGoulart/Clinica_backend/app/item"
	"github.com/GianGoulart/Clinica_backend/app/session"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/sirupsen/logrus"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	Health  health.App
	Item    item.App
	Session session.App
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Stores *store.Container

	StartedAt time.Time
	Version   string
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	container := &Container{
		Health:  health.NewApp(opts.Stores, opts.Version, opts.StartedAt),
		Item:    item.NewApp(opts.Stores),
		Session: session.NewApp(nil),
	}

	logrus.Info("Registered -> App")

	return container

}
