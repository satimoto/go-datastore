package param

import "github.com/satimoto/go-datastore/pkg/db"

func NewUpdateChannelRequestParams(channelRequest db.ChannelRequest) db.UpdateChannelRequestParams {
	return db.UpdateChannelRequestParams{
		ID:            channelRequest.ID,
		FundingAmount: channelRequest.FundingAmount,
		SettledMsat:   channelRequest.SettledMsat,
		Status:        channelRequest.Status,
	}
}
