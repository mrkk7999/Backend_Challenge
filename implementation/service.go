package implementation

import "backend_challenge/iface"

type service struct {
}

func New() iface.Service {
	return &service{}
}
