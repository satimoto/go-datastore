package channelrequest

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type ChannelRequestRepository interface {
	CreateChannelRequest(ctx context.Context, arg db.CreateChannelRequestParams) (db.ChannelRequest, error)
	CreateChannelRequestHtlc(ctx context.Context, arg db.CreateChannelRequestHtlcParams) (db.ChannelRequestHtlc, error)
	GetChannelRequest(ctx context.Context, id int64) (db.ChannelRequest, error)
	GetChannelRequestByPaymentHash(ctx context.Context, paymentHash []byte) (db.ChannelRequest, error)
	GetChannelRequestHtlc(ctx context.Context, channelRequestID int64) (db.ChannelRequestHtlc, error)
	GetChannelRequestHtlcByCircuitKey(ctx context.Context, arg db.GetChannelRequestHtlcByCircuitKeyParams) (db.ChannelRequestHtlc, error)
	ListChannelRequestHtlcs(ctx context.Context, channelRequestID int64) ([]db.ChannelRequestHtlc, error)
	ListUnsettledChannelRequestHtlcs(ctx context.Context, channelRequestID int64) ([]db.ChannelRequestHtlc, error)
	UpdateChannelRequest(ctx context.Context, arg db.UpdateChannelRequestParams) (db.ChannelRequest, error)
	UpdateChannelRequestByChannelPoint(ctx context.Context, arg db.UpdateChannelRequestByChannelPointParams) (db.ChannelRequest, error)
	UpdateChannelRequestStatus(ctx context.Context, arg db.UpdateChannelRequestStatusParams) (db.ChannelRequest, error)
	UpdateChannelRequestHtlc(ctx context.Context, arg db.UpdateChannelRequestHtlcParams) (db.ChannelRequestHtlc, error)
}

func NewRepository(repositoryService *db.RepositoryService) ChannelRequestRepository {
	return ChannelRequestRepository(repositoryService)
}
