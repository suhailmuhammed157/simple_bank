package gapi

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GrpcLogger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

	currentTime := time.Now()
	resp, err = handler(ctx, req)
	responseTime := time.Since(currentTime) // calculate how much time it took after receiving the response

	statusCode := codes.OK

	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	logger := log.Info()
	if err != nil {
		logger = log.Err(err)
	}

	logger.
		Str("method", info.FullMethod).
		Str("protocol", "grpc").
		Dur("Response_time", responseTime).
		Int("status_code", int(statusCode)).
		Str("status_text", statusCode.String()).
		Msg("received grpc request")

	return resp, err

}
