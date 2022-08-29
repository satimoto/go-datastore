package param

import "github.com/satimoto/go-datastore/pkg/db"

func NewUpdatePsbtFundingStateParams(psbtFundingState db.PsbtFundingState) db.UpdatePsbtFundingStateParams {
	return db.UpdatePsbtFundingStateParams{
		ID:         psbtFundingState.ID,
		Psbt:       psbtFundingState.Psbt,
		FundedPsbt: psbtFundingState.FundedPsbt,
		SignedPsbt: psbtFundingState.SignedPsbt,
		SignedTx:   psbtFundingState.SignedTx,
		IsFailed:   psbtFundingState.IsFailed,
	}
}
