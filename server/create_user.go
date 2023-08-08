package server

import (
	"context"

	"github.com/bsdlp/taxonomy/rpc/taxonomy"
)

func (srv *Server) CreateUser(context.Context, *taxonomy.CreateUserRequest) (*taxonomy.CreateUserResponse, error) {
	return nil, nil
}
