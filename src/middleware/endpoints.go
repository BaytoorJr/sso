package middleware

import "github.com/BaytoorJr/sso/src/service"

type Endpoints struct {
}

func MakeEndpoints(s service.MainService) *Endpoints {
	return &Endpoints{}
}
