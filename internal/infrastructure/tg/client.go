package tg

import (
	"context"
	"errors"
	"github.com/bmstu-itstech/itsreg-api/pkg/endpoint"
	tgv1 "github.com/bmstu-itstech/itsreg-proto/gen/go/tg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log/slog"
)

var (
	ErrNotFound = status.Error(codes.InvalidArgument, "not found")
)

type Service struct {
	log    *slog.Logger
	client tgv1.TgServiceClient
}

func NewService(
	log *slog.Logger,
	host string,
	port int,
) (*Service, error) {
	conn, err := grpc.NewClient(endpoint.Addr(host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := tgv1.NewTgServiceClient(conn)

	return &Service{
		log:    log,
		client: client,
	}, nil
}

func (s *Service) Start(
	ctx context.Context,
	botId int64,
) error {
	req := &tgv1.StartRequest{
		BotId: botId,
	}

	_, err := s.client.Start(ctx, req)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return errors.Join(ErrNotFound, err)
		}
		return err
	}

	return nil
}

func (s *Service) Stop(
	ctx context.Context,
	botId int64,
) error {
	req := &tgv1.StopRequest{
		BotId: botId,
	}

	_, err := s.client.Stop(ctx, req)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return errors.Join(ErrNotFound, err)
		}
		return err
	}

	return nil
}
