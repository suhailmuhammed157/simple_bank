package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	UserAgent string
	ClientIp  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtd := &Metadata{
		UserAgent: "",
		ClientIp:  "",
	}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		userAgents := md.Get("user-agent")

		if len(userAgents) > 0 {
			mtd.UserAgent = userAgents[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtd.ClientIp = p.Addr.String()
	}

	return mtd

}
