package types

import (
	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/ibc-go/v7/modules/core/exported"
)

var _ exported.ClientMessage = &Header{}

// ClientType defines that the Header is a Wasm client consensus algorithm
func (h Header) ClientType() string {
	return exported.Wasm
}

// ValidateBasic defines a basic validation for the wasm client header.
func (h Header) ValidateBasic() error {
	if len(h.Data) == 0 {
		return errorsmod.Wrap(ErrInvalidData, "data cannot be empty")
	}

	return nil
}
