package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
)

// ContractKeeper defines the entry points to a smart contract that must be exposed by the VM module
type ContractKeeper interface {
	// IBCAcknowledgementPacketCallback is called in the source chain when a packet acknowledgement
	// is received. The contract is expected to handle the callback within the user defined
	// gas limit, and handle any errors, state reversals, out of gas, or panics gracefully.
	// The user may also pass a custom message to the contract. (May be nil)
	// Gas limit is set to 0 if the user does not specify a custom gas limit or
	// if the remaining gas is less than the custom gas limit.
	IBCAcknowledgementPacketCallback(
		ctx sdk.Context,
		packet channeltypes.Packet,
		customMsg []byte,
		ackResult channeltypes.Acknowledgement,
		relayer sdk.AccAddress,
		contractAddr string,
		gasLimit uint64,
	) error
	// IBCPacketTimeoutCallback is called in the source chain when a packet is not received before
	// the timeout height. The contract is expected to handle the callback within the user defined
	// gas limit, and handle any errors, state reversals, out of gas, or panics gracefully.
	// Gas limit is set to 0 if the user does not specify a custom gas limit or
	// if the remaining gas is less than the custom gas limit.
	IBCPacketTimeoutCallback(
		ctx sdk.Context,
		packet channeltypes.Packet,
		relayer sdk.AccAddress,
		contractAddr string,
		gasLimit uint64,
	) error
	// IBCReceivePacketCallback is called in the destination chain when a packet is received.
	// The contract is expected to handle the callback within the user defined gas limit, and
	// handle any errors, state reversals, out of gas, or panics gracefully.
	// The user may also pass a custom message to the contract. (May be nil)
	// Gas limit is set to 0 if the user does not specify a custom gas limit or
	// if the remaining gas is less than the custom gas limit.
	IBCReceivePacketCallback(
		ctx sdk.Context,
		packet channeltypes.Packet,
		customMsg []byte,
		ackResult channeltypes.Acknowledgement,
		relayer sdk.AccAddress,
		contractAddr string,
		gasLimit uint64,
	) error
}