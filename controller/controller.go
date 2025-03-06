package controller

import "backend_challenge/iface"

type Controller struct {
	service iface.Service
}

func New(service iface.Service) *Controller {
	return &Controller{
		service: service,
	}
}
