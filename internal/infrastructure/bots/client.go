package bots

import (
	"context"
	"errors"
	"github.com/bmstu-itstech/itsreg-api/internal/model"
	"github.com/bmstu-itstech/itsreg-api/pkg/endpoint"
	botsv1 "github.com/bmstu-itstech/itsreg-proto/gen/go/bots"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log/slog"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
)

type Service struct {
	log    *slog.Logger
	client botsv1.BotsServiceClient
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

	client := botsv1.NewBotsServiceClient(conn)

	return &Service{
		log:    log,
		client: client,
	}, nil
}

func (s *Service) Create(
	ctx context.Context,
	bot model.Bot,
) (int64, error) {
	req := &botsv1.CreateRequest{
		Name:   bot.Name,
		Token:  bot.Token,
		Start:  bot.Start,
		Blocks: blocksToPb(bot.Blocks),
	}

	res, err := s.client.Create(ctx, req)
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			return 0, errors.Join(ErrInvalidArgument, err)
		}
		return 0, err
	}

	return res.BotId, nil
}
