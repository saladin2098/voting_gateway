package handlers

import "github.com/saladin2098/month3/lesson11/voting_gateway/servers"

type Handler struct {
	Clients servers.Clients
}

func NewHandler() *Handler {
	return &Handler{
        Clients: *servers.NewClients(),
    }
}
