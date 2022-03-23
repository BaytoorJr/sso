package middleware

import "github.com/BaytoorJr/sso/src/service"

type Middleware func(service service.MainService) service.MainService
