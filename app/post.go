package void

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	feemarketpost "github.com/skip-mev/feemarket/x/feemarket/post"
)

// PostHandlerOptions are the options required for constructing a FeeMarket PostHandler.
type PostHandlerOptions struct {
	AccountKeeper   feemarketpost.AccountKeeper
	BankKeeper      feemarketpost.BankKeeper
	FeeMarketKeeper feemarketpost.FeeMarketKeeper
}

// NewPostHandler returns a PostHandler chain with the fee deduct decorator.
func NewPostHandler(options PostHandlerOptions) (sdk.PostHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errors.New("account keeper is required for post builder")
	}

	if options.BankKeeper == nil {
		return nil, errors.New("bank keeper is required for post builder")
	}

	if options.FeeMarketKeeper == nil {
		return nil, errors.New("feemarket keeper is required for post builder")
	}

	postDecorators := []sdk.PostDecorator{
		feemarketpost.NewFeeMarketDeductDecorator(
			options.AccountKeeper,
			options.BankKeeper,
			options.FeeMarketKeeper,
		),
	}

	return sdk.ChainPostDecorators(postDecorators...), nil
}
