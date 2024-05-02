package cmd

import (
	"CRUD_SERVER/config"
	"CRUD_SERVER/network"
	"fmt"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filepath string) *Cmd {

	fmt.Println("여기는 뜹니다.")

	c := &Cmd{
		config:  config.NewConfig(filepath),
		network: network.NewNetwork(),
	}

	c.network.ServerStart(c.config.Server.Port)

	return c
}
