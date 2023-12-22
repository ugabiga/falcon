package model

import "github.com/ugabiga/falcon/internal/service"

type Layout struct {
	Claim   *service.JWTClaim
	IsLogin bool
}

type LayoutPage struct {
	Layout Layout
}

type Toast struct {
	Message string
}
