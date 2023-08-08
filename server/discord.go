package server

import (
	"context"

	"github.com/bsdlp/taxonomy/rpc/taxonomy"
)

func (srv *Server) LookupDiscordUser(context.Context, *taxonomy.LookupDiscordUserRequest) (*taxonomy.LookupDiscordUserResponse, error) {
	return nil, nil
}

func (srv *Server) LinkDiscordUser(context.Context, *taxonomy.LinkDiscordUserRequest) (*taxonomy.LinkDiscordUserResponse, error) {
	return nil, nil
}

func (srv *Server) UnlinkDiscordUser(context.Context, *taxonomy.UnlinkDiscordUserRequest) (*taxonomy.UnlinkDiscordUserResponse, error) {
	return nil, nil
}
