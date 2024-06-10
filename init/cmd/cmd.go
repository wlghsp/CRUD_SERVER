package cmd

import (
	"CRUD_SERVER/config"
	"CRUD_SERVER/network"
	"CRUD_SERVER/repository"
	"CRUD_SERVER/service"
)

type Cmd struct {
	config *config.Config

	network    *network.Network
	repository repository.Repository
	service    *service.Service
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filepath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
