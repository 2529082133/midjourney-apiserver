package service

import (
	"github.com/2529082133/midjourney-apiserver/internal/common"
	"github.com/2529082133/midjourney-apiserver/pkg/api"
)

type Service struct {
	api.UnimplementedAPIServiceServer
	*common.Base
}

func New(base *common.Base) *Service {
	return &Service{
		Base: base,
	}
}
